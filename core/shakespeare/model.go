package shakespeare

type Translation struct {
	Success  Success `json:"success"`
	Contents Content `json:"contents"`
}

type Success struct {
	Total int `json:"total"`
}

type Content struct {
	Translated  string `json:"translated"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}

type translationRequest struct {
	Text string `json:"text"`
}
