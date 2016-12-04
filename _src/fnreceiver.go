package main


import "fmt"


type Score int
func (s Score) Show() { fmt.Printf("Your score is %d\n", s) }


func main(){
	var myScore Score = 255
	myScore.Show()

	// 型のキャスト
	x := 12
	(Score(x)).Show()
}
