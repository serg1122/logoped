package logoped

// Report
type IReport interface {
	GetHeader() IHeader
	GetMessage() []byte
}

type DefaultReport struct {
	IReport

	header  IHeader
	message []byte
}

func CreateDefaultReport(header IHeader, message []byte) *DefaultReport {

	return &DefaultReport{
		header:  header,
		message: message,
	}
}

func (report DefaultReport) GetHeader() IHeader {

	return report.header
}

func (report DefaultReport) GetMessage() []byte {

	return report.message
}
