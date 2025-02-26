package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type contextKey string

const (
	LogFieldsContextKey = contextKey("log_fields")

	ProjectDirectoryName = "mallbots"
	ModuleName           = "github.com/Chengxufeng1994/event-driven-arch-in-go"

	// durationStr is the suffix for the field holding a Duration as a
	// string.
	durationStr = "_str"
)

// log_fields keys
const (
	// RepositoryFieldKey repository name (string)
	RepositoryFieldKey = "repository"
	// MatchedHostFieldKey matched host (bool) true when domain extracted from host
	MatchedHostFieldKey = "matched_host"
	// RefHostFieldKey reference id (string)
	RefHostFieldKey = "ref"
	// PathFieldKey path / request URI (string)
	PathFieldKey = "path"
	// UploadIDFieldKey s3 multipart upload ID (string) "upload_id"
	UploadIDFieldKey = "upload_id"
	// ListTypeFieldKey s3 list type version (string, ex: v1 or v2)
	ListTypeFieldKey = "list_type"
	// PhysicalAddressFieldKey object physical address (string)
	PhysicalAddressFieldKey = "physical_address"
	// PartNumberFieldKey s3 multipart upload part number (string)
	PartNumberFieldKey = "part_number"
	// RequestIDFieldKey request ID (string) based on the request ID found on context
	RequestIDFieldKey = "request_id"
	// HostFieldKey request's host (string)
	HostFieldKey = "host"
	// MethodFieldKey request's method (string)
	MethodFieldKey = "method"
	// UserFieldKey user's name associated with the request (string)
	UserFieldKey = "user"
	// ServiceNameFieldKey service name (string, ex: rest_api)
	ServiceNameFieldKey = "service_name"
	// LogAudit kind of information to audit (string, ex: API)
	LogAudit = "log_audit"
)

var (
	formatterInitOnce sync.Once
	defaultLogger     = logrus.New()
	openLoggers       []io.Closer
	syslogOnce        sync.Once
)

func Level() string {
	return defaultLogger.GetLevel().String()
}

type Fields map[string]interface{}

// logCallerTrimmer is used to trim the caller paths to be relative to the project root
func logCallerTrimmer(frame *runtime.Frame) (function string, file string) {
	indexOfModule := strings.Index(strings.ToLower(frame.File), ProjectDirectoryName)
	if indexOfModule != -1 {
		file = frame.File[indexOfModule+len(ProjectDirectoryName):]
	} else {
		file = frame.File
	}
	file = fmt.Sprintf("%s:%d", strings.TrimPrefix(file, string(os.PathSeparator)), frame.Line)
	function = strings.TrimPrefix(frame.Function, fmt.Sprintf("%s%s", ModuleName, string(os.PathSeparator)))
	return
}

func SetLevel(level string) {
	switch strings.ToLower(level) {
	case "trace":
		defaultLogger.SetLevel(logrus.TraceLevel)
	case "debug":
		defaultLogger.SetLevel(logrus.DebugLevel)
	case "info":
		defaultLogger.SetLevel(logrus.InfoLevel)
	case "warn", "warning":
		defaultLogger.SetLevel(logrus.WarnLevel)
	case "error":
		defaultLogger.SetLevel(logrus.ErrorLevel)
	case "panic":
		defaultLogger.SetLevel(logrus.PanicLevel)
	case "null", "none":
		defaultLogger.SetLevel(logrus.PanicLevel)
		defaultLogger.SetOutput(io.Discard)
	}
}

func CloseWriters() error {
	for _, c := range openLoggers {
		if err := c.Close(); err != nil {
			return fmt.Errorf("close log writer: %w", err)
		}
	}
	openLoggers = nil
	return nil
}

func SetOutputs(outputs []string, fileMaxSizeMB, filesKeep int) error {
	var writers []io.Writer
	if err := CloseWriters(); err != nil {
		return fmt.Errorf("close previous log writers: %w", err)
	}
	for _, output := range outputs {
		var w io.Writer
		switch output {
		case "":
			continue
		case "-":
			w = os.Stdout
		case "=":
			w = os.Stderr
		default:
			l := &lumberjack.Logger{
				Filename:   output,
				MaxSize:    fileMaxSizeMB,
				MaxBackups: filesKeep,
			}
			w = l
			openLoggers = append(openLoggers, l)
		}
		writers = append(writers, w)
	}
	if len(writers) == 1 {
		defaultLogger.SetOutput(writers[0])
	} else if len(writers) > 1 {
		defaultLogger.SetOutput(io.MultiWriter(writers...))
	}
	return nil
}

