package api

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/knabben/showcase/api/v1alpha1"
)

var (
	scheme = runtime.NewScheme()
	codecs = serializer.NewCodecFactory(scheme, serializer.EnableStrict)
)

func init() {
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
}

// LoadYAML decode the input read YAML into a configuration object
func LoadYAML(data []byte) (*v1alpha1.Demo, error) {
	var deserializer = codecs.UniversalDeserializer()
	obj, gvk, err := deserializer.Decode(data, nil, nil)
	if err != nil {
		return nil, err
	}
	demo, ok := obj.(*v1alpha1.Demo)
	if !ok {
		return nil, fmt.Errorf("got unexpected config type: %v", gvk)
	}
	return demo, nil
}
