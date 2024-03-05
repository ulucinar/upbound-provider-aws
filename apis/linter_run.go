//go:build linter_run

// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Package apis contains Kubernetes API for the provider.
package apis

import "k8s.io/apimachinery/pkg/runtime"

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	panic(`Cannot be called in runtime. The provider should not have been built with the "linter_run" build constraint.`)
}
