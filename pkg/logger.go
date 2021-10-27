package pkg

import (
	"encoding/json"
	"io"
	"runtime/debug"
	"sync"
	"time"
)

//Logger represents our custom logger type.
//It holds info about the output destination
//and also a mutex to coordinate writes.
type Logger struct {
	out      io.Writer
	minLevel Level
	mutx     sync.Mutex
}

type message struct {
	Level      string            `json:"level"`
	Time       string            `json:"time"`
	Message    string            `json:"message"`
	Properties map[string]string `json:"properties,omitempty"`
	Trace      string            `json:"trace,omitempty"`
}

// Return a new Logger instance which writes log entries
// at or above a minimum severity level to a specific output destination.
func New(o io.Writer, l Level) *Logger {
	return &Logger{
		out:      o,
		minLevel: l,
	}
}

func (l *Logger) print(lvl Level, msg string, prop map[string]string) (int, error) {
	//ignore below minimum level
	if lvl < l.minLevel {
		return 0, nil
	}

	m := message{
		Level:      lvl.ToString(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    msg,
		Properties: prop,
	}

	// Include a stack trace for entries at the ERROR and FATAL levels.
	if lvl >= Error {
		m.Trace = string(debug.Stack())
	}

	jMsg, err := json.Marshal(m)
	if err != nil {
		jMsg = []byte(Error.ToString() + ": unable to marshal log message:" + err.Error())
	}

	// Lock the mutex so that no two writes to the
	// output destination cannot happen concurrently.
	// It's possible that the text for two or more
	// log entries will be intermingled in the output.
	l.mutx.Lock()
	defer l.mutx.Unlock()

	return l.out.Write(append(jMsg, '\n'))
}

//Info logs a message and optional properties with a level of Info.
//Generally useful information to log.
func (l *Logger) Info(msg string, prop map[string]string) {
	l.print(Info, msg, prop)
}

//Error logs a message and optional properties with a level of Error.
//Includes stackstrace information.
func (l *Logger) Error(msg string, prop map[string]string) {
	l.print(Error, msg, prop)
}

//Fatal logs a message and optional properties with a level of Fatal.
//Program is shutdown afterwards.
func (l *Logger) Fatal(msg string, prop map[string]string) {
	l.print(Fatal, msg, prop)
}

//Warn logs a message and optional properties with a level of Warn.
func (l *Logger) Warn(msg string, prop map[string]string) {
	l.print(Warn, msg, prop)
}

//Debug logs a message and optional properties with a level of Debug.
//Use it for debug purposes.
func (l *Logger) Debug(msg string, prop map[string]string) {
	l.print(Debug, msg, prop)
}

//Trace logs a message and optional properties with a level of Trace.
func (l *Logger) Trace(msg string, prop map[string]string) {
	l.print(Trace, msg, prop)
}