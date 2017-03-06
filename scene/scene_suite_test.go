package scene_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestScene(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scene Suite")
}
