package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func init() {
	contentWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	file, err := os.OpenFile(
		"erp_backend.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(contentWriter, file)

	// defer file.Close()

	Logger = zerolog.New(multiWriter).With().Timestamp().Logger()

}
