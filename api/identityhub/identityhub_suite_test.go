package identityhub

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIdentityHub(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IdentityHub Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("IdentityHub Suite Started")
})

var _ = AfterSuite(func() {
	fmt.Println("IdentityHub Suite Finished")
})

type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
	Fail(fmt.Sprintf(format, args...))
}
