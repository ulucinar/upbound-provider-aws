package v1beta1

import (
	"encoding/json"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/crossplane/crossplane-runtime/pkg/errors"

	"github.com/upbound/provider-aws/apis/elasticache/v1"
)

func (g *ReplicationGroup) ConvertTo(dst conversion.Hub) error {
	to, ok := dst.(*v1.ReplicationGroup)
	if !ok {
		return errors.Errorf("invalid destination type: %T", dst)
	}
	bs, err := json.Marshal(g)
	if err != nil {
		return errors.Wrap(err, "cannot marshal JSON data")
	}
	if err := json.Unmarshal(bs, dst); err != nil {
		return errors.Wrap(err, "cannot unmarshal JSON data")
	}
	to.APIVersion = v1.CRDGroupVersion.String()

	// join v1beta1 AvailabilityZones and PreferredCacheClusterAzs into v1 PreferredCacheClusterAzs
	to.Spec.ForProvider.PreferredCacheClusterAzs = append(g.Spec.ForProvider.AvailabilityZones, g.Spec.ForProvider.PreferredCacheClusterAzs...)
	bs, _ = json.Marshal(g.Spec.ForProvider.AvailabilityZones)
	to.Annotations["conversion.aws.provider.upbound.io/availability-zones"] = string(bs)

	return nil
}

func (g *ReplicationGroup) ConvertFrom(src conversion.Hub) error {
	from, ok := src.(*v1.ReplicationGroup)
	if !ok {
		return errors.Errorf("invalid source type: %T", src)
	}
	bs, err := json.Marshal(src)
	if err != nil {
		return errors.Wrap(err, "cannot marshal JSON data")
	}
	if err := json.Unmarshal(bs, g); err != nil {
		return errors.Wrap(err, "cannot unmarshal JSON data")
	}
	g.APIVersion = CRDGroupVersion.String()

	// split v1 PreferredCacheClusterAzs into v1beta1 AvailabilityZones and PreferredCacheClusterAzs
	var azs []*string
	ann := from.Annotations["conversion.aws.provider.upbound.io/availability-zones"]
	if ann != "" {
		if err := json.Unmarshal([]byte(ann), &azs); err != nil {
			return errors.Wrap(err, "cannot unmarshal JSON data")
		}
		delete(g.Annotations, "conversion.aws.provider.upbound.io/availability-zones")
	}
	if isPrefix(azs, from.Spec.ForProvider.PreferredCacheClusterAzs) {
		g.Spec.ForProvider.AvailabilityZones = azs
		g.Spec.ForProvider.PreferredCacheClusterAzs = from.Spec.ForProvider.PreferredCacheClusterAzs[len(azs):]
	} else {
		g.Spec.ForProvider.AvailabilityZones = from.Spec.ForProvider.PreferredCacheClusterAzs
	}

	return nil
}

func isPrefix(prefix []*string, l []*string) bool {
	if len(prefix) > len(l) {
		return false
	}
	for i, v := range prefix {
		if v == l[i] {
			continue
		}
		if v == nil || l[i] == nil || *v != *l[i] {
			return false
		}
	}
	return true
}
