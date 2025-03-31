package cli

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCli(t *testing.T) {
	t.Run("Parameters have been provided", func(t *testing.T) {
		os.Args = []string{"./bin/wc", "-c", "test.txt"}

		actual := HasProvided()

		assert.Truef(t, actual, "expected %t, got %t", true, actual)
	})

	t.Run("Parameters have not been provided", func(t *testing.T) {
		os.Args = []string{"./bin/wc"}

		actual := HasProvided()

		assert.Falsef(t, actual, "expected %t, got %t", false, actual)
	})

	t.Run("Get filename from parameter provided", func(t *testing.T) {
		os.Args = []string{"./bin/wc", "-c", "test.txt"}

		actual := NewCli()

		expected := "test.txt"
		assert.Equal(t, expected, actual.FileName, "expected %t, got %t", expected, actual)
	})
}
