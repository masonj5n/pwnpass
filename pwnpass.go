package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/howeyc/gopass"
	"github.com/masonj88/passwordChecker"
)

func printResults(isHidden, isPwned bool, checkedPass, numTimes string) {
			if(isPwned && isHidden) {
				fmt.Printf("This password has been pwned %s time(s)\n", numTimes)
			} else if (!isPwned && isHidden){
				fmt.Printf("This password has not been pwned\n")
			} else if (isPwned && !isHidden) {
				fmt.Printf("The password: %s has been pwned %s times\n", checkedPass, numTimes)
			} else {
				fmt.Printf("The password: %s has NOT been pwned\n", checkedPass)
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
		isPwned, checkedPass, numTimes := passwordChecker.CheckForPwnage(string(passwd))
		printResults(true, isPwned, checkedPass, numTimes)
	} else {
		file, err := os.Open(*batchPntr)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
      isPwned, checkedPass, numTimes := passwordChecker.CheckForPwnage(scanner.Text())
			printResults(false, isPwned, checkedPass, numTimes)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	}
}
