package backup

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Archiver は、バックアップ内容と保存先のパスを受け取りアーカイブする関数をもつ
type Archiver interface {
	DestFmt() func(int64) string
	Archive(src string, dest string) error
}

type zipper struct{}

// var ZIP Archiver = (*zipper)(nil) // nilを*zipper型にキャスト
// ・ZIP というexportedなinterface{}型の変数に代入しておくことで、実装を外部に見せない。
//   こうすることで、外部への影響なしに内部の実装を変更できる
// ・*zipperがArchiveメソッドを実装していないと、以下のようにアサーションが出る。
//   cannot use (*zipper)(nil) (type *zipper) as type Archiver in assignment:
//       *zipper does not implement Archiver (missing Archive method)

// DestFmt メソッドはファイル名の拡張子をzipに変更する func(int64) string {} を返す
func (z *zipper) DestFmt() func(int64) string {
	return func(i int64) string {
		return fmt.Sprintf("%d.zip", i)
	}
}

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		io.Copy(f, in)
		return nil
	})
}
