package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunVersion(t *testing.T) {
	arg1 := ""
	arg2 := ""

	runVersion(func(a ...interface{}) (n int, err error) {
		arg1 = a[0].(string)
		arg2 = a[1].(string)
		return 0, nil
	})(nil, []string{})

	assert.Equal(t, "Version:", arg1)
	assert.Equal(t, "0.0.1", arg2)
}
