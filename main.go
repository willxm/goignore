package main

import (
	"fmt"
	"log"
	"os"

	"github.com/willxm/goignore/cmd/goignore"
)

const (
	Override = "override"
	Update   = "update"
)

func main() {
	input := os.Args
	if len(input) <= 1 {
		fmt.Printf("please input gitignore type\n")
		return
	}
	typeArg := input[1]
	OperationType := ""

	if goignore.IsExistGitignore() {
		fmt.Printf(".gitignore is existed, do you want override or update it ?\n")
		fmt.Scanf("%s", &OperationType)
		if OperationType != Override && OperationType != Update {
			fmt.Printf("please input the true Operation!\n")
			return
		}
	} else {
		err := goignore.CreateGitignore()
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	var gitignoreContent string

	if OperationType == Update {
		data, err := goignore.ReadTheExistGitignore()
		if err != nil {
			log.Fatal(err)
			return
		}
		gitignoreContent = string(data)
	} else {
		err := goignore.TruncateGitignore()
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	f, err := goignore.OpenGitignore()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	if typeArg == "go" {
		gitignoreContent = goignore.GoGitignore
	} else if typeArg == "c" {
		gitignoreContent = goignore.CGitignore
	}
	_, err = f.WriteString(gitignoreContent)
	if err != nil {
		log.Fatal(err)
		return
	}
}
