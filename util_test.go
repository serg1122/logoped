package logoped

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateRandomString(t *testing.T) {

	str, err := CreateRandomString([]rune("A"), 3)

	assert.Equal(t, `AAA`, str)
	assert.Nil(t, err)
}
