package main

import "fmt"

type Employee interface {
	Pay() int
	Work(days int) int
}

type Employ func(emp []Employee) int

func Manage(fn Employ) Employ {
	return func(emp []Employee) int {
		fmt.Println("=== START ===")
		result := fn(emp)
		fmt.Printf("=== END[result:%v] ===\n", result)
		return result
	}
}

type WageSlave struct {
	id int
	name string
	wage int
}

// func Foo(emp... WageSlave) int{
// 	acc := 0
// 	for _, x := range emp{
// 		acc += x.wage
// 	}
// 	return acc
// }

func (self *WageSlave) Pay() int {
	return self.wage
}

func (self *WageSlave) Work(days int) int {
	return self.wage * days
}

func main(){
	fn := Manage(func (emp []Employee) int{
		acc := 0
		// lenはスライスじゃないと呼べません。配列には適用できません。
		for i := 0; i < len(emp); i++ {
			acc += emp[i].Pay()
		}
		return acc
	})

	// x, y, z := WageSlave{1, "foo", 30}, WageSlave{2, "bar", 16}, WageSlave{3, "baz", 62}
	// result := fn(&x, &y, &z)

	// xs = []WageSlave{
	// 	{1, "foo", 30},
	// 	{2, "bar", 16},
	// 	{3, "baz", 62},
	// }

	xs := make([]Employee, 3)
	// Pay, WorkのレシーバーをEmployeeのポインタで宣言しているので、WageSlaveの頭に&を付けないとコンパイルが通りません。
	xs[0] = Employee(&WageSlave{1, "foo", 30})
	xs[1] = Employee(&WageSlave{2, "bar", 16})
	xs[2] = Employee(&WageSlave{3, "baz", 62})

	// 可変長引数で宣言した場合も、このようにひとつずつ頭に&を付けないとコンパイルエラーです。
	// result := fn(&xs[0], &xs[1])
	// result := fn(&xs[0])

	// []WageSlaveで宣言したスライスを一発で[]Employeeにキャストはできません。
	// 以下のコードもコンパイルが通りません。
	// xs := make([]WageSlave, 3)
	// xs[0] = WageSlave{1, "foo", 30}
	// xs[1] = WageSlave{2, "bar", 16}
	// xs[2] = WageSlave{3, "baz", 62}
	// cannot convert xs (type []WageSlave) to type []Employee
	// result := fn([]Employee(xs))

	result := fn(xs)

	// 要するに今のわたしは可変長引数とスライスと配列とポインタがよくわかっていないことがよくわかりました。
	fmt.Println(result)
}
