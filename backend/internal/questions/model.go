package questions

type Question struct {
	ID         int      `json:"id"`
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	Difficulty int      `json:"difficulty"`
	Tags       []string `json:"tags"`
}
