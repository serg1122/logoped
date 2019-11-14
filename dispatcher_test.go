package logoped

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s1 = []byte{}

type WriterToS1 struct {
	io.Writer
}

func (w WriterToS1) Write(bytes []byte) (int, error) {

	s1 = bytes

	return len(bytes), nil
}

var s2 = []byte{}

type WriterToS2 struct {
	io.Writer
}

func (w WriterToS2) Write(bytes []byte) (int, error) {

	s2 = bytes

	return len(bytes), nil
}

var s3 = []byte{}

type WriterToS3 struct {
	io.Writer
}

func (w WriterToS3) Write(bytes []byte) (int, error) {

	s3 = bytes

	return len(bytes), nil
}

// Tests
func Test_CreateDefaultDispatcher(t *testing.T) {

	dispatcher := CreateDefaultDispatcher([]IAcceptor{CreateAcceptor(CreateDefaultFormatter(), LogLevelInfo, []io.Writer{os.Stderr})})

	assert.IsType(t, &DefaultDispatcher{}, dispatcher)
}

func Test_CreateReport(t *testing.T) {

	dispatcher := CreateDefaultDispatcher([]IAcceptor{CreateAcceptor(CreateDefaultFormatter(), LogLevelInfo, []io.Writer{os.Stderr})})

	report := dispatcher.CreateReport(LogLevelWarn, []byte{})

	assert.IsType(t, &DefaultReport{}, report)
}

func Test_Dispatch(t *testing.T) {

	dispatcher := CreateDefaultDispatcher(
		[]IAcceptor{
			CreateAcceptor(CreateDefaultFormatter(), LogLevelInfo, []io.Writer{&WriterToS1{}, &WriterToS2{}}),
			CreateAcceptor(CreateDefaultFormatter(), LogLevelInfo, []io.Writer{&WriterToS3{}}),
		},
	)

	report := dispatcher.CreateReport(LogLevelWarn, []byte("Эту фразу нужно првоерить"))

	dispatcher.Dispatch(report)

	expected := CreateDefaultFormatter().Format(report)
	assert.Equal(t, expected, s1)
	assert.Equal(t, expected, s2)
	assert.Equal(t, expected, s3)
}
