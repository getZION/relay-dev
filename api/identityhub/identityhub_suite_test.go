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
