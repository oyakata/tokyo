package main

import (
	"fmt"
	fx "github.com/alangpierce/go-forceexport"
)

func main(){
	var nextStdChunk func(layout string) (prefix string, std int, suffix string)
	fx.GetFunc(&nextStdChunk, "time.nextStdChunk")

	fmt.Println(nextStdChunk("2006"))  // 273 => stdLongYear

	// var std0x = [...]int{stdZeroMonth, stdZeroDay, stdZeroHour12, stdZeroMinute, stdZeroSecond, stdYear}
	fmt.Println(nextStdChunk("2005"))  // 263 005 => std0x[layout[i+1]-'1']
	fmt.Println(nextStdChunk("005"))  // 0 528
	fmt.Println(nextStdChunk("2016"))  // 263 016
	fmt.Println(nextStdChunk("016"))  // 260 6 => std0x[layout[i+1]-'1']
	fmt.Println(nextStdChunk("6"))  // 6 0 => meaningless value
}
