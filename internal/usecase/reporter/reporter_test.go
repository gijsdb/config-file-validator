package reporter

import (
	"errors"
	"testing"

	"github.com/Boeing/config-file-validator/internal/entity"
)

func Test_stdoutReport(t *testing.T) {
	reportNoValidationError := entity.Report{
		"good.xml",
		"/fake/path/good.xml",
		true,
		nil,
	}

	reportWithValidationError := entity.Report{
		"bad.xml",
		"/fake/path/bad.xml",
		false,
		errors.New("Unable to parse bad.xml file"),
	}

	reports := []entity.Report{reportNoValidationError, reportWithValidationError}

	stdoutReporter := StdoutReporter{}
	err := stdoutReporter.Print(reports)
	if err != nil {
		t.Errorf("Reporting failed")
	}
}
