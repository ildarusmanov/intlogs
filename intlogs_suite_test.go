package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestIntlogs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intlogs Suite")
}
