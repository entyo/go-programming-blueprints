package backup

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DirHash は、パスで指定したディレクトリ以下のファイルの各種データからハッシュ値を生成し、返す
func DirHash(path string) (string, error) {
	hash := md5.New()
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		io.WriteString(hash, path)
		// hash(io.Writerを実装したオブジェクト)にファイルデータを書き込む
		fmt.Fprintf(hash, "%v", info.IsDir()) // hash.Hash はio.Writerを実装している
		fmt.Fprintf(hash, "%v", info.ModTime())
		fmt.Fprintf(hash, "%v", info.Mode())
		fmt.Fprintf(hash, "%v", info.Name())
		fmt.Fprintf(hash, "%v", info.Size())
		return nil
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil // ハッシュ値を計算しhexで返す
}
