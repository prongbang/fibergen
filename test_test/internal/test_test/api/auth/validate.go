package auth

type Validate interface {
}

type validate struct {
}

func NewValidate() Validate {
	return &validate{}
}
