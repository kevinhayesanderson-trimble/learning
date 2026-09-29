package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func getConfigValue(key string) (string, bool) {
	return "localhost", true
}

func main1() {
	const defaultServer = "localhost"
	// getConfigValue returns false if the named value does not exist
	serverAddress, found := getConfigValue("server")
	if !found {
		serverAddress := defaultServer
		fmt.Println("default server set:", serverAddress)
	}
	fmt.Println("connecting to server:", serverAddress)
}

func wordCountOld() int {
	text := "let's count some words"

	var numSpaces int

	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			numSpaces++
		} else {
			text := text[i]
			fmt.Println("Found", text, "words")
		}
	}

	return numSpaces
}

func wordCount() int {
	text := "let's count some words"
	return len(strings.Fields(text))
}

func main() {
	if len(os.Args)<2{
		log.Println("need to provide filename!")
		os.Exit(1)
	}
	fmt.Println(os.Args[0])
	for _, fileName := range os.Args[1:] {
		file, err := os.Open(fileName)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		// words := strings.Fields(string(string(fileContents)))
		// lenWords := len(words)
		// wordsB := bytes.Fields(fileContents)
		// for i, word := range words{
		// 	fmt.Printf("Word %d: %s\n", i, string(word))
		// }
		// lenWordsB := len(wordsB)
		// //fmt.Println("Found", wordCount(), "words")
		// fmt.Printf("Found %d words", lenWords)
		// fmt.Printf("Found %d words", lenWordsB)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		var wordcount int
		for scanner.Scan() {
			wordcount++
		}
		if scanner.Err() != nil{
			log.Println(scanner.Err())
			os.Exit(1)
		}
		fmt.Printf("Found %d words", wordcount)
	}
	os.Exit(0)
}