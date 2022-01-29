package template

type coreParamsTemplate struct {
}

func (c *coreParamsTemplate) Text() []byte {
	return []byte(`package core

type Params struct {
	OffsetNo int64
	LimitNo  int64
}
`)
}

func CoreParamsTemplate() Template {
	return &coreParamsTemplate{}
}
