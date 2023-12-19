package do

import "{{.PKG_DIR}}/internal/apps/{{.SERVICE_LOWER}}/domain/entity"

type {{.ENTITY_UPPER}}DO struct {
}

func New{{.ENTITY_UPPER}}FromEntity({{.ENTITY_LOWER}} *entity.{{.ENTITY_UPPER}}) *{{.ENTITY_UPPER}}DO {
	//TODO
	return &{{.ENTITY_UPPER}}DO{}
}

func (do *{{.ENTITY_UPPER}}DO) To{{.ENTITY_UPPER}}() *entity.{{.ENTITY_UPPER}} {
	//TODO
	return &entity.{{.ENTITY_UPPER}}{
	}
}

func To{{.ENTITY_UPPER}}s({{.ENTITY_LOWER}}Dos []*{{.ENTITY_UPPER}}DO) []*entity.{{.ENTITY_UPPER}} {
	{{.ENTITY_LOWER}}s := make([]*entity.{{.ENTITY_UPPER}}, len({{.ENTITY_LOWER}}Dos))
	for i := range {{.ENTITY_LOWER}}Dos {
		{{.ENTITY_LOWER}}s[i] = {{.ENTITY_LOWER}}Dos[i].To{{.ENTITY_UPPER}}()
	}
	return {{.ENTITY_LOWER}}s
}