package main_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntlogs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intlogs Suite")
}
