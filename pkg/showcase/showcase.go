package showcase

import (
	"github.com/knabben/showcase/api/v1alpha1"
	"github.com/knabben/showcase/pkg/api"
	"github.com/knabben/showcase/pkg/ui"
	"os"
)

type Interface interface {
	LoadFromAPI() (*v1alpha1.Demo, error)
	Run() error
}

// Showcase holds the CRD and metadata for the demo
type Showcase struct {
	CRDPath string `json:"crd,omitempty"`
	Demo    *v1alpha1.Demo
}

// NewShowcase returns a new Showcase internal object
func NewShowcase(crdPath string) Interface {
	return &Showcase{CRDPath: crdPath}
}

// LoadFromAPI returns the decoded Demo specification
func (s *Showcase) LoadFromAPI() (demo *v1alpha1.Demo, err error) {
	var data []byte
	if data, err = os.ReadFile(s.CRDPath); err != nil {
		return nil, err
	}
	if demo, err = api.LoadYAML(data); err != nil {
		return nil, err
	}
	s.Demo = demo
	return
}

// Run present the final demo for the user
func (s *Showcase) Run() error {
	return ui.Run(s.Demo)
}
