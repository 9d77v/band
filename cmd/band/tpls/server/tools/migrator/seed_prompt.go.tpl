package main

import (
	"encoding/json"
	"log"

	aigcDo "hhy-services/apps/aigc/infrastructure/persistence/do"

	"github.com/9d77v/band/pkg/stores/orm/base"
	"github.com/9d77v/band/pkg/stores/orm/impl/postgres"
	"gorm.io/gorm"
)

// --- 种子数据: AIGC Prompt ---

// mustMarshal 将 v 序列化为 JSON 字符串，失败时 panic
func mustMarshal(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Panicf("JSON 序列化失败: %s", err)
	}
	return string(b)
}

// defaultPrompts 默认 AIGC 提示词种子数据
func defaultPrompts() []*aigcDo.AigcPrompt {
	return []*aigcDo.AigcPrompt{
		{
			Model: base.Model{ID: 1},
			Code:  "AI_TRANSLATE",
			Name:  "AI翻译",
			Content: mustMarshal(map[string]any{
				"instruction": "请将input中的{{.SourceLang}}内容翻译成自然流畅的{{.TargetLang}},{{.Style}},确保保留原文的段落结构",
				"output":      "输出翻译后的内容,不添加任何注释",
				"input":       "{{.Text}}",
			}),
		},
		{
			Model: base.Model{ID: 2},
			Code:  "AI_GENERATE_FFMPEG_COMMAND",
			Name:  "AI生成FFMpeg命令",
			Content: `你是一个可以将自然语言翻译为 FFmpeg 命令的助手。你的任务是：
1. 根据用户的自然语言指令 ` + "{{.UserCommand}}" + `，结合以下补充信息生成一条完整的 FFmpeg 命令：
   - 当前输入视频文件名：input.mp4
   - 视频总时长：{{.Duration}}
   - 当前画面所在时间戳：{{.CurrentTime}}
2. 如果用户的描述不明确或无法解析为合法的 FFmpeg 命令，请返回以下内容：
   - 一个明确的提示，例如："无法根据输入生成 FFmpeg 命令，请尝试更清晰地描述目标。"
3. 只输出FFmpeg 命令或提示信息，不要提供额外的解释。
4. 如果需要用到字体文件，fontfile=default.ttf
5. 如果输出图片，默认jpeg格式
6. 默认不使用任何编码器，且音视频同步
7. 输出结果为ffmpeg命令字符串，不包含反引号`,
		},
		{
			Model:   base.Model{ID: 3},
			Code:    "AI_OCR",
			Name:    "AI图像识别",
			Content: `Read all the text in the image.`,
		},
		{
			Model:   base.Model{ID: 4},
			Code:    "AI_WRITING",
			Name:    "AI写作",
			Content: `帮我写文章, {{.Style}}, {{.Length}}, 主题和要求： {{.Text}}`,
		},
		{
			Model:   base.Model{ID: 5},
			Code:    "AI_WRITING_POLISH",
			Name:    "AI写作润色",
			Content: `润色以下段落, {{.Style}}, 只输出润色后的内容,不要包含任何关于润色选择的解释或额外信息。我的输入为： {{.Text}}`,
		},
		{
			Model:   base.Model{ID: 6},
			Code:    "AI_WRITING_SUMMARY",
			Name:    "AI写作总结",
			Content: `阅读并理解下面提供的文本内容，然后给出一个简洁且全面的总结。在总结中，确保保留原文中的关键信息和核心观点，并尽量减少不必要的细节。总结应当逻辑清晰、易于理解，并能够准确反映原文的主要论点和结论,  {{.Style}}。我的输入为： {{.Text}}`,
		},
		{
			Model: base.Model{ID: 7},
			Code:  "AI_WRITING_RED_NOTE",
			Name:  "AI写小红书笔记",
			Content: `请撰写一篇适合发布在小红书上的笔记, 
 {{.Style}}, {{.Length}}主题和要求如下：{{.Text}}。

注意事项：
- 包含一个不超过20字的标题和正文内容。
- 不使用Markdown格式，但可以并鼓励适当使用emoji。
- 避免提及具体网站链接和硬广告，确保内容真实有价值。
-  分段描述，每段围绕一个中心点，保持段落简洁明了，易于阅读。
- 在正文末尾自然地添加3到5个相关标签，直接作为正文的一部分。

目标是生成符合小红书风格且吸引读者注意的内容。`,
		},
		{
			Model:   base.Model{ID: 8},
			Code:    "AI_IMAGE_QA",
			Name:    "图片问答",
			Content: `{{.Text}}`,
		},
		{
			Model: base.Model{ID: 9},
			Code:  "AI_WRITING_DRAW_PROMPT",
			Name:  "AI绘画提示词",
			Content: mustMarshal(map[string]any{
				"instruction": "请根据input中的内容，使用补全规则生成合适的AI绘画提示词",
				"completion_rules": map[string]string{
					"style":       "若未检测到风格关键词，则根据主题关键词匹配默认风格（如自然→吉卜力/科幻→赛博朋克/神话→古典油画）",
					"composition": "自动补充视角（广角/仰视/微距）和层次（前景/背景元素）",
					"color":       "根据主题氛围补充主色调和光影（如战斗→暗红橙光/宁静→莫兰迪色）",
					"details":     "为物体添加合理材质和动态效果（如金属反光、粒子飞溅）",
					"tech":        "默认添加'8K超高清，锐化细节'",
					"exclusion":   "默认排除'低质量，畸变，水印'",
				},
				"output": "仅输出整合补全内容后的完整提示词文本，不要添加任何注释",
				"input":  "{{.Text}}",
			}),
		},
		{
			Model: base.Model{ID: 10},
			Code:  "AI_WRITING_VIDEO_PROMPT",
			Name:  "AI视频提示词",
			Content: mustMarshal(map[string]any{
				"instruction": "根据input内容生成专业AI视频提示词，自动补全影视化要素",
				"completion_rules": map[string]any{
					"motion": map[string]string{
						"type":       "未指定时智能匹配推/拉/摇/移/升降/手持运镜",
						"transition": "按主题添加转场（动作→冲击波特效/情感→柔光渐变）",
					},
					"subject": map[string]string{
						"detail":  "强化材质表现（丝绸飘动/金属反光）和微动作（发丝飘动/眨眼频率）",
						"physics": "自动添加布料模拟/流体动力学效果",
					},
					"scene": map[string]string{
						"environment": "生成互动元素（风吹草动/雨滴涟漪/动态投影）",
						"depth":       "智能分层（浅景深虚化/多层视差滚动）",
					},
					"cinematic": map[string]string{
						"framing": "自动组合景别（大特写→全景切换）",
						"angle":   "智能生成视角（蚂蚁视角→上帝视角过渡）",
					},
					"style": map[string]string{
						"color":    "未指定时匹配调色方案（科幻→青橙对比/武侠→水墨晕染）",
						"lighting": "自动构建光影层次（体积光雾/霓虹光污染）",
					},
				},
				"output": "仅输出整合补全内容后的完整提示词文本，不要添加任何注释",
				"input":  "{{.Text}}",
			}),
		},
		{
			Model: base.Model{ID: 11},
			Code:  "AI_WRITING_CONVERSATION_SCRIPT",
			Name:  "AI会话脚本生成提示词",
			Content: mustMarshal(map[string]any{
				"instruction": "请根据input中的内容，用户口语CERF等级,用户感兴趣的话题，背景信息，用户期望生成的对话轮数，用户重点关注的语法，生成符合output示例的json格式的,自然流畅的英语口语会话脚本，输出内容都是英文，标题长度不超过50字，角色名称随机，words为会话中出现的，符合当前CERF等级的不低于4个且不超过10个的单词或固定搭配",
				"input": map[string]string{
					"level":    "{{.Level}}",
					"topic":    "{{.Topic}}",
					"scenario": "{{.Scenario}}",
					"rounds":   "{{.Rounds}}",
					"grammar":  "{{.Grammar}}",
				},
				"output": map[string]any{
					"title": "Talking about preject",
					"content": []map[string]string{
						{"name": "Alice", "sentence": "Hi Bob, I've just finished reviewing our project report. Have you had a chance to look at it too?"},
						{"name": "Bob", "sentence": "Yes, Alice. I have gone through it this morning. It looks quite comprehensive."},
						{"name": "Alice", "sentence": "That's good to hear. Since we started this project, we have faced several challenges, haven't we?"},
						{"name": "Bob", "sentence": "Absolutely. But thanks to everyone's hard work, we have managed to overcome them all."},
					},
					"key_phrases": []string{"colleague", "manage to", "comprehensive", "challenge", "overcome"},
				},
			}),
		},
	}
}

// seedPrompts 初始化 AIGC 提示词种子数据
func seedPrompts(db *postgres.PgDB) {
	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, prompt := range defaultPrompts() {
			var existing aigcDo.AigcPrompt
			result := tx.Model(&aigcDo.AigcPrompt{}).Where("code = ?", prompt.Code).First(&existing)
			if result.Error == nil {
				continue
			}
			if result.Error != gorm.ErrRecordNotFound {
				return result.Error
			}
			if err := tx.Create(prompt).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("AIGC 提示词种子数据创建失败: %s", err)
	}
	log.Println("AIGC 提示词种子数据创建成功")
}
