package prompt

import (
	"bytes"
	"html/template"
	"log"
)

type CoStarPromt struct {
	Context   string
	Objective string
	Style     string
	Tone      string
	Audience  string
	Response  string
}

const defaultTemplate = `# 上下文 #
{{.Context}}
# 目标 #
{{.Objective}}
# 风格 #
{{.Style}}
# 语调 #
{{.Tone}}
# 受众 #
{{.Audience}}
# 响应 #
{{.Response}}`

func (c CoStarPromt) ToString() string {
	tpl := template.Must(template.New("first").Parse(defaultTemplate))
	var buf bytes.Buffer
	err := tpl.Execute(&buf, c)
	if err != nil {
		log.Println(err)
	}
	return buf.String()
}

func (c CoStarPromt) Format(tplStr string) string {
	tpl := template.Must(template.New("custom").Parse(tplStr))
	var buf bytes.Buffer
	err := tpl.Execute(&buf, c)
	if err != nil {
		log.Println(err)
	}
	return buf.String()
}
