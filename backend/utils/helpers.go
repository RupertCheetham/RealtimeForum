package utils

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func PrintErrOnCommandLine(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func WriteMessageToLogFile(message interface{}) {
	now := time.Now()
	formatTime := now.Format(time.UnixDate)
	file, line, functionName := trace()
	filename := strings.Split(file, "/")
	stringMessage := fmt.Sprintf("%s on line %s in %s at file %s", message, strconv.Itoa(line), functionName, filename[len(filename)-1])
	// stringMessage := string(AssertString(message)) + " on line " + strconv.Itoa(line) + " in " + functionName + " at file " + filename[len(filename)-1]
	MessageWithFormatTime := formatTime + ": " + stringMessage + "\n"
	WriteToLogFile(MessageWithFormatTime)
}

func HandleError(message string, err error) {
	if err != nil {
		now := time.Now()
		formatTime := now.Format(time.UnixDate)
		file, line, functionName := trace()
		filename := strings.Split(file, "/")
		errorMessage := fmt.Sprintf("***** %s: %v on line %s in %s at file %s *****", message, err, strconv.Itoa(line), functionName, filename[len(filename)-1])
		errorMessageWithFormatTime := formatTime + ": " + errorMessage + "\n"
		WriteToLogFile(errorMessageWithFormatTime)
	}
}

func WriteToLogFile(message string) {
	file, err := os.OpenFile("./logfile.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	PrintErrOnCommandLine(err)

	n, err := file.Write([]byte(message))
	PrintErrOnCommandLine(err)
	if n != len(message) {
		fmt.Println("message length not the same")
	}
}

func trace() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "?", 0, "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return file, line, "?"
	}

	return file, line, fn.Name()
}

func AssertString(val interface{}) string {
	v := val.(string)
	return v
}
