package finder

import (
	"testing"
)

func Test_fsFinder(t *testing.T) {
	fsFinder := NewService("../../../test/fixtures", nil)
	files, err := fsFinder.Find()

	if len(files) < 1 {
		t.Errorf("Unable to find files")
	}

	if err != nil {
		t.Errorf("Unable to find files")
	}

}

func Test_fsFinderExcludeDirs(t *testing.T) {
	fsFinder := NewService("../../../test/fixtures", []string{"subdir"})
	files, err := fsFinder.Find()

	if len(files) < 1 {
		t.Errorf("Unable to find files")
	}

	if err != nil {
		t.Errorf("Unable to find files")
	}
}

func Test_fsFinderPathNoExist(t *testing.T) {
	fsFinder := NewService("/bad/path", nil)
	_, err := fsFinder.Find()

	if err == nil {
		t.Errorf("Error not returned")
	}
}
