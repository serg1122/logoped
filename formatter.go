package logoped

type IFormatter interface {
	Format(IReport) []byte
}

type DefaultFormatter struct {
	delimiter       []byte
	timepointFormat string
}

func CreateDefaultFormatter() *DefaultFormatter {

	return &DefaultFormatter{

		delimiter:       []byte("\n"),
		timepointFormat: "2006-01-02 15:04:05.999999",
	}
}

func (formatter DefaultFormatter) Format(report IReport) []byte {

	header := report.GetHeader()

	formatted := []byte{}
	formatted = append(formatted, []byte(header.GetTimepoint().Format(formatter.timepointFormat)+" "+header.GetLogLevel().Name())...)
	formatted = append(formatted, formatter.delimiter...)
	formatted = append(formatted, report.GetMessage()...)

	if len(report.GetPostMessage()) > 0 {
		formatted = append(formatted, formatter.delimiter...)
		formatted = append(formatted, report.GetPostMessage()...)
	}

	return formatted
}
