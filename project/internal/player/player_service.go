package player

import (
	"github.com/eddyvy/template/internal/database"
)

func readAllPlayers() ([]*PlayerModel, error) {
	sqlStr := `SELECT * FROM player`
	rows, err := database.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := make([]*PlayerModel, 0)
	for rows.Next() {
		player := new(PlayerModel)
		err := rows.Scan(&player.Id, &player.Name, &player.Age, &player.Country)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}

	return players, nil
}

func readOnePlayer(id int) (*PlayerModel, error) {
	sqlStr := `SELECT * FROM player WHERE id = $1`

	player := new(PlayerModel)
	row := database.DB.QueryRow(sqlStr, id)
	err := row.Scan(&player.Id, &player.Name, &player.Age, &player.Country)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func createPlayer(player *PlayerInput) (*PlayerModel, error) {
	sqlStr := `INSERT INTO player (name, age, country) VALUES ($1, $2, $3) RETURNING id`

	var id int

	err := database.DB.QueryRow(sqlStr, player.Name, player.Age, player.Country).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &PlayerModel{
		Id:      id,
		Name:    player.Name,
		Age:     player.Age,
		Country: player.Country}, nil
}

func updatePlayer(id int, player *PlayerInput) (*PlayerModel, error) {
	sqlStr := `UPDATE player SET name = $1, age = $2, country = $3 WHERE id = $4`

	_, err := database.DB.Exec(sqlStr, player.Name, player.Age, player.Country, id)
	if err != nil {
		return nil, err
	}

	return &PlayerModel{
		Id:      id,
		Name:    player.Name,
		Age:     player.Age,
		Country: player.Country}, nil
}

func deletePlayer(id int) error {
	sqlStr := `DELETE FROM player WHERE id = $1`

	_, err := database.DB.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
