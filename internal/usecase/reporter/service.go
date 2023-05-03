package reporter

import (
	"fmt"

	"github.com/Boeing/config-file-validator/internal/entity"
)

type InteractorI interface {
	Print(reports []entity.Report) error
}

type Service struct{}

func NewService(reporter string) (InteractorI, error) {
	switch reporter {
	case "json":
		return JsonReporter{}, nil
	case "standard":
		return StdoutReporter{}, nil
	default:
		return nil, fmt.Errorf("unsupported reporter type: %s", reporter)
	}
}
