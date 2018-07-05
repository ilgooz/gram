package telegram

type UpdateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int64   `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	ID   int64 `json:"message_id"`
	From struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		UserName     string `json:"username"`
		LanguageCode string `json:"language_code"`
	} `json:"from"`
	Chat struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		UserName  string `json:"username"`
		Type      string `json:"type"`
	} `json:"chat"`
	Date int64  `json:"date"`
	Text string `json:"text"`
}