func HasLogFileOutput(outputs []string) bool {
	return slices.ContainsFunc(outputs, func(e string) bool {
		return e != "" && e != "-" && e != "="
	})
}

func GetLogFileOutputPath(outputs []string) string {
	outFileIdx := slices.IndexFunc(outputs, func(e string) bool {
		return e != "" && e != "-" && e != "="
	})
	return outputs[outFileIdx]
}

type OutputFormatOptions struct {
	CallerPrettyfier func(*runtime.Frame) (function string, file string)
}

type OutputFormatOptionFunc func(options *OutputFormatOptions)

func SetOutputFormat(format string, opts ...OutputFormatOptionFunc) {
	// setup options
	var options OutputFormatOptions
	for _, opt := range opts {
		opt(&options)
	}
	if options.CallerPrettyfier == nil {
		options.CallerPrettyfier = logCallerTrimmer
	}

	// setup formatter
	var formatter logrus.Formatter
	switch strings.ToLower(format) {
	case "text":
		disableColors := false
		noColor := os.Getenv("NO_COLOR")
		if noColor != "" && noColor != "0" {
			disableColors = true
		}

		formatter = &logrus.TextFormatter{
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			PadLevelText:           true,
			QuoteEmptyFields:       true,
			CallerPrettyfier:       options.CallerPrettyfier,
			DisableColors:          disableColors,
		}
	case "json":
		formatter = &logrus.JSONFormatter{
			CallerPrettyfier: options.CallerPrettyfier,
			PrettyPrint:      false,
		}
	default:
		return // no known formatter found
	}

	// wrap it with our caller formatter
	defaultLogger.SetFormatter(logrusCallerFormatter{formatter})
}

type Logger interface {
	WithContext(ctx context.Context) Logger
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Log(level logrus.Level, args ...interface{})
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Logf(level logrus.Level, format string, args ...interface{})
	IsTracing() bool
	IsDebugging() bool
	IsInfo() bool
	IsError() bool
	IsWarn() bool
}

type logrusEntryWrapper struct {
	e *logrus.Entry
}

func (l *logrusEntryWrapper) WithContext(ctx context.Context) Logger {
	return addFromContext(
		&logrusEntryWrapper{l.e.WithContext(ctx)},
		ctx,
	)
}

var durationType = reflect.TypeOf(time.Duration(0))

// splitDurationFields modifies fields to split every field of type
// time.Duration into 2 fields, one "_nsecs" and one "_str".
func (l *logrusEntryWrapper) WithFields(fields Fields) Logger {
	var durationKeys []string
	for key, value := range fields {
		if value != nil && reflect.TypeOf(value).AssignableTo(durationType) {
			durationKeys = append(durationKeys, key)
		}
	}

	for _, key := range durationKeys {
		duration := fields[key].(time.Duration)
		fields[key] = duration.Nanoseconds()
		fields[key+durationStr] = duration.String()
	}

	return &logrusEntryWrapper{l.e.WithFields(logrus.Fields(fields))}
}

func (l *logrusEntryWrapper) WithField(key string, value interface{}) Logger {
	return l.WithFields(Fields{key: value})
}

func (l *logrusEntryWrapper) WithError(err error) Logger {
	return &logrusEntryWrapper{l.e.WithError(err)}
}

func (l *logrusEntryWrapper) Trace(args ...interface{}) {
	l.e.Trace(args...)
}

func (l *logrusEntryWrapper) Debug(args ...interface{}) {
	l.e.Debug(args...)
}

func (l *logrusEntryWrapper) Info(args ...interface{}) {
	l.e.Info(args...)
}

func (l *logrusEntryWrapper) Warn(args ...interface{}) {
	l.e.Warn(args...)
}

func (l *logrusEntryWrapper) Warning(args ...interface{}) {
	l.e.Warning(args...)
}

