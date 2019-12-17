package controller

import (
	"github.com/example-inc/apln-operator/pkg/controller/apln"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, apln.Add)
}
