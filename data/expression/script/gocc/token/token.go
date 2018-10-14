// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"||",
		"&&",
		"==",
		"!=",
		"<",
		"<=",
		">",
		">=",
		"+",
		"-",
		"*",
		"/",
		"%",
		"!",
		"ident",
		"(",
		")",
		"functionName",
		"?",
		":",
		"true",
		"false",
		"nil",
		"null",
		"intLit",
		"floatLit",
		"stringLit",
		"ref",
		"selector",
		"index",
		"empty",
		",",
	},

	idMap: map[string]Type{
		"INVALID":      0,
		"$":            1,
		"||":           2,
		"&&":           3,
		"==":           4,
		"!=":           5,
		"<":            6,
		"<=":           7,
		">":            8,
		">=":           9,
		"+":            10,
		"-":            11,
		"*":            12,
		"/":            13,
		"%":            14,
		"!":            15,
		"ident":        16,
		"(":            17,
		")":            18,
		"functionName": 19,
		"?":            20,
		":":            21,
		"true":         22,
		"false":        23,
		"nil":          24,
		"null":         25,
		"intLit":       26,
		"floatLit":     27,
		"stringLit":    28,
		"ref":          29,
		"selector":     30,
		"index":        31,
		"empty":        32,
		",":            33,
	},
}
