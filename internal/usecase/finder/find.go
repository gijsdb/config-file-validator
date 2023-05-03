package finder

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Boeing/config-file-validator/internal/entity"
)

func (s *Service) Find() ([]entity.FileMetadata, error) {
	var matchingFiles []entity.FileMetadata

	// check that the path exists before walking it or the error returned
	// from filepath.Walk will be very confusing and undescriptive
	if _, err := os.Stat(*s.SearchPath); os.IsNotExist(err) {
		return nil, err
	}

	err := filepath.WalkDir(*s.SearchPath, func(path string, dirEntry fs.DirEntry, err error) error {
		// determine if directory is in the excludeDirs list
		for _, dir := range s.Exclude {
			if dirEntry.IsDir() && dirEntry.Name() == dir {
				err := filepath.SkipDir
				if err != nil {
					return err
				}
			}
		}

		if !dirEntry.IsDir() {
			walkFileExtension := filepath.Ext(path)

			for _, fileType := range s.FileTypes {
				for _, extension := range fileType.Extensions {
					// filepath.Ext() returns the extension name with a dot
					// so it needs to be prepended to the FileType extension
					// in order to match
					if ("." + extension) == walkFileExtension {
						fileMetadata := entity.FileMetadata{dirEntry.Name(), path, fileType}
						matchingFiles = append(matchingFiles, fileMetadata)
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return matchingFiles, nil
}
