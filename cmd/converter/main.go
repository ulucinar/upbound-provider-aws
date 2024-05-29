// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"

	"github.com/crossplane/upjet/pkg/examples/conversion"

	"github.com/upbound/provider-aws/config"
)

func main() {
	pc, err := config.GetProvider(context.Background(), false)
	if err != nil {
		panic(err)
	}

	if err = conversion.ConvertSingletonListToEmbeddedObject(pc,
		"examples",
		"hack/boilerplate.examples.txt"); err != nil {
		panic(err)
	}
}
