package qywx

type TextMsg struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
}

type Text struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list"`
	MentionedMobileList []string `json:"mentioned_mobile_list"`
}

func (c *QyWx) Text(text Text) error {
	params := &TextMsg{
		Msgtype: "text",
		Text:    text,
	}
	return c.Send(params)
}
