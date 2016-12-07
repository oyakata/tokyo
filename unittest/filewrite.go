package unittest

import (
	"encoding/csv"
	"os"
)

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
