package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/btcsuite/btcutil"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first part of the key: ")
	firstPart, _ := reader.ReadString('\n')
	firstPart = strings.TrimSpace(firstPart)

	fmt.Print("Enter the second part of the key: ")
	secondPart, _ := reader.ReadString('\n')
	secondPart = strings.TrimSpace(secondPart)

	var wg sync.WaitGroup
	keyChan := make(chan string)
	correctKeys := make([]string, 0)

	wg.Add(1)
	go func() {
		defer wg.Done()
		generateCombinations(secondPart, "", keyChan)
		close(keyChan)
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for key := range keyChan {
				if verifyChecksum(firstPart + key) {
					log.Printf("Correct key found: %s\n", firstPart+key)
					correctKeys = append(correctKeys, firstPart+key)
				}
			}
		}()
	}

	wg.Wait()

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, line := range correctKeys {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Possible correct keys written to file.")
}

func generateCombinations(chars, prefix string, keyChan chan<- string) {
	if len(chars) == 0 {
		keyChan <- prefix
		return
	}

	for i := 0; i < len(chars); i++ {
		newPrefix := prefix + string(chars[i])
		generateCombinations(chars[:i]+chars[i+1:], newPrefix, keyChan)
	}
}

func verifyChecksum(key string) bool {
	decodedKey, err := btcutil.DecodeWIF(key)
	if err != nil {
		return false
	}

	return decodedKey != nil
}
