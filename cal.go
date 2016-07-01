// calc is intended to be a "same flavor" of lexing/parsing calculator
// as in
// https://github.com/xharaken/step2015/blob/master/calculator_modularize_2.py
// Please note that this is not a very canonical "Go" style. But this
// intentionally follows
// https://github.com/xharaken/step2015/blob/master/calculator_modularize_2.py
// nearly line-by-line and avoids concepts like interfaces, functors,
// or slices of strings.
//
// If you're curious, there's a very good talk by Rob Pike (one of the
// authors of Go) about a much cleaner way to write lexers like this,
// explaining how https://golang.org/pkg/text/template/ works
// internally at https://www.youtube.com/watch?v=HxaD_trXwRE (with
// slides at https://talks.golang.org/2011/lex.slide ).
package main

import (
	"bufio"
	"fmt"
	"os"
)

// kind identifies the kinds of different tokens.
// This would actually be simpler in Go if different kinds of tokens
// were just different types that shared a common interface.
type kind int

const (
	kindNumber = iota
	kindPlus
	kindMinus
	kindTimes
	kindDIvides
)

type token struct {
	kind   kind
	number float64
}

func readNumber(line string, index int) (token, int) {
	tok := token{kind: kindNumber}
	for ; index < len(line); index++ {
		c := line[index]
		if '0' <= c && c <= '9' {
			tok.number = tok.number*10 + float64(c-'0')
		} else {
			break
		}
	}
	if index < len(line) && line[index] == '.' {
		keta := 0.1
		for index++; index < len(line); index++ {
			c := line[index]
			if '0' <= c && c <= '9' {
				tok.number += float64(c-'0') * keta
			} else {
				break
			}
			keta *= 0.1
		}
	}
	return tok, index
}

func readPlus(line string, index int) (token, int) {
	return token{kind: kindPlus}, index + 1
}

func readMinus(line string, index int) (token, int) {
	return token{kind: kindMinus}, index + 1
}

func readTimes(line string, index int)(token, int){
	return token{kind: kindTimes},index + 1
}

func readDivides(line string, index int)(token, int){
		return token{kind: kindDivides}, index + 1

func tokenize(line string) []token {
	var tokens []token
	for index := 0; index < len(line); {
		var tok token
		switch {
		case '0' <= line[index] && line[index] <= '9':
			tok, index = readNumber(line, index)
		case '+' == line[index]:
			tok, index = readPlus(line, index)
		case '-' == line[index]:
			tok, index = readMinus(line, index)
		case '*' == line[index]:
			tok, index = readTimes(line,index)
		case '/' == line[index]:
			tok, index = readDivides(line,index)
		default:
			panic(fmt.Errorf("invalid character from %q@%v: %v", line, index, line[index]))
		}
		tokens = append(tokens, tok)
	}
	return tokens
}

func evaluate(tokens []token) float64 {
	answer := float64(0)
	tokens = append([]token{token{kind: kindPlus}}, tokens...) // insert a dummy '+' token
	for idx := 1; idx < len(tokens); idx++ {
		switch tokens[idx].kind {
		case kindNumber:
			switch tokens[idx-1].kind {
			case kindTimes:
				answer = answer * tokens[idx].number
			case kindDivides:
				answer = answer / tokens[idx].number
			case kindPlus:
				answer += tokens[idx].number
			case kindMinus:
				answer -= tokens[idx].number
			}
		}
	}
	return answer
}

func evalLine(line string) float64 {
	tokens := tokenize(line)
	answer := evaluate(tokens)
	return answer
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		answer := evalLine(line)
		fmt.Printf("answer = %f\n", answer)
	}
}
