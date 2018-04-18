package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/masonj88/pwchecker"
	"log"
	"os"
)

func printResults(isHidden, isPwned bool, checkedPass, numTimes string) {
	if isPwned {
		if isHidden {
			fmt.Printf("This password has been pwned %s time(s)\n", numTimes)
		} else {
			fmt.Printf("The password: %s has been pwned %s time(s)\n", checkedPass, numTimes)
		}
	} else {
		if isHidden {
			fmt.Printf("This password has not been pwned\n")
		} else {
			fmt.Printf("The password: %s has NOT been pwned\n", checkedPass)
		}
	}
}

func main() {

	// Parse flags for batch processing
	batchPntr := flag.String("batch", "none", "Input path of file to be batch processed")
	flag.Parse()
	// Get password input if not batch processing
	if *batchPntr == "none" {
		fmt.Printf("Password: ")
		passwd, err := gopass.GetPasswd()
		if err != nil {
			fmt.Println("Error parsing password")
		}
		rpwd, err := pwchecker.CheckForPwnage(string(passwd))
		if err != nil {
			fmt.Println("Couldn't return processed password")
			panic(err)
		}
		printResults(true, rpwd.Pwnd, rpwd.Pwd, rpwd.TmPwnd)

	} else {
		file, err := os.Open(*batchPntr)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			rpwd, err := pwchecker.CheckForPwnage(scanner.Text())
			if err != nil {
				fmt.Println("Couldn't return processed password")
				panic(err)
			}
			printResults(false, rpwd.Pwnd, rpwd.Pwd, rpwd.TmPwnd)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
