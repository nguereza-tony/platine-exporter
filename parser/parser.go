package parser

import (
	"bytes"
	"strconv"
)

type Entry struct {
	Method []byte
	Path   []byte
	Status int
	Time   float64
	DBTime float64
}

func Parse(line []byte, e *Entry) {
	fields := bytes.SplitSeq(line, []byte("\t"))

	for f := range fields {

		if bytes.HasPrefix(f, []byte("method=")) {
			e.Method = f[7:]
		}

		if bytes.HasPrefix(f, []byte("path=")) {
			e.Path = f[5:]
		}

		if bytes.HasPrefix(f, []byte("status=")) {
			e.Status, _ = strconv.Atoi(string(f[7:]))
		}

		if bytes.HasPrefix(f, []byte("time=")) {
			v := bytes.TrimSuffix(f[5:], []byte("ms"))
			ms, _ := strconv.ParseFloat(string(v), 64)
			e.Time = ms / 1000
		}

		if bytes.HasPrefix(f, []byte("db_time=")) {
			v := bytes.TrimSuffix(f[8:], []byte("ms"))
			ms, _ := strconv.ParseFloat(string(v), 64)
			e.DBTime = ms / 1000
		}
	}
}
