package main

import (
	"github.com/knabben/showcase/pkg/showcase"
	"path"
)

func main() {
	crd := path.Join("./pkg/showcase/testdata", "demo.yaml")
	sc := showcase.NewShowcase(crd)
	_, err := sc.LoadFromAPI()
	if err != nil {
		panic(err)
	}
	if err = sc.Run(); err != nil {
		panic(err)
	}
}
