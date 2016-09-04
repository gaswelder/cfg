package cfg

import (
	"fmt"
)

type scanner struct {
	str string
	pos int
	len int
	err error
}

func newScanner(s string) *scanner {
	return &scanner{s, 0, len(s), nil}
}

func (s *scanner) more() bool {
	if s.err != nil {
		return false
	}
	return s.pos < s.len
}

func (s *scanner) next() byte {
	if !s.more() {
		return 0
	}
	return s.str[s.pos]
}

func (s *scanner) get() byte {
	if !s.more() {
		return 0
	}
	ch := s.str[s.pos]
	s.pos++
	return ch
}

func (s *scanner) rest() string {
	return s.str[s.pos:]
}

func (s *scanner) expect(ch byte) {
	if s.err != nil {
		return
	}
	n := s.get()
	if n != ch {
		s.err = fmt.Errorf("expected %c, got %c", ch, n)
	}
}
