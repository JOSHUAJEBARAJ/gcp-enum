package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	GCP_URL = "https://www.googleapis.com/storage/v1/b/"
	green   = "\033[32m" // Green color
	red     = "\033[31m" // Red color
	reset   = "\033[0m"  // Reset color
)

func getwords(filename string) []string {

	file, err := os.Open(filename)
	_ = file
	if err != nil {
		exit(fmt.Sprintf("Error Opening %s", filename))
	}

	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		wordlist = append(wordlist, word)
	}

	file.Close()
	return wordlist

}
func main() {

	jobs := make(chan string, 100)

	key := flag.String("k", "", "keyword that you want to enumerate")
	filename := flag.String("file", "", "File name containing the word list")
	concurrency := flag.Int("c", 5, "Default concurrency value is 5 you can change the value using the c flag")
	flag.Parse()

	if *key == "" && *filename == "" {
		flag.PrintDefaults()
		exit("")

	}

	// function to get the words from the file
	//fmt.Println(time.Since(start))

	wordlist := getwords(*filename)
	//fmt.Println(time.Since(start))

	// function to generate the worldlist using the file key and words from the file
	endpoints := mutate(*key, wordlist)
	// fmt.Println(time.Since(start))

	//fmt.Println(len(endpoints))
	//  go routine

	var wg sync.WaitGroup

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go enumerate(jobs, &wg)
	}

	for _, value := range endpoints {
		jobs <- value

	}
	close(jobs)
	wg.Wait()

}

func enumerate(jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// reading the permutaed world
		url := fmt.Sprintf("https://www.googleapis.com/storage/v1/b/%s", job)
		res, err := http.Head(url)
		if err != nil {
			panic(err)
		}

		if res.StatusCode != (400) && res.StatusCode != (404) {
			fmt.Println(green + "Valid:" + job + reset)

		} else {
			fmt.Println(red + "Invalid:" + job + reset)
		}

	}

}

func mutate(key string, words []string) []string {
	var ret []string
	for _, word := range words {
		ret = append(ret, key+"-"+word, word+"-"+key, key+"_"+word, word+"_"+key, key+word, word+key)
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
