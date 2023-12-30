package login

type Validate interface {
}

type validate struct {
}

func NewValidate() Validate {
	return &validate{}
}
