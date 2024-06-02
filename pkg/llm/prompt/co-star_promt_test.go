package prompt

import "testing"

func TestCoStarPromt_ToString(t *testing.T) {
	type fields struct {
		Context   string
		Objective string
		Style     string
		Tone      string
		Audience  string
		Response  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"costar示例", fields{
			Context:   "Context",
			Objective: "Objective",
			Style:     "Style",
			Tone:      "Tone",
			Audience:  "Audience",
			Response:  "Response",
		}, `### CONTEXT（上下文） ###
Context
### OBJECTIVE（目标） ###
Objective
### STYLE（风格） ###
Style
### TONE（语调） ###
Tone
### AUDIENCE（受众） ###
Audience
### RESPONSE（响应） ###
Response`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CoStarPromt{
				Context:   tt.fields.Context,
				Objective: tt.fields.Objective,
				Style:     tt.fields.Style,
				Tone:      tt.fields.Tone,
				Audience:  tt.fields.Audience,
				Response:  tt.fields.Response,
			}
			if got := c.ToString(); got != tt.want {
				t.Errorf("CoStarPromt.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
