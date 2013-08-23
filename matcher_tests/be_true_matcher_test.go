package matcher_tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/matchers"
)

var _ = Describe("BeTrue", func() {
	It("should handle true and false correctly", func() {
		Ω(true).Should(BeTrue())
		Ω(false).ShouldNot(BeTrue())
	})

	It("should only support booleans", func() {
		success, _, err := (&BeTrueMatcher{}).Match("foo")
		Ω(success).Should(BeFalse())
		Ω(err).Should(HaveOccured())
	})
})
