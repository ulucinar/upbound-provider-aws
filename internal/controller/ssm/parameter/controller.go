/*
Copyright 2022 Upbound Inc.
*/

package parameter

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	tjcontroller "github.com/upbound/upjet/pkg/controller"
	"github.com/upbound/upjet/pkg/resource/json"
	"github.com/upbound/upjet/pkg/terraform"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/provider-aws/apis/ssm/v1beta1"
)

// Setup adds a controller that reconciles Parameter managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1beta1.Parameter_GroupVersionKind.String())
	var initializers managed.InitializerChain
	for _, i := range o.Provider.Resources["aws_ssm_parameter"].InitializerFns {
		initializers = append(initializers, i(mgr.GetClient()))
	}
	initializers = append(initializers, managed.NewNameAsExternalName(mgr.GetClient()))
	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.SecretStoreConfigGVK != nil {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), *o.SecretStoreConfigGVK))
	}
	r := managed.NewReconciler(mgr,
		xpresource.ManagedKind(v1beta1.Parameter_GroupVersionKind),
		managed.WithExternalConnecter(&connector{
			connector: tjcontroller.NewConnector(mgr.GetClient(), o.WorkspaceStore, o.SetupFn, o.Provider.Resources["aws_ssm_parameter"],
				tjcontroller.WithCallbackProvider(tjcontroller.NewAPICallbacks(mgr, xpresource.ManagedKind(v1beta1.Parameter_GroupVersionKind))),
			),
		}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithFinalizer(terraform.NewWorkspaceFinalizer(o.WorkspaceStore, xpresource.NewAPIFinalizer(mgr.GetClient(), managed.FinalizerName))),
		managed.WithTimeout(3*time.Minute),
		managed.WithInitializers(initializers),
		managed.WithConnectionPublishers(cps...),
		managed.WithPollInterval(o.PollInterval),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.Parameter{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

type connector struct {
	connector *tjcontroller.Connector
}

func (c *connector) Connect(ctx context.Context, mg xpresource.Managed) (managed.ExternalClient, error) {
	ec, err := c.connector.Connect(ctx, mg)
	return &external{
		external: ec,
	}, errors.Wrap(err, "failed to get ExternalClient")
}

type external struct {
	external managed.ExternalClient
}

func (e *external) Observe(ctx context.Context, mg xpresource.Managed) (managed.ExternalObservation, error) {
	eo, err := e.external.Observe(ctx, mg)
	if err != nil {
		return eo, errors.Wrap(err, "failed to observe parameter")
	}
	cr, ok := mg.(*v1beta1.Parameter)
	if !ok {
		return managed.ExternalObservation{}, errors.New("not a Parameter")
	}
	// todo: reset status.atProvider.value if sensitive is set
	if cr.Spec.ForProvider.Sensitive != nil && !*cr.Spec.ForProvider.Sensitive {
		stateFile := filepath.Join(os.TempDir(), string(cr.GetUID()), "terraform.tfstate")
		buff, err := os.ReadFile(stateFile)
		if err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "failed to read the tfstate file")
		}
		s := &json.StateV4{}
		if err := json.JSParser.Unmarshal(buff, s); err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "failed to unmarshal tfstate file")
		}
		var attr map[string]any
		if err := json.JSParser.Unmarshal(s.GetAttributes(), &attr); err != nil {
			return managed.ExternalObservation{}, errors.Wrap(err, "failed to unmarshal attributes")
		}
		val := attr["value"].(string)
		cr.Status.AtProvider.Value = &val
	}
	return eo, nil
}

func (e *external) Create(ctx context.Context, mg xpresource.Managed) (managed.ExternalCreation, error) {
	return e.external.Create(ctx, mg)
}

func (e *external) Update(ctx context.Context, mg xpresource.Managed) (managed.ExternalUpdate, error) {
	return e.external.Update(ctx, mg)
}

func (e *external) Delete(ctx context.Context, mg xpresource.Managed) error {
	return e.external.Delete(ctx, mg)
}
