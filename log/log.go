package log

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

// ----------------------------
const (
	LNONE = iota
	LFATAL
	LERROR
	LWARN
	LINFO
	LDEBUG
	LALL
)

// 返回包名/文件名:lineno:funcName:
func getFileInfo(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		panic("getFileInfo()")
	}
	lastIndexPos := strings.LastIndex(file, "/")
	secondLastIndexPos := strings.LastIndex(file[:lastIndexPos], "/")
	f := runtime.FuncForPC(pc)
	funcName := f.Name()
	return fmt.Sprintf(`%s:%d %s(): `, file[secondLastIndexPos+1:], line, funcName)
}

type Log struct {
	l     *log.Logger
	level int
}

func NewLog(out io.Writer, prefix string, level int) *Log {
	if len(prefix) != 0 {
		prefix = prefix + " "
	}
	l := log.New(out, prefix, log.LstdFlags)
	return &Log{l, level}
}

func (p *Log) Debug(v ...interface{}) {
	if p.level < LDEBUG {
		return
	}
	p.l.Output(2,
		"[DEBUG] "+getFileInfo(2)+
			fmt.Sprint(v...),
	)
}
func (p *Log) Debugf(f string, v ...interface{}) {
	if p.level < LDEBUG {
		return
	}
	p.l.Output(2,
		"[DEBUG] "+
			fmt.Sprintf(f, v...),
	)
}

func (p *Log) Info(v ...interface{}) {
	if p.level < LINFO {
		return
	}
	p.l.Output(2,
		"[INFO] "+
			fmt.Sprint(v...),
	)
}
func (p *Log) Infof(f string, v ...interface{}) {
	if p.level < LINFO {
		return
	}
	p.l.Output(2,
		"[INFO] "+
			fmt.Sprintf(f, v...),
	)
}

func (p *Log) Warn(v ...interface{}) {
	if p.level < LWARN {
		return
	}
	p.l.Output(2,
		"[WARN] "+
			fmt.Sprint(v...),
	)
}
func (p *Log) Warnf(f string, v ...interface{}) {
	if p.level < LWARN {
		return
	}
	p.l.Output(2,
		"[WARN] "+
			fmt.Sprintf(f, v...),
	)
}

func (p *Log) Error(v ...interface{}) {
	if p.level < LERROR {
		return
	}
	p.l.Output(2,
		"[ERROR] "+
			fmt.Sprint(v...),
	)
}
func (p *Log) Errorf(f string, v ...interface{}) {
	if p.level < LERROR {
		return
	}
	p.l.Output(2,
		"[ERROR] "+
			fmt.Sprintf(f, v...),
	)
}
func (p *Log) Fatal(v ...interface{}) {
	if p.level < LFATAL {
		return
	}
	p.l.Output(2,
		"[FATAL] "+
			fmt.Sprint(v...),
	)
}
func (p *Log) Fatalf(f string, v ...interface{}) {
	if p.level < LFATAL {
		return
	}
	p.l.Output(2,
		"[FATAL] "+
			fmt.Sprintf(f, v...),
	)
}
