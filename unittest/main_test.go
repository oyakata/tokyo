package unittest

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

// ファイルを入力に受け付けるプログラムのテスト: strings.NewReader/bytes.NewReaderで対処
func TestCSVToMaps(t *testing.T) {
	fp := strings.NewReader("" +
		"id,name,age\n" +
		"1,foo,18\n" +
		"2,bar,\n" +
		"3,,5")

	xs, err := CSVToMaps(fp)
	if err != nil {
		t.Errorf("ファイルの読み込みに失敗しました: %v", err)
	}

	expected := []Entry{
		{"id": "1", "name": "foo", "age": "18"},
		{"id": "2", "name": "bar", "age": ""},
		{"id": "3", "name": "", "age": "5"},
	}

	if !reflect.DeepEqual(xs, expected) {
		t.Errorf("結果が期待値と一致しません: %v != %v", xs, expected)
	}
}

// ファイルを出力するプログラムのテスト: ioutil.TempFileで対処
func TestWriteCSV(t *testing.T) {
	tmp, _ := ioutil.TempFile("", "_")
	defer func() {
		os.Remove(tmp.Name())
	}()

	data := []Entry{
		{"id": "1", "name": "foo", "age": "18"},
		{"id": "2", "name": "bar", "age": ""},
		{"id": "3", "name": "", "age": "5"},
	}
	names := []string{"id", "name", "age"}
	WriteCSV(names, data, tmp.Name())

	in, _ := os.Open(tmp.Name())
	defer in.Close()
	bs, err := ioutil.ReadAll(in)
	if err != nil {
		t.Errorf("処理に失敗しました: %v", err)
	}

	result := string(bs)	

	expected := ("" +
		"id,name,age\n" +
		"1,foo,18\n" +
		"2,bar,\n" +
		"3,,5\n")
	if result != expected {
		t.Errorf("結果が期待通りではありませんでした: %v != %v", result, expected)
	}
}
