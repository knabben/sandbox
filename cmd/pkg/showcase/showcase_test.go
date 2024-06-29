package showcase

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

var fakePath = "./testdata"

func TestLoadAPI(t *testing.T) {
	sc := NewShowcase(path.Join(fakePath, "demo.yaml"))
	demo, err := sc.LoadFromAPI()
	assert.Nil(t, err)
	assert.Len(t, demo.Spec.Steps, 2)

	for _, step := range demo.Spec.Steps {
		assert.Greater(t, len(step.Description), 0)
		assert.Greater(t, len(step.Command), 1)
	}
}
