package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

// 独自の出力用フォーマット
func (p path) String() string {
	return fmt.Sprintf("%s [%s]", p.Path, p.Hash)
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()
	var dbpath = flag.String("db", "./backupdata", "dbへのパス")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("エラー: コマンドを指定してください")
		return
	}
	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()
	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	switch strings.ToLower(args[0]) {
	case "list":
		var path path
		col.ForEach(func(i int, data []byte) bool {
			err := json.Unmarshal(data, &path) // パス(フィールド)に対応した値をセットする
			if err != nil {
				fatalErr = err
				return true
			}
			fmt.Printf("= %s\n", path.String())
			return false
		})
	case "add":
		// add 以降に何も指定されなかった場合
		if len(args[1:]) == 0 {
			fatalErr = errors.New("追加するパスを指定してください")
			return
		}
		for _, p := range args[1:] {
			path := &path{Path: p, Hash: "まだアーカイブされていません"}
			if err := col.InsertJSON(path); err != nil {
				fatalErr = err
				return
			}
			fmt.Printf("+ %s\n", path.String())
		}
	case "remove":
		var path path
		col.RemoveEach(func(i int, data []byte) (bool, bool) {
			err := json.Unmarshal(data, &path)
			if err != nil {
				fatalErr = err
				return false, true
			}
			for _, p := range args[1:] {
				if path.Path == p {
					fmt.Printf("- %s\n", path.String())
					return true, false
				}
			}
			return false, false
		})
	}
}
