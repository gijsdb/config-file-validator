package finder

import (
	"github.com/Boeing/config-file-validator/internal/entity"
	"github.com/Boeing/config-file-validator/internal/usecase/validator"
)

type InteractorI interface {
	Find() ([]entity.FileMetadata, error)
}

func NewService(search_path string, exclude []string) InteractorI {
	return &Service{
		FileTypes:  init_file_types(),
		SearchPath: search_path,
		Exclude:    exclude,
	}
}

type Service struct {
	FileTypes  []entity.FileType
	SearchPath string
	Exclude    []string
}

func init_file_types() []entity.FileType {
	return []entity.FileType{
		{
			Name:       "json",
			Extensions: []string{"json"},
			Validator:  validator.NewJSONValidator(),
		},
		{
			Name:       "yaml",
			Extensions: []string{"yml", "yaml"},
			Validator:  validator.NewYAMLValidator(),
		},
		{
			Name:       "xml",
			Extensions: []string{"xml"},
			Validator:  validator.NewXMLValidator(),
		},
		{
			Name:       "toml",
			Extensions: []string{"toml"},
			Validator:  validator.NewTOMLValidator(),
		},
	}
}
