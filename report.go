package logoped

// Report
type IReport interface {
	GetHeader() IHeader
	GetMessage() []byte
}

type Report struct {
	IReport

	header  IHeader
	message []byte
}

func CreateReport(header IHeader, message []byte) *Report {

	return &Report{
		header:  header,
		message: message,
	}
}

func (report Report) GetHeader() IHeader {

	return report.header
}

func (report Report) GetMessage() []byte {

	return report.message
}
