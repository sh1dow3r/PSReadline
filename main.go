package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check_errors(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err)
		os.Exit(0)
	}
}

func check_user_files(argument string) (matches []string, err error) {
	pattern := ""
	if argument == "All" || argument == "all" || argument == "ALL" {
		pattern = "C:\\Users\\*\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\ConsoleHost_history.txt"
		fmt.Println(pattern)
	} else {
		pattern = "C:\\Users\\" + argument + "\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine\\ConsoleHost_history.txt"
		fmt.Println(pattern)
	}

	all_occurrences, err := filepath.Glob(pattern)
	check_errors(err)
	return all_occurrences, err
}

func read_print_csv_entries(file_path string) {
	f, err := os.OpenFile(file_path, os.O_RDWR, os.ModePerm)
	check_errors(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	username := strings.Split(file_path, "\\")[2]
	for sc.Scan() {
		line := sc.Text() // GET the command and append username string
		setup := username + ", " + line
		fmt.Println(setup)
	}
}

func main() {
	argument_length := len(os.Args)
	if argument_length != 2 {
		fmt.Println("[E] Failed to run with the correct format...")
		fmt.Println("Format: main.exe {All | <username>}")
	} else {
		argument := os.Args[1]
		header := "User, Command"
		fmt.Println(header)
		fmt.Println(argument)
		paths, err := check_user_files(argument)
		// check for errors, if so, exit gracefully
		check_errors(err)

		// loop through all the files and write them to a csv file with user appended
		for _, path := range paths {
			read_print_csv_entries(path)
		}
	}
}
