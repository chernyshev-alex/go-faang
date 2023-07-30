package parscomb

import (
	"fmt"
	"io"
	"strings"
)

type MTerminal struct {
	char rune
	len  int
}
type MParser func(io.RuneScanner) (interface{}, io.RuneScanner)

func (t MTerminal) str() string {
	if t.len == 1 {
		return fmt.Sprintf("%c", t.char)
	} else {
		return fmt.Sprintf("%d%c", t.len, t.char)
	}
}

func parseAndCompress(input string) string {
	seqParser := SeqStringParser()
	parser := ApplyParser(seqParser)
	nodes, _ := parser(strings.NewReader(input))
	s := ""
	for _, t := range nodes.([]MTerminal) {
		s += t.str()
	}
	return s
}

func SeqStringParser() MParser {
	return func(s io.RuneScanner) (interface{}, io.RuneScanner) {
		first, _, e := s.ReadRune()
		cnt, next := 1, first
		for e == nil {
			if next, _, e = s.ReadRune(); e == nil {
				if first != next {
					s.UnreadRune()
					return MTerminal{first, cnt}, s
				} else {
					cnt += 1
				}
			} else {
				return MTerminal{first, cnt}, nil
			}
		}
		return MTerminal{first, cnt}, nil
	}
}

func ApplyParser(mp MParser) MParser {
	return func(s io.RuneScanner) (interface{}, io.RuneScanner) {
		result := make([]MTerminal, 0)
		var node interface{}
		for scanner := s; scanner != nil; {
			node, scanner = mp(scanner)
			if node != nil {
				result = append(result, node.(MTerminal))
			}
		}
		return result, nil
	}
}
