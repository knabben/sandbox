package main

import (
	"github.com/knabben/tutorial-istio-sec/magefiles/writter"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
)

const (
	PATH_2 = "2-goclient/code"
)

type Demo2 mg.Namespace

// PrintCode presents the go client source code
func (Demo2) PrintCode() error {
	sh.RunV("batcat", PATH_2+"/main.go")
	return nil
}

// Run finds the configmap in the cluster
func (Demo2) Run() error {
	os.Chdir(PATH_2)
	sh.RunV("go", "run", "main.go")

	writter.Kubectl("get", "cm", "-n", "kube-system")
	return nil
}
