package shakespeare

// Translation model
type Translation struct {
	Success  Success `json:"success"`
	Contents Content `json:"contents"`
}

// Success model
type Success struct {
	Total int `json:"total"`
}

// Content model
type Content struct {
	Translated  string `json:"translated"`
	Text        string `json:"text"`
	Translation string `json:"translation"`
}

// translationRequest  model
type translationRequest struct {
	Text string `json:"text"`
}
