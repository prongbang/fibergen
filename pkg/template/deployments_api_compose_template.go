package template

type apiComposeTemplate struct {
	Project string
}

func (a *apiComposeTemplate) Text() []byte {
	return []byte(`version: '3.9'
services:
  ` + a.Project + `-api:
    image: ` + a.Project + `-api:1.0
    ports:
      - "9001:9001"
    environment:
      TZ: "Asia/Bangkok"
    networks:
      - ` + a.Project + `-network

networks:
  ` + a.Project + `-network:
    external: true`)
}

func DeploymentsApiComposeTemplate(project string) Template {
	return &apiComposeTemplate{
		Project: project,
	}
}
