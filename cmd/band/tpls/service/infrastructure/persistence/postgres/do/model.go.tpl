package do

import "{{.PKG_DIR}}/domain/entity"

type {{.ENTITY_UPPER}} struct {
}

func New{{.ENTITY_UPPER}}FromEntity({{.ENTITY_LOWER}} *entity.{{.ENTITY_UPPER}}) *{{.ENTITY_UPPER}} {
	//TODO
	return &{{.ENTITY_UPPER}}{}
}

func (do *{{.ENTITY_UPPER}}) To{{.ENTITY_UPPER}}() *entity.{{.ENTITY_UPPER}} {
	//TODO
	return &entity.{{.ENTITY_UPPER}}{
	}
}

func To{{.ENTITY_UPPER}}s({{.ENTITY_LOWER}}Dos []*{{.ENTITY_UPPER}}) []*entity.{{.ENTITY_UPPER}} {
	{{.ENTITY_LOWER}}s := make([]*entity.{{.ENTITY_UPPER}}, len({{.ENTITY_LOWER}}Dos))
	for i := range {{.ENTITY_LOWER}}Dos {
		{{.ENTITY_LOWER}}s[i] = {{.ENTITY_LOWER}}Dos[i].To{{.ENTITY_UPPER}}()
	}
	return {{.ENTITY_LOWER}}s
}
