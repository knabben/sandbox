package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
)

const (
	PATH_3 = "3-kubebuilder/code"
)

type Demo3 mg.Namespace

// StartController initialize the controller and API
func (Demo3) StartController() error {
	sh.RunV("rm", "-fr", PATH_3)
	sh.RunV("mkdir", PATH_3)
	os.Chdir(PATH_3)
	sh.RunV("kubebuilder", "init", "--domain", "corp.beer", "--repo", "github.com/knabben/showcase")
	sh.RunV("kubebuilder", "create", "api", "--group", "showcase", "--version", "v1", "--kind", "Presentation")
	return nil
}