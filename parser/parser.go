package parser

import (
	"bytes"
)

type LogEntry struct {
	Ip      []byte
	Request []byte
	Method  []byte
	Path    []byte
	User    *int
	Status  int
	Time    float64
	DBTime  float64
}

var idToken = []byte("{id}")

func Parse(line []byte, e *LogEntry) {

	start := 0
	n := len(line)

	for i := 0; i <= n; i++ {

		if i == n || line[i] == '\t' {

			field := line[start:i]
			parseField(field, e)

			start = i + 1
		}
	}
}

func parseField(f []byte, e *LogEntry) {
	switch {
	case bytes.HasPrefix(f, []byte("req=")):
		e.Request = f[4:]

	case bytes.HasPrefix(f, []byte("ip=")):
		e.Ip = f[3:]

	case bytes.HasPrefix(f, []byte("user=")):
		if isNumeric(f[5:]) {
			val := atoiBytes(f[5:])
			e.User = &val
		}

	case bytes.HasPrefix(f, []byte("method=")):
		e.Method = f[7:]

	case bytes.HasPrefix(f, []byte("path=")):
		e.Path = normalizePath(f[5:])

	case bytes.HasPrefix(f, []byte("status=")):
		e.Status = atoiBytes(f[7:])

	case bytes.HasPrefix(f, []byte("time=")):
		e.Time = parseMs(f[5:])

	case bytes.HasPrefix(f, []byte("db_time=")):
		e.DBTime = parseMs(f[8:])
	}
}

func normalizePath(path []byte) []byte {

	out := make([]byte, 0, len(path))

	segments := bytes.Split(path, []byte("/"))

	for i, seg := range segments {

		if i > 0 {
			out = append(out, '/')
		}

		if isNumeric(seg) {
			out = append(out, idToken...)
		} else {
			out = append(out, seg...)
		}
	}

	return out
}

func isNumeric(b []byte) bool {
	if len(b) == 0 {
		return false
	}

	for _, c := range b {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func atoiBytes(b []byte) int {
	n := 0
	for _, c := range b {
		if c < '0' || c > '9' {
			break
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func parseMs(b []byte) float64 {

	end := len(b)
	if end > 2 && b[end-2] == 'm' {
		end -= 2
	}

	var val float64
	var frac float64 = 0.1
	decimal := false

	for i := 0; i < end; i++ {
		c := b[i]

		if c == '.' {
			decimal = true
			continue
		}

		if c < '0' || c > '9' {
			break
		}

		if !decimal {
			val = val*10 + float64(c-'0')
		} else {
			val += float64(c-'0') * frac
			frac *= 0.1
		}
	}

	return val
}
