package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check_errors(err error) {
	if err != nil {
		log.Fatal(err)
		//os.Exit(0)
	}
}

func check_user_files() (matches []string, err error) {
	pattern := "C:\\Users\\*\\AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadline\\ConsoleHost_history.txt"
	all_occurrences, err := filepath.Glob(pattern)
	check_errors(err)

	return all_occurrences, err

}

func write_csvfile(csv_writer *csv.Writer, file_path string) {
	f, err := os.OpenFile(file_path, os.O_WRONLY, os.ModePerm)
	check_errors(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	username := strings.Split(file_path, "\\")[2]
	for sc.Scan() {
		line := sc.Text() // GET the command and append username string
		setup := []string{username, line}
		fmt.Println(setup)
		csv_writer.Write(setup)
	}
	check_errors(err)

}

func main() {
	// get all the powershell paths
	csvfilename := "powershell_history.csv"
	csvfile, err := os.Create(csvfilename)
	check_errors(err)
	csvwriter := csv.NewWriter(csvfile)
	header := []string{"User", "Command"}
	csvwriter.Write(header)
	paths, err := check_user_files()
	// check for errors, if so, exit gracefully
	check_errors(err)
	// loop through all the files and write them to a csv file with user appended
	for _, path := range paths {
		write_csvfile(csvwriter, path)
	}
	csvwriter.Flush()
	csvfile.Close()

}
