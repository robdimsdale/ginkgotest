package ginkgotest_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinkgotest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgotest Suite")
}

var _ = BeforeSuite(func() {
	fmt.Printf("BeforeSuite\n")
})

var _ = AfterSuite(func() {
	fmt.Printf("AfterSuite\n")
})
