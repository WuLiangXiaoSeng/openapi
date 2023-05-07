package common

import (
	"flag"
	"fmt"
	"path"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

var debug_on bool = false

func SetDebugOn() {
	debug_on = true
}

func SetDebugOff() {
	debug_on = false
}

func GetDebugOn() bool {
	return debug_on
}

func Println(a ...interface{}) {
	if GetDebugOn() {
		if pc, file, line, ok := runtime.Caller(1); ok {
			funcName := runtime.FuncForPC(pc).Name()
			fileName := path.Base(file)
			funcNameNumber := strings.LastIndex(funcName, ".")
			fileNameNumber := strings.LastIndex(fileName, ".")
			if funcNameNumber != -1 {
				fmt.Printf("[%s:%s:%d] ", fileName[:fileNameNumber], funcName[funcNameNumber+1:], line)
			} else {
				fmt.Printf("[%s:%s:%d] ", fileName[:fileNameNumber], funcName, line)
			}
		}
		fmt.Println(a...)
	}
}

func Printf(format string, a ...interface{}) {
	if GetDebugOn() {
		if pc, file, line, ok := runtime.Caller(1); ok {
			funcName := runtime.FuncForPC(pc).Name()
			fileName := path.Base(file)
			matchNumber := strings.LastIndex(funcName, ".")
			fileNameNumber := strings.LastIndex(fileName, ".")
			if matchNumber != -1 {
				fmt.Printf("[%s:%s:%d] ", fileName[:fileNameNumber], funcName[matchNumber+1:], line)
			} else {
				fmt.Printf("[%s:%s:%d] ", fileName[:fileNameNumber], funcName, line)
			}
		}
		fmt.Printf(format, a...)
	}
}

func DumpStack() {
	debug.PrintStack()
}

func Debug() {
	var b = flag.Bool("debug", false, "")
	flag.Parse()
	if *b {
		SetDebugOn()
		gin.SetMode(gin.DebugMode)
	} else {
		SetDebugOff()
		gin.SetMode(gin.ReleaseMode)
	}
}
