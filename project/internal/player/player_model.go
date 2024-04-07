package player

type PlayerModel struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

type PlayerInput struct {
	Name    string
	Age     int
	Country string
}
