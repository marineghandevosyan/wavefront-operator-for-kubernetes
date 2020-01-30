// Copyright 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	"github.com/marineghandevosyan/wavefront-operator-for-kubernetes/pkg/controller/wavefrontcollector"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, wavefrontcollector.Add)
}
