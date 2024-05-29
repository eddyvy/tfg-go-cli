package player

type Model struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

type Input struct {
	Name    string
	Age     int
	Country string
}
