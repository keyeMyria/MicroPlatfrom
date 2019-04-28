package log

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)
// 对外接口
type os struct {
	*logger
	opts Options
}

type logger struct {
	f  Fields
	fn logFunc
}

type logFunc func(l Level, f Fields, m string) error

func newOS(opts ...Option) Log {
	options := Options{
		Level:   DefaultLevel,
		Fields:  make(Fields),
		Context: context.TODO(),
	}

	for _, o := range opts {
		o(&options)
	}

	if len(options.Outputs) == 0 {
		options.Outputs = []Output{NewOutput()}
	}

	//https://blog.csdn.net/linuxandroidwince/article/details/79113398...
	o := &os{
		&logger{
			f: make(Fields),
		},
		options,
	}

	// so ugly
	o.logger.fn = o.log

	return o
}

func (o *os) log(l Level, f Fields, m string) error {
	// discard if we're not at the right level
	if l < o.opts.Level {
		return nil
	}

	fields := make(Fields)

	for k, v := range o.opts.Fields {
		fields[k] = v
	}

	for k, v := range f {
		fields[k] = v
	}

	e := &Event{
		Timestamp: time.Now().Format("2006-1-2 15:04:05"),
		Level:     l,
		Fields:    fields,
		Message:   m,
	}



	var gerr error
	for i, v := range o.opts.Outputs {
		// sjlin 新增判断文件大小的功能， 如果超过指定大小，则替换原来的output
		if v.IsFull(){
			//log1 := mylog.NewLog(mylog.WithOutput(myfile.NewOutput(mylog.OutputName("sjlin.json"))))
			//file.NewOutput().
			newOutput := NewOutput(OutputName(v.String()))

			o.opts.Outputs[i] = newOutput
		}
		v = o.opts.Outputs[i]
		if err := v.Send(e); err != nil {
			gerr = err
		}
	}
	return gerr
}

func (o *os) Init(opts ...Option) error {
	for _, opt := range opts {
		opt(&o.opts)
	}
	return nil
}

func (o *os) Options() Options {
	return o.opts
}

func (o *os) String() string {
	return "os"
}

func (l *logger) Debug(args ...interface{}) {
	l.fn(DebugLevel, l.f, fmt.Sprint(args...))
}

func (l *logger) Info(args ...interface{}) {
	l.fn(InfoLevel, l.f, fmt.Sprint(args...))
}
// sjlin 新增漏的方法
func (l *logger) Warn(args ...interface{}) {
	l.fn(WarnLevel, l.f, fmt.Sprint(args...))
}
func (l *logger) Error(args ...interface{}) {
	l.fn(ErrorLevel, l.f, fmt.Sprint(args...))
}

func (l *logger) Fatal(args ...interface{}) {
	l.fn(FatalLevel, l.f, fmt.Sprint(args...))
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.fn(DebugLevel, l.f, fmt.Sprintf(format, args...))
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.fn(InfoLevel, l.f, fmt.Sprintf(format, args...))
}

// sjlin 新增漏的方法
func (l *logger) Warnf(format string, args ...interface{}) {
	l.fn(WarnLevel, l.f, fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.fn(ErrorLevel, l.f, fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.fn(FatalLevel, l.f, fmt.Sprintf(format, args...))
}

func (l *logger) Log(level Level, args ...interface{}) {
	l.fn(level, l.f, fmt.Sprint(args...))
}

func (l *logger) Logf(level Level, format string, args ...interface{}) {
	l.fn(level, l.f, fmt.Sprintf(format, args...))
}

func (l *logger) WithFields(f Fields) Logger {
	fields := make(Fields)

	for k, v := range l.f {
		fields[k] = v
	}

	for k, v := range f {
		fields[k] = v
	}

	return &logger{
		f:  fields,
		fn: l.fn,
	}
}
