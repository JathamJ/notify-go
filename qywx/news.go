package qywx

type NewsMsg struct {
	Msgtype string `json:"msgtype"`
	News    News   `json:"news"`
}

type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

type News struct {
	Articles []Articles `json:"articles"`
}

func (c *QyWx) News(articles []Articles) error {
	params := &NewsMsg{
		Msgtype: "news",
		News: News{
			Articles: articles,
		},
	}
	return c.Send(params)
}
