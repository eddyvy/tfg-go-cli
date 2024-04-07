package referee

type RefereeModel struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type RefereeInput struct {
	Name    string
	Country string
}
