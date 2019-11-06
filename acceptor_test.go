package logoped

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateAcceptor(t *testing.T) {

	acceptor := CreateAcceptor(CreateDefaultFormatter(), LogLevelInfo, []io.Writer{os.Stderr})

	assert.IsType(t, &Acceptor{}, acceptor)
}
