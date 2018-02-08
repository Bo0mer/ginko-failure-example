//+build integration

package integrationtest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test", func() {
	It("should pass", func() {
		Ω(true).Should(BeTrue())
	})
})
