package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check_errors(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func check_user_files() (matches []string, err error) {
	pattern := "C:\\Users\\*\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadline\\ConsoleHost_history.txt"
	all_occurrences, err := filepath.Glob(pattern)
	check_errors(err)

	return all_occurrences, err

}

func write_csvfile(csv_file string, file_path string) {
	f, err := os.OpenFile(file_path, os.O_RDONLY, os.ModePerm)
	check_errors(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	username := strings.Split(file_path, "\\")[2]
	for sc.Scan() {
		line := sc.Text() // GET the line string
		fmt.Println(username, "Executed: ", line)

	}
	check_errors(err)

}

func main() {
	// get all the powershell paths
	csvfile := "powershell_history.csv"
	paths, err := check_user_files()
	// check for errors, if so, exit gracefully
	check_errors(err)
	// loop through all the files and write them to a csv file with user appended
	for _, path := range paths {
		write_csvfile(csvfile, path)
	}

}
