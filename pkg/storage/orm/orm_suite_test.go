package orm_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Orm Suite")
}
