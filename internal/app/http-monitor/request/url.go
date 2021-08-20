package request

type Url struct {
	UrlBody   string `json:"url"`
	Threshold int    `json:"threshold"`
}

func NewUrl(body string, threshold int) *Url {
	url := &Url{UrlBody: body, Threshold: threshold}

	return url
}
