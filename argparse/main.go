package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.NewFlagSet("help", flag.ExitOnError)
	verbose := help.Bool("v", false, "詳しく")
	moreVerbose := help.Bool("vv", false, "より詳しく")
	mostVerbose := help.Bool("vvv", false, "最も詳しく") // FIXME: 4つ以上vが続いたら3つで処理したい

	money := flag.NewFlagSet("money", flag.ExitOnError)

	if len(os.Args) == 1 {
		fmt.Println("usage: argparse <command> [<args>]")
		fmt.Println("\thelp: ヘルプをプリントします")
		fmt.Println("\tmoney: 為替レートを調べます")
		return
	}

	switch os.Args[1] {
	case "help":
		help.Parse(os.Args[2:])
	case "money":
		money.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}

	if help.Parsed() {
		message := "Help me."
		if *verbose {
			message = "Help me!"
		}
		if *moreVerbose {
			message = "Help me!!"
		}
		if *mostVerbose {
			message = "HELP ME!!"
		}
		fmt.Println(message)
	}

	if money.Parsed() {
		if len(money.Args()) < 2 {
			// e.g. $ <cmd> money USD
			fmt.Println("エラー: 通貨をふたつ入力してください")
			return
		}

		from, to := money.Arg(0), money.Arg(1)
		if from == "" || to == "" {
			// e.g. $ <cmd> money USD ''
			fmt.Println("エラー: 通貨をふたつ入力してください")
			return
		}
		fmt.Printf("%v/%v: 100円/$\n", from, to) // 適当
	}
}
