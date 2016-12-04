package main

import (
	"fmt"
)

type Employee struct {
	name string
	age int
	salary int
}

type LaborInsurance interface {
	Show()
	Join(number int) string
}

type Engineer Employee
type Longshoreman Employee

func (e *Engineer) Show() {
	fmt.Println(e.name)
}

func (e *Engineer) Join(number int) string {
	return fmt.Sprintf("'%v' joined to %v person(s)!", e.name, number)
}

// FIXME: これは、どうしてダメなんだろう・・？
// func Cert(labor *LaborInsurance, number int) {
// 	labor.Show()
// 	labor.Join(number)
// }

func main(){
	ken := Engineer{"Ken Pepple", 18, 240}
	bob := Longshoreman{"Bob Marine", 20, 100}
	fmt.Println(ken)
	fmt.Println(bob)
	ken.Show()
	fmt.Println(ken.Join(24))
	// Cert(ken, 1208)

	who := LaborInsurance(&ken)
	who.Show()

	// bob.Show undefined (type Longshoreman has no field or method Show)
	// bob.Show()
}
