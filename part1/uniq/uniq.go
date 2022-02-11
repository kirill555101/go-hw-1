package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	c bool
	d bool
	u bool
	f int
	s int
	i bool
}

func NewOptions(c, d, u bool, f, s int, i bool) *Options {
	return &Options{
		c: c,
		d: d,
		u: u,
		f: f,
		s: s,
		i: i,
	}
}

func (options *Options) CorrectParams() bool {
	return options.c && !options.d && !options.u || !options.c && options.d && !options.u ||
		!options.c && !options.d && options.u || !options.c && !options.d && !options.u
}

func (options *Options) ResetParams() {
	if !options.c && !options.u && !options.d {
		options.u = true
		options.d = true
	}
}

func Uniq(options *Options, text []string) []string {
	var result []string
	var repeats int
	previousLine := skip(options, text[0])

	if !options.c && options.u && options.d {
		show(options, &result, text[0], repeats)
	}

	for index := 1; index < len(text); index++ {
		currentLine := text[index]

		if options.i {
			previousLine = strings.ToLower(previousLine)
			currentLine = strings.ToLower(currentLine)
		}
		currentLine = skip(options, currentLine)

		if previousLine == currentLine {
			repeats++
			continue
		}

		if options.c || !options.d || !options.u {
			show(options, &result, text[index-1], repeats)
		}
		if !options.c && options.u && options.d {
			show(options, &result, text[index], repeats)
		}

		previousLine = currentLine
		repeats = 0
	}

	if options.c || !options.d || !options.u {
		show(options, &result, text[len(text)-1], repeats)
	}

	return result
}

func show(options *Options, result *[]string, str string, repeats int) {
	switch {
	case options.c:
		*result = append(*result, strconv.Itoa(repeats + 1) + " " + str)

	case (options.d && repeats != 0) || (options.u && repeats == 0):
		*result = append(*result, str)
	}
}

func skip(options *Options, str string) string {
	fields := strings.Split(str, " ")
	if len(fields) < options.f {
		return "\n"
	}

	str = strings.Join(fields[options.f:], " ")
	if len(str) < options.s {
		return "\n"
	}

	return str[options.s:]
}
