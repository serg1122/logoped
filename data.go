package logoped

import "time"

// Report
type IReport interface {
	GetHeader() IHeader
	GetMessage() []byte
	GetPostMessage() []byte
}

type Report struct {
	IReport

	header      IHeader
	message     []byte
	postmessage []byte
}

func CreateReport(header IHeader, message []byte, postmessage []byte) *Report {

	return &Report{
		header:      header,
		message:     message,
		postmessage: postmessage,
	}
}

func (report Report) GetHeader() IHeader {

	return report.header
}

func (report Report) GetMessage() []byte {

	return report.message
}

func (report Report) GetPostMessage() []byte {

	return report.postmessage
}

// Header
type IHeader interface {
	GetTimepoint() time.Time
	GetLogLevel() ILogLevel
}

type Header struct {
	timepoint time.Time
	loglevel  ILogLevel
}

func CreateHeader(timepoint time.Time, loglevel ILogLevel) *Header {

	return &Header{
		timepoint: timepoint,
		loglevel:  loglevel,
	}
}

func (header Header) GetTimepoint() time.Time {

	return header.timepoint
}

func (header Header) GetLogLevel() ILogLevel {

	return header.loglevel
}
