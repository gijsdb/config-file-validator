package entity

import "github.com/Boeing/config-file-validator/internal/usecase/validator"

type Report struct {
	FileName        string
	FilePath        string
	IsValid         bool
	ValidationError error
}

type FileMetadata struct {
	Name     string
	Path     string
	FileType FileType
}

type FileType struct {
	Name       string
	Extensions []string
	Validator  validator.InteractorI
}
