package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Boeing/config-file-validator/internal/entity"
	"github.com/Boeing/config-file-validator/internal/usecase/finder"
	"github.com/Boeing/config-file-validator/internal/usecase/reporter"
)

/*
Validator recusively scans a directory to search for configuration files and
validates them using the go package for each configuration type.

Currently json, yaml, toml, and xml configuration file types are supported.

Usage:

    validator [flags]

The flags are:
    -search-path string
		The search path for configuration files
    -exclude-dirs string
    	Subdirectories to exclude when searching for configuration files
*/

func main() {
	search_path := flag.String("search-path", "", "The search path for configuration files")
	exclude_dirs := flag.String("exclude-dirs", "", "Subdirectories to exclude when searching for configuration files")
	report_format := flag.String("report-type", "standard", "The format the report should be printed to stdout")
	flag.Parse()
	exclude := strings.Split(*exclude_dirs, ",")

	finder := finder.NewService(search_path, exclude)
	reporter, err := reporter.NewService(*report_format)
	if err != nil {
		fmt.Printf("error creating reporter service: %s", err)
		os.Exit(1)
	}

	exitCode := execute(finder, reporter)

	os.Exit(exitCode)
}

func execute(finder finder.InteractorI, reporter reporter.InteractorI) int {
	found_files, err := finder.Find()
	if err != nil {
		fmt.Printf("error finding files: %s", err)
		return 1
	}

	var reports []entity.Report

	for _, file := range found_files {
		fileContent, err := ioutil.ReadFile(file.Path)
		if err != nil {
			fmt.Printf("error reading file: %s", err)
			return 1
		}

		isValid, err := file.FileType.Validator.Validate(fileContent)

		report := entity.Report{
			FileName:        file.Name,
			FilePath:        file.Path,
			IsValid:         isValid,
			ValidationError: err,
		}

		reports = append(reports, report)
	}

	err = reporter.Print(reports)
	if err != nil {
		fmt.Printf("error printing report: %s", err)
		return 1
	}

	return 0
}
