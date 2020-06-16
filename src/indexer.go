package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func getFilesFromDir(dir string) []string {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".txt" {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return paths
}

func buildTrie(content string) (*Trie, error) {
	if content == "" {
		return nil, errors.New("Missing content")
	}
	trie := NewTrie()
	for _, word := range strings.Split(content, " ") {
		trie.Add(strings.ToLower(word))
	}
	return trie, nil
}
