package imap

import "fmt"

type defaultLogger struct{}

func (l *defaultLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLogger) Logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
