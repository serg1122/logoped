package logoped

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateRandomString(t *testing.T) {

	string, err := CreateRandomString([]rune("A"), 3)

	assert.Equal(t, `AAA`, string)
	assert.Nil(t, err)
}
