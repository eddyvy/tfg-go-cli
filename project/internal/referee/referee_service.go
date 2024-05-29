package referee

import (
	"strconv"

	"github.com/eddyvy/template/internal/database"
)

func findAll() ([]*Model, error) {
	sqlStr := `SELECT * FROM referee`
	rows, err := database.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	referees := make([]*Model, 0)
	for rows.Next() {
		referee := new(Model)
		err := rows.Scan(&referee.Id, &referee.Name, &referee.Country)
		if err != nil {
			return nil, err
		}
		referees = append(referees, referee)
	}

	return referees, nil
}

func findOne(id int) (*Model, error) {
	sqlStr := `SELECT * FROM referee WHERE id = $1`

	referee := new(Model)
	row := database.DB.QueryRow(sqlStr, id)
	err := row.Scan(&referee.Id, &referee.Name, &referee.Country)
	if err != nil {
		return nil, err
	}

	return referee, nil
}

func create(referee *Input) (*Model, error) {
	sqlStr := `INSERT INTO referee (name, country) VALUES ($1, $2) RETURNING id`

	var id int

	err := database.DB.QueryRow(sqlStr, referee.Name, referee.Country).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &Model{
		Id:      id,
		Name:    referee.Name,
		Country: referee.Country}, nil
}

func update(id int, referee *Input) (*Model, error) {
	sqlStr := `UPDATE referee SET name = $1, country = $2 WHERE id = $3`

	_, err := database.DB.Exec(sqlStr, referee.Name, referee.Country, id)
	if err != nil {
		return nil, err
	}

	return &Model{
		Id:      id,
		Name:    referee.Name,
		Country: referee.Country}, nil
}

func delete(id int) error {
	sqlStr := `DELETE FROM referee WHERE id = $1`

	_, err := database.DB.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}

func parseId(idParam string) (int, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}

	return id, nil
}
