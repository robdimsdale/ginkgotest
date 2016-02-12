package ginkgotest_test

import (
	"fmt"

	"github.com/robdimsdale/ginkgotest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ginkgotest", func() {
	BeforeEach(func() {
		fmt.Printf("BeforeEach\n")
	})

	AfterEach(func() {
		fmt.Printf("AfterEach\n")
	})

	It("Prints", func() {
		fmt.Printf("Inside It\n")
		Expect(ginkgotest.Hello()).To(Equal("hello"))
		fmt.Printf("Completed It\n")
	})
})
