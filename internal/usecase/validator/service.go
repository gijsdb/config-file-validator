package validator

type InteractorI interface {
	Validate(b []byte) (bool, error)
}
