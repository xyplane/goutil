// Copyright 2011 Dylan Maxwell.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package log implements a logger that extends the 
// simple Go Logger to provide functionality for
// disabling output from the logger.  Several
// general purpose loggers are predefined and methods
// are provided for convenient access to them.
package log

import (
	"os"
	"io"
	"log"
)

var (
	info *Logger
	warn *Logger
	trace *Logger
	debug *Logger
	panik *Logger
	fatal *Logger
)

// Initialize the predefined loggers with appropriate prefixes. 
func init() {
	info = New(os.Stderr, "INFO: ", log.LstdFlags)
	warn = New(os.Stderr, "WARN: ", log.LstdFlags)
	trace = New(os.Stderr, "TRACE: ", log.LstdFlags)
	debug = New(os.Stderr, "DEBUG: ", log.LstdFlags)
	panik = New(os.Stderr, "PANIC: ", log.LstdFlags)
	fatal = New(os.Stderr, "FATAL: ", log.LstdFlags)
}

// SetOutput changes the output destination for the predefined loggers.
func SetOutput(w io.Writer) {
	info.SetOutput(w)
	warn.SetOutput(w)
	trace.SetOutput(w)
	debug.SetOutput(w)
	panik.SetOutput(w)
	fatal.SetOutput(w)
}

// Logger delegates to a simple Go Logger,
// but provides functionality to disable output. 
type Logger struct {
	enabled bool
	logger *log.Logger
}

// New create a new Logger with given writer, prefix and flags.
func New(w io.Writer, prefix string, flags int) *Logger {
	logger := log.New(w, prefix, flags)
	return &Logger{ enabled:true, logger:logger }
}

// SetOutput changes the output destination writer of the logger.
func (l *Logger) SetOutput(w io.Writer) {
	flags := l.logger.Flags()
	prefix := l.logger.Prefix()
	l.logger = log.New(w, prefix, flags)
}

// Disable or enable the logger.
func (l *Logger) setEnabled(enabled bool) {
	l.enabled = enabled
}

// Print calls the Print function of the underlying
// simple Go Logger if the logger is enabled. Returns
// true if the logger is enabled and false otherwise.
func (l *Logger) Print(v ...interface{}) bool {
	if l.enabled {
		if len(v) > 0 {
			l.logger.Print(v...)
		}
		return true
	}
	return false
}

// Println calls the Println function of the underlying
// simple Go Logger if the logger is enabled.
func (l *Logger) Println(v ...interface{}) {
	if l.enabled {
		l.logger.Println(v...)
	}
}

// Printf calls the Printf function of the underlying
// simple Go Logger if the logger is enabled.
func (l *Logger) Printf(format string, v ...interface{}) {
	if l.enabled {
		l.logger.Printf(format, v...)
	}
}

// Panic calls the Panic function of the underlying
// simple Go Logger irregardless if the logger is enabled.
func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panic(v...)
}

// Panicln calls the Panicln function of the underlying
// simple Go Logger irregardless if the logger is enabled.
func (l *Logger) Panicln(v ...interface{}) {
	l.logger.Panicln(v...)
}

// Panicf calls the Panicf function of the underlying
// simple Go Logger irregardless if the logger is enabled.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

// Fatal calls the Fatal function of the underlying
// simple Go Logger irregardless if the logger is enabled.
func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal(v...)
}

// Fatalln calls the Fatalln function of the underlying
// simple Go Logger irregardless if the logger is enabled.
func (l *Logger) Fatalln(v ...interface{}) {
	l.logger.Fatalln(v...)
}

// Fatalf calls the Fatalf function of the underlying
// simple Go Logger irregardless if the logger is enabled.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}

// SetInfo disables or enables the predefined Info Logger.
func SetInfo(enabled bool) {
	info.setEnabled(enabled)
}	

// Info calls the Print function of the predefined 'info' logger.
func Info(v ...interface{}) bool {
	return info.Print(v...)
}

// Infoln calls the Println function of the predefined 'info' logger.
func Infoln(v ...interface{}) {
	info.Println(v...)
}

// Infof calls the Printf function of the predefined 'info' logger.
func Infof(format string, v ...interface{}) {
	info.Printf(format, v...)
}

// SetWarn disables or enables the predefined 'warn' logger.
func SetWarn(enabled bool) {
	warn.setEnabled(enabled)
}

// Warn calls the Print function of the predefined 'warn' logger.
func Warn(v ...interface{}) bool {
	return warn.Print(v...)
}

// Warnln calls the Println function of the predefined 'warn' logger.
func Warnln(v ...interface{}) {
	warn.Println(v...)
}

// Warnf calls the Printf function of the predefined 'warn' logger.
func Warnf(format string, v ...interface{}) {
	warn.Printf(format, v...)
}

// SetTrace disables or enables the predefined 'trace' logger.
func SetTrace(enabled bool) {
	trace.setEnabled(enabled)
}

// Trace calls the Print function of the predefined 'trace' logger.
func Trace(v ...interface{}) bool {
	return trace.Print(v...)
}

// Traceln calls the Println function of the predefined 'trace' logger.
func Traceln(v ...interface{}) {
	trace.Println(v...)
}

// Tracef calls the Printf function of the predefined 'trace' logger.
func Tracef(format string, v ...interface{}) {
	trace.Printf(format, v...)
}

// SetDebug disables or enables the predefined 'debug' logger.
func SetDebug(enabled bool) {
	debug.setEnabled(enabled)
}

// Debug calls the Print function of the predefined 'debug' logger.
func Debug(v ...interface{}) bool {
	return debug.Print(v...)
}

// Debugln calls the Println function of the predefined 'debug' logger.
func Debugln(v ...interface{}) {
	debug.Println(v...)
}

// Debugf calls the Printf function of the predefined 'debug' logger.
func Debugf(format string, v ...interface{}) {
	debug.Printf(format, v...)
}

// Panic calls the Panic function of the predefined 'panic' logger.
func Panic(v ...interface{}) {
	panik.Panic(v...)
}

// Panicln calls the Panicln function of the predefined 'panic' logger.
func Panicln(v ...interface{}) {
	panik.Panicln(v...)
}

// Panicf calls the Panicf function of the predefined 'panic' logger.
func Panicf(format string, v ...interface{}) {
	panik.Panicf(format, v...)
}

// Fatal calls the Fatal function of the predefined 'fatal' logger.
func Fatal(v ...interface{}) {
	fatal.Fatal(v...)
}

// Fatalln calls the Fatalln function of the predefined 'fatal' logger.
func Fatalln(v ...interface{}) {
	fatal.Fatalln(v...)
}

// Fatalf calls the Fatalf function of the predefined 'fatal' logger.
func Fatalf(format string, v ...interface{}) {
	fatal.Fatalf(format, v...)
}
