package unittest

import (
	"encoding/csv"
	"io"
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
