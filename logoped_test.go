package logoped

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var l1 = []byte{}
var l2 = []byte{}
var l3 = []byte{}

var l3message = []byte("все плохо")

type WriterToL1 struct {
	io.Writer
}

func (w WriterToL1) Write(bytes []byte) (int, error) {

	l1 = bytes

	return len(bytes), nil
}

type WriterToL2 struct {
	io.Writer
}

func (w WriterToL2) Write(bytes []byte) (int, error) {

	l2 = bytes

	return len(bytes), nil
}

type WriterToL3 struct {
	io.Writer
}

func (w WriterToL3) Write(bytes []byte) (int, error) {

	l3 = bytes

	return len(bytes), nil
}

type Formatter1 struct {
	IFormatter
}

func (f Formatter1) Format(report IReport) []byte {

	return append([]byte("prefix for formatter 1\n"), report.GetMessage()...)
}

type Formatter2 struct {
	IFormatter
}

func (f Formatter2) Format(report IReport) []byte {

	return append(report.GetMessage(), []byte("\npostfix for formatter 2")...)
}

type Formatter3 struct {
	IFormatter
}

func (f Formatter3) Format(report IReport) []byte {

	return l3message
}

func Test_CreateLogoped(t *testing.T) {

	logoped := CreateLogoped(
		CreateDispatcher(
			[]IAcceptor{CreateAcceptor(CreateFormatter(), LogLevelInfo, []io.Writer{os.Stderr})},
		),
		LogLevelInfo,
	)

	assert.IsType(t, &Logoped{}, logoped)
}

func Test_Usage(t *testing.T) {

	l1 = []byte{}
	l2 = []byte{}
	l3 = []byte{}

	formatter1 := &Formatter1{}
	formatter2 := &Formatter2{}
	formatter3 := &Formatter3{}

	logoped := CreateLogoped(
		CreateDispatcher(
			[]IAcceptor{
				CreateAcceptor(
					formatter1,
					LogLevelInfo,
					[]io.Writer{
						&WriterToL1{},
					},
				),
				CreateAcceptor(
					formatter2,
					LogLevelWarn,
					[]io.Writer{
						&WriterToL2{},
					},
				),
				CreateAcceptor(
					formatter3,
					LogLevelError,
					[]io.Writer{
						&WriterToL3{},
					},
				),
			},
		),
		LogLevelInfo,
	)

	message := []byte("Наш президент - вор")

	logoped.Warn(message)

	report := logoped.dispatcher.CreateReport(LogLevelWarn, message)

	expected1 := formatter1.Format(report)
	expected2 := formatter2.Format(report)
	expected3 := []byte{}

	assert.Equal(t, expected1, l1)
	assert.Equal(t, expected2, l2)
	assert.Equal(t, expected3, l3)
}
