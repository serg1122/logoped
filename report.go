package logoped

// Report
type IReport interface {
	GetHeader() IHeader
	GetMessage() []byte
	GetPostMessage() []byte
}

type DefaultReport struct {
	IReport

	header      IHeader
	message     []byte
	postmessage []byte
}

func CreateDefaultReport(header IHeader, message []byte, postmessage []byte) *DefaultReport {

	return &DefaultReport{
		header:      header,
		message:     message,
		postmessage: postmessage,
	}
}

func (report DefaultReport) GetHeader() IHeader {

	return report.header
}

func (report DefaultReport) GetMessage() []byte {

	return report.message
}

func (report DefaultReport) GetPostMessage() []byte {

	return report.postmessage
}
