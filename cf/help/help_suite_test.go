package help_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	"github.com/cloudfoundry/cli/cf/commands"
)

func TestHelp(t *testing.T) {
	RegisterFailHandler(Fail)

	commands.Load()

	RunSpecs(t, "Help Suite")
}
