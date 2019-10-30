package logoped

type ILogLevel interface {
	Name() string
	Weight() int
}

type LogLevel struct {
	name   string
	weight int

	ILogLevel
}

func CreateLogLevel(name string, weight int) *LogLevel {
	return &LogLevel{
		weight: weight,
		name:   name,
	}
}

func (ll LogLevel) Name() string {
	return ll.name
}

func (ll LogLevel) Weight() int {
	return ll.weight
}
