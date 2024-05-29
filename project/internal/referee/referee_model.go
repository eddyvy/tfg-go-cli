package referee

type Model struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Input struct {
	Name    string
	Country string
}
