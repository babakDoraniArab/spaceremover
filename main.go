package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	dir, err := os.Getwd()
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	files, err := ioutil.ReadDir(dir)
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := file.Name()
		newPath := strings.ReplaceAll(oldPath, " ", "_")
		
		if oldPath != newPath {
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Printf("Error renaming file %s to %s: %v\n", oldPath, newPath, err)
			} else {
				color.Green("Renamed file %s to %s\n", oldPath, newPath)
			}
		}
	}
}
