package finder

import (
	"github.com/Boeing/config-file-validator/internal/entity"
	"github.com/Boeing/config-file-validator/internal/usecase/validator"
)

type InteractorI interface {
	Find() ([]entity.FileMetadata, error)
}

func NewService(search_path *string, exclude []string) InteractorI {
	file_types := init_file_types()

	return &Service{
		FileTypes:  file_types,
		SearchPath: search_path,
		Exclude:    exclude,
	}
}

type Service struct {
	FileTypes  []entity.FileType
	SearchPath *string
	Exclude    []string
}

func init_file_types() []entity.FileType {
	var JsonFileType = entity.FileType{
		Name:       "json",
		Extensions: []string{"json"},
		Validator:  validator.NewJSONValidator(),
	}

	// Instance of the FileType object to
	// represent a YAML file
	var YamlFileType = entity.FileType{
		Name:       "yaml",
		Extensions: []string{"yml", "yaml"},
		Validator:  validator.NewYAMLValidator(),
	}

	// Instance of FileType object to
	// represent a XML file
	var XmlFileType = entity.FileType{
		Name:       "xml",
		Extensions: []string{"xml"},
		Validator:  validator.NewXMLValidator(),
	}

	// Instance of FileType object to
	// represent a Toml file
	var TomlFileType = entity.FileType{
		Name:       "toml",
		Extensions: []string{"toml"},
		Validator:  validator.NewTOMLValidator(),
	}

	// An array of files types that are supported
	// by the validator
	return []entity.FileType{
		JsonFileType,
		YamlFileType,
		XmlFileType,
		TomlFileType,
	}
}
