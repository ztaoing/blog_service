/**
* @Author:zhoutao
* @Date:2020/7/30 下午9:11
 */

package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

//日志分级
type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

//在不同的使用场景中记录不同级别的日志
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

//日志标准化
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context // 上下文
	level     Level           //级别
	fileds    Fields          //日志公共字段
	callers   []string        //设置当前某一层调用栈的信息（程序计数器、文件信息、行号）
}

func NewLlogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

//设置日志等级
func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

//设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fileds == nil {
		ll.fileds = make(Fields)
	}
	for k, v := range f {
		ll.fileds[k] = v
	}
	return ll

}

//设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

//设置当前都一层调用栈的信息（程序计数器、文件信息、行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	//
	pc, file, line, ok := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	if ok {
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}
	return ll
}

//设置当前整个调用栈的信息
func (l *Logger) WithCallerFrames() *Logger {
	//最大调用栈层数
	maxCallerDepth := 25
	minCallerDepth := 1

	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)

	depth := runtime.Callers(minCallerDepth, pcs)
	freams := runtime.CallersFrames(pcs[:depth])
	//?
	for frame, more := freams.Next(); more; frame, more = freams.Next() {
		s := fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

//日志的格式化和输出
func (l *Logger) JSONFormat(message string) map[string]interface{} {
	data := make(Fields, len(l.fileds)+4)
	data["level"] = l.level
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fileds) > 0 {
		for k, v := range l.fileds {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

//日志分级输出
func (l *Logger) Output(message string) {
	body, _ := json.Marshal(l.JSONFormat(message))
	content := string(body)

	switch l.level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Print(content)
	case LevelPanic:
		l.newLogger.Print(content)
	}
}

//日志输出的暴露方法
func (l *Logger) Debug(v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}
