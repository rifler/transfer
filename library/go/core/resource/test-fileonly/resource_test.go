package test

import (
	"os"
	"testing"

	"github.com/doublecloud/tross/library/go/core/resource"
	"github.com/stretchr/testify/assert"
)

func TestResource(t *testing.T) {
	assert.Equal(t, []byte("hello world"), resource.Get("/a.txt"))

	bindata, err := os.ReadFile("testdata/b.bin")
	assert.NoError(t, err)
	assert.Equal(t, bindata, resource.Get("/b.bin"))

	assert.Equal(t, []byte("handle this"), resource.Get("testdata/collision.txt"))
}
