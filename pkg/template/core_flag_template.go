package template

type coreFlagTemplate struct {
}

func (c *coreFlagTemplate) Text() []byte {
	return []byte(`package core

const (
	FlagIgnore   = 0
	FlagInactive = 1
	FlagActive   = 2
)
`)
}

func CoreFlagTemplate() Template {
	return &coreFlagTemplate{}
}
