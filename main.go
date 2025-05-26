/*
Pay attention in order to run this script first
you need to give this script a module name
the format is {REMOTE}/{USERNAME}/{module name}
we give this module a name with:
$ go mod init {REMOTE}/{USERNAME}/{module name}
after that we need to download the specified modules
and list them in go.mod go dose that automaticlly by
$ go mod tidy
finaly we can run the program by
$ go run <gofile>
To generate a binary file
$ go build <gofile>
tip1 A package named main has an entrypoint at the main() function. A main package is compiled into an executable program.
tip2 A package by any other name is a library package. Libraries have no entry point.
tip3 Go programs are organized into packages. A package is a directory of Go code that's all compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package (directory).
tip4 The go install command compiles and installs a package or packages on your local machine for your personal usage. It installs the package's compiled binary in the GOBIN directory.
$ go install
*/
/*
Simple implrementation of a graphql server with gqlgen
*/
package main

// import packages
import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ANSI color codes for terminal output
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	// Text Colors
	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"
	// Background Colors
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

/*
SetupLogger initializes zerolog to write to both console and a file.
*/
func SetupLogger(logFilePath string, logLevel zerolog.Level) (zerolog.Logger, error) {
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return zerolog.Logger{}, err
	}

	// Console writer with human-friendly formatting
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Set global log level
	zerolog.SetGlobalLevel(logLevel)

	// Combine both writers
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Set global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Create the logger
	logger := zerolog.New(multi).With().Caller().Timestamp().Logger()

	// Set as the global logger
	log.Logger = logger

	return logger, nil
}

// The function that will be executed
func main() {
	logger, err := SetupLogger("app.log", zerolog.InfoLevel)
	startTime := time.Now() // Record start time
	if err != nil {
		panic(err)
	}
	logger.Info().Str("FunctionName:", "main").Msg(FgCyan + "Main function started" + Reset)
	defer func() {
		logger.Info().Str("FunctionName:", "main").TimeDiff("Duration (ms)", time.Now(), startTime).Msg(FgCyan + "Main function ended." + Reset)
	}()

}

/*

First initialize the repository with this command

$ go get github.com/99designs/gqlgen
$ go run github.com/99designs/gqlgen init

Write your schema in graph\schema.graphqls then run the following command to generate the code

$ go get github.com/99designs/gqlgen
$ go run github.com/99designs/gqlgen generate

The codes will be generated now implement the resolvers in graph/resolver.go

finally run the server with

$ go run main.go

*/