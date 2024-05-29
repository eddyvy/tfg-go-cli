package player

import (
	"github.com/eddyvy/template/internal/database"
)

func findAll() ([]*Model, error) {
	sqlStr := `SELECT * FROM player`
	rows, err := database.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := make([]*Model, 0)
	for rows.Next() {
		player := new(Model)
		err := rows.Scan(&player.Id, &player.Name, &player.Age, &player.Country)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}

func findOne(id int) (*Model, error) {
	sqlStr := `SELECT * FROM player WHERE id = $1`

	player := new(Model)
	row := database.DB.QueryRow(sqlStr, id)
	err := row.Scan(&player.Id, &player.Name, &player.Age, &player.Country)

	if err != nil {
		return nil, err
	}

	return player, nil
}

func create(player *Input) (*Model, error) {
	sqlStr := `INSERT INTO player (name, age, country) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := database.DB.QueryRow(sqlStr, player.Name, player.Age, player.Country).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &Model{
		Id:      id,
		Name:    player.Name,
		Age:     player.Age,
		Country: player.Country}, nil
}

func update(id int, player *Input) (*Model, error) {
	sqlStr := `UPDATE player SET name = $1, age = $2, country = $3 WHERE id = $4`

	_, err := database.DB.Exec(sqlStr, player.Name, player.Age, player.Country, id)
	if err != nil {
		return nil, err
	}

	return &Model{
		Id:      id,
		Name:    player.Name,
		Age:     player.Age,
		Country: player.Country}, nil
}

func delete(id int) error {
	sqlStr := `DELETE FROM player WHERE id = $1`

	_, err := database.DB.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
