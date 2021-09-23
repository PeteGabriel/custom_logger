package pkg

//Level definition to represent the severity level of a log entry.
type Level int8


const (
	Trace Level = iota
	Debug
	Info 
	Warn
	Error
	Fatal
)

//ToString obtains a string representation of a given level
func (l Level) ToString() string {
	switch l{
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"	
	case Info:
		return "INFO"
	case Warn:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return ""
	}
}

