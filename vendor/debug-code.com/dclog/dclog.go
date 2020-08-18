package dclog

import (
"fmt"
"runtime"
"strconv"
"time"
)

//type LogInfo

var fileFlag bool
var filePath string
var flag = "[DCLog]"
func InitSlog(flag bool, path string) {
	if flag {
		fileFlag = flag
		filePath = path
	}

}

func Info(arg ...interface{}) {

	msg := fmt.Sprint(arg...)
	timeS := time.Now().Format("2006-01-02 15:04:05 ")
	title := flag+" [INFO] " + timeS + " "

	_, file, line, ok := runtime.Caller(1)
	if ok {
		title += file + ":" + strconv.Itoa(line)
	}

	msg = title + "  " + msg
	fmt.Println(msg)

}

func Error(arg ...interface{}) {

	msg := fmt.Sprint(arg...)
	timeS := time.Now().Format("2006-01-02 15:04:05 ")
	title := flag+" [ERROR] " + timeS + " "

	_, file, line, ok := runtime.Caller(1)
	if ok {
		title += file + ":" + strconv.Itoa(line)
	}

	msg = title + "  " + msg
	fmt.Println(msg)

}

func Debug(arg ...interface{}) {

	msg := fmt.Sprint(arg...)
	timeS := time.Now().Format("2006-01-02 15:04:05 ")
	title := flag+" [DEBUG] " + timeS + " "

	_, file, line, ok := runtime.Caller(1)
	if ok {
		title += file + ":" + strconv.Itoa(line)
	}

	msg = title + "  " + msg
	fmt.Println(msg)

}
