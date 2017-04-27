package main

import (
	"fmt"
	"github.com/alauda/loggo-writer"
	"github.com/juju/loggo"
	"path/filepath"
	"time"
)

var myLogger = loggo.GetLogger("client")

func myFormatter(entry loggo.Entry) string {
	ts := entry.Timestamp.In(time.UTC).Format("2006-01-02 15:04:05.999")
	// Just get the basename from the filename
	filename := filepath.Base(entry.Filename)
	return fmt.Sprintf("%s %s %s %s:%d %s", ts, entry.Level, entry.Module, filename, entry.Line, entry.Message)
}

func main() {
	conf := loggo_writer.FileWriterConfig{
		LogDir:       "/tmp",
		AppName:      "myapp",
		Formatter:    myFormatter,
		MaxAge:       24 * time.Hour,
		RotationTime: 1 * time.Minute,
	}

	writer, err := loggo_writer.NewFileWriter(&conf)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	loggo.RemoveWriter("default")
	loggo.RegisterWriter("default", writer)
	loggo.ConfigureLoggers("client=DEBUG")

	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		myLogger.Debugf("Hello alauda loggo filewriter %d", i)
	}
}
