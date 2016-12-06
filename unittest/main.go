package unittest

import (
	"encoding/csv"
	"io"
	"os"
)

type Entry map[string]string

// ファイルを処理するプログラム
func CSVToMaps(in io.Reader) (xs []Entry, err error) {
	var r *csv.Reader
	r = csv.NewReader(in)

	var names []string
	var entry Entry
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if names == nil {
			// 最初の行
			names = record
		} else {
			entry = Entry{}
			for i := 0; i < len(names); i++ {
				entry[names[i]] = record[i]
			}
			xs = append(xs, entry)
		}
	}
	return
}

// ローカルのファイルを出力するプログラム (S3とかに出力する場合はこの方法ではダメ)
func WriteCSV(names []string, xs []Entry, filepath string) {
	out, _ := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer out.Close()
	w := csv.NewWriter(out)

	defer w.Flush()
	w.Write(names)

	values := make([]string, len(names))
	for _, x := range xs {
		for i := 0; i < len(names); i++ {
			values[i] = x[names[i]]
		}
		w.Write(values)
	}
}

// ログを出力するプログラム
// モックやスタブ
// 設定ファイルの設定値に依存するテスト
// 例外を出すことをテスト
