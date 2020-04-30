package geyser_test

import (
	"github.com/13k/geyser/v2"
)

type NoopLogger struct{}

var _ geyser.Logger = (*NoopLogger)(nil)

func (l *NoopLogger) Errorf(format string, args ...interface{}) {}
func (l *NoopLogger) Warnf(format string, args ...interface{})  {}
func (l *NoopLogger) Debugf(format string, args ...interface{}) {}
