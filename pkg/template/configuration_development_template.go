package template

type configurationDevelopmentTemplate struct {
}

func (c *configurationDevelopmentTemplate) Text() []byte {
	return []byte(`env: "development"

api:
  port: 9001

role:
  admin: "admin"
  user: "user"

jwt:
  secret: "secret"

casbin:
  model: "policy/model.conf"
  policy: "policy/policy.csv"

mongodb:
  host: "localhost"
  port: 27017
  database: "mongoDB"
  user: "root"
  pass: "password"
`)
}

func ConfigurationDevelopmentTemplate() Template {
	return &configurationDevelopmentTemplate{}
}
