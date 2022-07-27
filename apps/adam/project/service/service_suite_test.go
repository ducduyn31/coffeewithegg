package service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProjectModuleService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Project Service Suite")
}
