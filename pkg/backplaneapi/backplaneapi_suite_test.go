package backplaneapi_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBackplaneapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backplaneapi Suite")
}
