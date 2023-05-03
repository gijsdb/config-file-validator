package main

import (
	"testing"

	"github.com/Boeing/config-file-validator/internal/usecase/finder"
	"github.com/Boeing/config-file-validator/internal/usecase/reporter"
)

func Test_ExecuteNoFilesFound(t *testing.T) {
	search_path := "/wrong/path/to/files"
	report_format := "standard"

	finder := finder.NewService(search_path, nil)
	reporter, _ := reporter.NewService(report_format)

	exitCode := execute(finder, reporter)

	if exitCode == 0 {
		t.Error("incorrect result: function returned OK exitcode")
	}
}
