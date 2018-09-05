package ginkgomod

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgomod(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgomod Suite")
}

var _ = Describe("Foo", func() {
	It("Bar", func() {
		Expect(10).To(Equal(10))
	})
})
