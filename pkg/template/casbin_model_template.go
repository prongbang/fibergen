package template

type casbinModelTemplate struct {
}

func (c *casbinModelTemplate) Text() []byte {
	return []byte(`[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch(r.obj, p.obj) || keyMatch2(r.obj, p.obj)) && (r.act == p.act || regexMatch(r.act, p.act))`)
}

func CasbinModelTemplate() Template {
	return &casbinModelTemplate{}
}
