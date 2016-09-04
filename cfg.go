package cfg

import (
	"fmt"
	"io/ioutil"
)

type section map[string]string
type config map[string]section

func ParseFile(path string) (cfg config, err error) {

	src, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	cfg = make(config)

	b := newScanner(string(src))
	for {
		skipSpaces(b)
		if !b.more() {
			break
		}

		var name string
		var sec section
		name, sec, err = readSection(b)
		if err != nil {
			break
		}
		cfg[name] = sec
	}
	return
}

func readSection(b *scanner) (name string, sec section, err error) {

	name = readName(b)
	if name == "" {
		err = fmt.Errorf("Identifier expected")
		return
	}

	skipSpaces(b)
	if !b.more() {
		return
	}
	b.expect('{')
	skipSpaces(b)

	sec = make(section)

	for b.next() != '}' {
		key := readName(b)
		if key == "" {
			err = fmt.Errorf("Property expected")
			return
		}

		for b.next() == ' ' || b.next() == '\t' {
			b.get()
		}

		val := ""
		for b.next() != '\n' && b.next() != '\r' {
			val += string(b.get())
		}

		sec[key] = val

		skipSpaces(b)
	}
	b.expect('}')
	err = b.err
	return
}

func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func skipSpaces(b *scanner) {
	for isSpace(b.next()) {
		b.get()
	}
}

func readName(b *scanner) string {
	name := ""
	if !isAlpha(b.next()) {
		return name
	}

	for isAlpha(b.next()) || isDigit(b.next()) {
		name += string(b.get())
	}
	return name
}
