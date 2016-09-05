package main

import "strings"

// PathSeparator は"/"
const PathSeparator = "/"

// Path は１つのURLを表す構造体
type Path struct {
	Path string
	ID   string
}

// NewPath は文字列からPathを生成するための関数
func NewPath(p string) *Path {
	var id string
	p = strings.Trim(p, PathSeparator)
	s := strings.Split(p, PathSeparator)
	if len(s) > 1 {
		endIndex := len(s) - 1
		id = s[endIndex]
		p = strings.Join(s[:endIndex], PathSeparator)
	}
	return &Path{Path: p, ID: id}
}

// HasID はPathがIDを持つかを調べる関数
func (p *Path) HasID() bool {
	return len(p.ID) > 0
}
