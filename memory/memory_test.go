package memory_test

import "testing"
import "github.com/onsi/gomega"
import "github.com/onsi/ginkgo"

func TestMemory(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Memory System")
}
