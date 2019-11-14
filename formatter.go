package logoped

type IFormatter interface {
	Format(IReport) []byte
}

type Formatter struct {
	delimiter       []byte
	timepointFormat string
}

func CreateFormatter() *Formatter {

	return &Formatter{

		delimiter:       []byte("\n"),
		timepointFormat: "2006-01-02 15:04:05.999999",
	}
}

func (formatter Formatter) Format(report IReport) []byte {

	header := report.GetHeader()

	formatted := []byte{}
	formatted = append(formatted, []byte(header.GetTimepoint().Format(formatter.timepointFormat)+" "+header.GetLogLevel().Name())...)
	formatted = append(formatted, formatter.delimiter...)
	formatted = append(formatted, report.GetMessage()...)

	return formatted
}
