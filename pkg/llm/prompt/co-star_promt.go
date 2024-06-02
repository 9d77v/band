package prompt

type CoStarPromt struct {
	Context   string
	Objective string
	Style     string
	Tone      string
	Audience  string
	Response  string
}

func (c CoStarPromt) ToString() string {
	return `### CONTEXT（上下文） ###
` + c.Context + `
### OBJECTIVE（目标） ###
` + c.Objective + `
### STYLE（风格） ###
` + c.Style + `
### TONE（语调） ###
` + c.Tone + `
### AUDIENCE（受众） ###
` + c.Audience + `
### RESPONSE（响应） ###
` + c.Response
}
