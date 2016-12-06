package main

import "fmt"

type Challenger struct {
	count     int
	threshold int
}

type Retryable struct{}

// 非公開メソッドの方はわざとpanicを利用する。
func (c *Challenger) attack() bool {
	if c.count < c.threshold {
		c.count++
		panic(Retryable{}) // パニック、しかしリトライ可能
	}
	if c.count > c.threshold {
		panic("No more chance.")
	}
	c.count++
	return true
}

// 公開メソッドはpanicを捕捉してリカバーをかける。
func (c *Challenger) Attack() (ok bool, err error) {
	defer func() {
		switch p := recover(); p {
		case nil:
			// ここはパニックが起きなかった場合が該当する
			fmt.Println("No error.")
		case Retryable{}:
			err = fmt.Errorf("failed(%v times)", c.count)
		default:
			panic(p)
		}
	}()
	ok = c.attack()
	return
}

func main() {
	var c *Challenger
	c = &Challenger{0, 3}
	fmt.Println(c.Attack()) // false failed(1 times)
	fmt.Println(c.Attack()) // false failed(2 times)
	fmt.Println(c.Attack()) // false failed(3 times)
	fmt.Println(c.Attack()) // true <nil>
	fmt.Println(c.Attack()) // 以下panic

	// panic: No more chance. [recovered]
	// 	panic: No more chance.
	//
	// goroutine 1 [running]:
	// panic(0x4b9f40, 0xc82000a400)
	// 	/usr/lib/go-1.6/src/runtime/panic.go:481 +0x3e6
	// main.(*Challenger).Attack.func1(0xc820039e30, 0xc820039e58)
	// 	/home/echizen/_go/src/github.com/oyakata/tokyo/trash/retry.go:35 +0x2fc
	// panic(0x4b9f40, 0xc82000a400)
	// 	/usr/lib/go-1.6/src/runtime/panic.go:443 +0x4e9
	// main.(*Challenger).attack(0xc820039e58, 0x52c678)
	// 	/home/echizen/_go/src/github.com/oyakata/tokyo/trash/retry.go:19 +0xd0
	// main.(*Challenger).Attack(0xc820039e58, 0x0, 0x0, 0x0)
	// 	/home/echizen/_go/src/github.com/oyakata/tokyo/trash/retry.go:38 +0x78
	// main.main()
	// 	/home/echizen/_go/src/github.com/oyakata/tokyo/trash/retry.go:49 +0x52c
	// exit status 2
}
