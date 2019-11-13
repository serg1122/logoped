package logoped

import "time"

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
