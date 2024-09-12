package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(".env")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close() 

	newFile, err := os.Create("template.env")
	if err != nil {
		log.Panic(err)
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		fmt.Fprint(newFile, parts[0],"=\n")
	}
}
