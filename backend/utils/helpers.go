package utils

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// The function "PrintErrOnCommandLine" prints an error message to the command line if an error is not
// nil.
func PrintErrOnCommandLine(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// The function `WriteMessageToLogFile` writes a formatted message with timestamp, file name, line
// number, and function name to a log file.
func WriteMessageToLogFile(message interface{}) {
	now := time.Now()
	formatTime := now.Format(time.UnixDate)
	file, line, functionName := trace()
	filename := strings.Split(file, "/")
	stringMessage := fmt.Sprintf("%s on line %s in %s at file %s", message, strconv.Itoa(line), functionName, filename[len(filename)-1])
	MessageWithFormatTime := formatTime + ": " + stringMessage + "\n"
	WriteToLogFile(MessageWithFormatTime)
}

// The function "HandleError" logs an error message along with the current time, file name, line
// number, and function name if an error occurs.
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

// The function `WriteToLogFile` appends a given message to a log file, creating the file if it doesn't
// exist.
func WriteToLogFile(message string) {
	_, err := os.Stat("./logfile.txt")
	if err != nil {
		os.Create("./logfile.txt")
	}

	file, err := os.OpenFile("./logfile.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	PrintErrOnCommandLine(err)

	n, err := file.Write([]byte(message))
	PrintErrOnCommandLine(err)
	if n != len(message) {
		fmt.Println("message length not the same")
	}

}

// The trace function returns the file name, line number, and function name of the caller.
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

// The AssertString function takes an interface{} value and asserts that it is a string, returning the
// string value.
func AssertString(val interface{}) string {
	v := val.(string)
	return v
}

// generates a new UUID
func GenerateNewUUID() string {
	newUUID := uuid.New()

	// Convert the UUID to a string for display
	uuidString := newUUID.String()

	return uuidString
}
