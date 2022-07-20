package pack1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnStr(t *testing.T) {
	assert.Equal(t, "Hello main!", ReturnStr())
}
