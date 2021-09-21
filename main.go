package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./queries.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	i := 0
	for scanner.Scan() {
		fmt.Printf("[%d]\n", i)
		fmt.Printf("query=%s\n", scanner.Text())
		fmt.Printf("query-results-file=/tmp/%d.csv\n", i)
		fmt.Println("count=1")
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
