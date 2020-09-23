package items

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	URL string `json:"url"`
}

type Item struct {
	Id          string      `json:"id"`
	Seller      int64       `json:"seller"`
	Title       string      `json:"title"`
	Description Description `json:"description"`
	Picture     []Picture   `json:"picture"`
	VideoURL    string      `json:"video_url"`
	Price       float32     `json:"price"`
	Status      string      `json:"status"`
}
