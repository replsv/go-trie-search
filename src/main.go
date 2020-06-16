package main

import (
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"
)

var cpuUsed = 1
var m sync.Mutex
var dir string
var search string
var tries []*Trie

func init() {
	dirPtr := flag.String("dir", "./docs/", "Directory for scanning documents")
	searchPtr := flag.String("search", "romania", "String to search for")
	flag.Parse()
	dir = string(*dirPtr)
	search = strings.ToLower(string(*searchPtr))
}

func main() {
	loadAsTries()

	fmt.Printf("[!] Ready. Loaded %d tries\n", len(tries))

	performSearch()
}

func loadAsTries() {
	start := time.Now()
	paths := getFilesFromDir(dir)
	wg := &sync.WaitGroup{}
	wg.Add(len(paths))
	fmt.Println(paths)
	for _, path := range paths {
		go func(path string) {
			defer wg.Done()
			content := readFile(path)
			trie, err := buildTrie(content)
			if err == nil {
				trie.SetData("file", path)
				tries = append(tries, trie)
			}
		}(path)
	}
	wg.Wait()
	fmt.Printf("[~] Elapsed: %s\n", time.Since(start))
}

func performSearch() {
	start := time.Now()
	if len(tries) > 0 {
		fmt.Printf("[!] Performing search for term: %s\n", search)
		results := []string{}
		wg := sync.WaitGroup{}
		for _, trie := range tries {
			wg.Add(1)
			go func(trie *Trie) {
				defer wg.Done()
				if trie.Find(search) != nil {
					results = append(results, trie.GetData("file").(string))
				}
			}(trie)
		}
		wg.Wait()
		if len(results) > 0 {
			fmt.Printf("[!] Found results in %d file(s)\n", len(results))
			for _, resultFile := range results {
				fmt.Printf("~ %s\n", resultFile)
			}
		} else {
			fmt.Println("No results :(")
		}
	} else {
		fmt.Println("Nothing to do")
	}
	fmt.Printf("[~] Elapsed: %s\n", time.Since(start))
}
