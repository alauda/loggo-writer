package loggo_writer

import (
	"bufio"
	"fmt"
	"github.com/juju/loggo"
	"github.com/lestrrat/go-file-rotatelogs"
	"time"
)

type FileWriterConfig struct {
	LogDir       string
	AppName      string
	Formatter    func(entry loggo.Entry) string
	MaxAge       time.Duration
	RotationTime time.Duration
}

type FileWriter struct {
	config    *FileWriterConfig
	writer    *bufio.Writer
	rotatelog *rotatelogs.RotateLogs
}

func NewFileWriter(conf *FileWriterConfig) (*FileWriter, error) {
	var err error
	fullPath := fmt.Sprintf("%s/%s.log.%s", conf.LogDir, conf.AppName, "%Y%m%d%H%M")
	if err = EnsureDir(fullPath); err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/%s.log", conf.LogDir, conf.AppName)

	if conf.MaxAge <= 0 {
		conf.MaxAge = 7 * 24 * time.Hour
	}

	if conf.RotationTime <= 0 {
		conf.RotationTime = 24 * time.Hour
	}

	rl, err := rotatelogs.New(fullPath,
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(conf.MaxAge),
		rotatelogs.WithRotationTime(conf.RotationTime),
	)

	if err != nil {
		return nil, err
	}

	writer := bufio.NewWriter(rl)

	return &FileWriter{
		config:    conf,
		rotatelog: rl,
		writer:    writer,
	}, nil
}

func (l *FileWriter) Close() {
	l.rotatelog.Close()
}

func (l *FileWriter) Write(entry loggo.Entry) {
	logLine := l.config.Formatter(entry)
	l.writer.WriteString(logLine + "\n")
	l.writer.Flush()
}
