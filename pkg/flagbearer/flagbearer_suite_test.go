package flagbearer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFlagbearer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flagbearer Suite")
}