func (l *logrusEntryWrapper) Error(args ...interface{}) {
	l.e.Error(args...)
}

func (l *logrusEntryWrapper) Fatal(args ...interface{}) {
	l.e.Fatal(args...)
}

func (l *logrusEntryWrapper) Panic(args ...interface{}) {
	l.e.Panic(args...)
}

func (l *logrusEntryWrapper) Log(level logrus.Level, args ...interface{}) {
	l.e.Log(level, args...)
}

func (l *logrusEntryWrapper) Tracef(format string, args ...interface{}) {
	l.e.Tracef(format, args...)
}

func (l *logrusEntryWrapper) Debugf(format string, args ...interface{}) {
	l.e.Debugf(format, args...)
}

func (l *logrusEntryWrapper) Infof(format string, args ...interface{}) {
	l.e.Infof(format, args...)
}

func (l *logrusEntryWrapper) Warnf(format string, args ...interface{}) {
	l.e.Warnf(format, args...)
}

func (l *logrusEntryWrapper) Warningf(format string, args ...interface{}) {
	l.e.Warningf(format, args...)
}

func (l *logrusEntryWrapper) Errorf(format string, args ...interface{}) {
	l.e.Errorf(format, args...)
}

func (l *logrusEntryWrapper) Fatalf(format string, args ...interface{}) {
	l.e.Fatalf(format, args...)
}

func (l *logrusEntryWrapper) Panicf(format string, args ...interface{}) {
	l.e.Panicf(format, args...)
}

func (l *logrusEntryWrapper) Logf(level logrus.Level, format string, args ...interface{}) {
	l.e.Logf(level, format, args...)
}

func (l *logrusEntryWrapper) IsTracing() bool {
	return l.e.Logger.IsLevelEnabled(logrus.TraceLevel)
}

func (l *logrusEntryWrapper) IsDebugging() bool {
	return l.e.Logger.IsLevelEnabled(logrus.DebugLevel)
}

func (l *logrusEntryWrapper) IsInfo() bool {
	return l.e.Logger.IsLevelEnabled(logrus.InfoLevel)
}

func (l *logrusEntryWrapper) IsError() bool {
	return l.e.Logger.IsLevelEnabled(logrus.ErrorLevel)
}

func (l *logrusEntryWrapper) IsWarn() bool {
	return l.e.Logger.IsLevelEnabled(logrus.WarnLevel)
}

type logrusCallerFormatter struct {
	f logrus.Formatter
}

func (lf logrusCallerFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Caller = getCaller()
	return lf.f.Format(e)
}

// ContextUnavailable returns a Logger when no context is available.  It
// should be used only in code during startup, teardown, or tests.  Prefer
// to use Default().
func ContextUnavailable() Logger {
	// wrap formatter with our own formatter that overrides caller
	formatterInitOnce.Do(func() {
		// defaultLogger.SetReportCaller(true)
		defaultLogger.SetNoLock()
		defaultLogger.Formatter = logrusCallerFormatter{defaultLogger.Formatter}
	})
	return &logrusEntryWrapper{
		e: logrus.NewEntry(defaultLogger),
	}
}

// GetFieldsFromContext returns the logging fields on ctx or nil.
func GetFieldsFromContext(ctx context.Context) Fields {
	fields := ctx.Value(LogFieldsContextKey)
	if fields == nil {
		return nil
	}
	return fields.(Fields)
}

func addFromContext(log Logger, ctx context.Context) Logger {
	loggerFields := GetFieldsFromContext(ctx)
	return log.WithFields(loggerFields)
}

// FromContext returns a Logger for reporting logs during ctx.  This logger
// will typically include request IDs from the context.
func FromContext(ctx context.Context) Logger {
	return addFromContext(ContextUnavailable(), ctx)
}

func AddFields(ctx context.Context, fields Fields) context.Context {
	ctxFields := ctx.Value(LogFieldsContextKey)
	loggerFields := Fields{}
	if ctxFields != nil {
		loggerFields = ctxFields.(Fields)
	}
	for k, v := range fields {
		loggerFields[k] = v
	}
	return context.WithValue(ctx, LogFieldsContextKey, loggerFields)
}
