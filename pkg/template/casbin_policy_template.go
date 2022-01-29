package template

type casbinPolicyTemplate struct {
}

func (c *casbinPolicyTemplate) Text() []byte {
	return []byte(`p, anonymous, /swagger, (GET)|(OPTIONS)
p, anonymous, /v1/login, (GET)|(OPTIONS)
p, admin, /v1/*, (GET)|(POST)|(PUT)|(DELETE)|(OPTIONS)
p, user, /v1/user/me, (GET)|(OPTIONS)`)
}

func CasbinPolicyTemplate() Template {
	return &casbinPolicyTemplate{}
}
