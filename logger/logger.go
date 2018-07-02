package logger

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/felixge/httpsnoop"
	"github.com/julienschmidt/httprouter"
)

type Logger struct {
	Info  *log.Logger
	Debug *log.Logger
	Err   *log.Logger
	resp  *log.Logger
}

func NewLogger(test ...bool) *Logger {
	if test != nil {
		i := log.New(ioutil.Discard, "", 0)
		return &Logger{i, i, i, i}
	}

	i := log.New(os.Stdout,
		"\033[1;32m[Info]:\033[0m ",
		log.Ldate|log.Ltime)

	d := log.New(os.Stdout,
		"\033[1;34m[Debug]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)

	e := log.New(os.Stderr,
		"\033[1;31m[Error]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)

	r := log.New(os.Stdout,
		"\033[1;37m[Response]:\033[0m ",
		log.Ldate|log.Ltime|log.Lshortfile)

	return &Logger{i, d, e, r}
}

func ResponseLogger(mux *httprouter.Router) http.HandlerFunc {
	l := NewLogger(true)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(mux, w, r)
		l.resp.Printf("%s %s | %d \n", r.Method, r.URL, m.Code)
	})
}
