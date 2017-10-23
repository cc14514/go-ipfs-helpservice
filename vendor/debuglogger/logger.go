package debuglogger

import (
	"fmt"
	"runtime"
	"strconv"
)

func Printf(format string, msg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("["+file+":"+strconv.Itoa(line)+"] "+format+"\n", msg...)
}

func Println(msg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	h := "[" + file + ":" + strconv.Itoa(line) + "] "
	arr := []interface{}{h}
	arr = append(arr, msg...)
	fmt.Println(msg...)
}
