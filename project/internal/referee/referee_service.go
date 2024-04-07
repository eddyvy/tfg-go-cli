package referee

import (
	"github.com/eddyvy/template/internal/database"
)

func readAllReferees() ([]*RefereeModel, error) {
	sqlStr := `SELECT * FROM referee`
	rows, err := database.DB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	referees := make([]*RefereeModel, 0)
	for rows.Next() {
		referee := new(RefereeModel)
		err := rows.Scan(&referee.Id, &referee.Name, &referee.Country)
		if err != nil {
			return nil, err
		}
		referees = append(referees, referee)
	}

	return referees, nil
}

func readOneReferee(id int) (*RefereeModel, error) {
	sqlStr := `SELECT * FROM referee WHERE id = $1`

	referee := new(RefereeModel)
	row := database.DB.QueryRow(sqlStr, id)
	err := row.Scan(&referee.Id, &referee.Name, &referee.Country)
	if err != nil {
		return nil, err
	}

	return referee, nil
}

func createReferee(referee *RefereeInput) (*RefereeModel, error) {
	sqlStr := `INSERT INTO referee (name, country) VALUES ($1, $2) RETURNING id`

	var id int

	err := database.DB.QueryRow(sqlStr, referee.Name, referee.Country).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &RefereeModel{
		Id:      id,
		Name:    referee.Name,
		Country: referee.Country}, nil
}

func updateReferee(id int, referee *RefereeInput) (*RefereeModel, error) {
	sqlStr := `UPDATE referee SET name = $1, country = $2 WHERE id = $3`

	_, err := database.DB.Exec(sqlStr, referee.Name, referee.Country, id)
	if err != nil {
		return nil, err
	}

	return &RefereeModel{
		Id:      id,
		Name:    referee.Name,
		Country: referee.Country}, nil
}

func deleteReferee(id int) error {
	sqlStr := `DELETE FROM referee WHERE id = $1`

	_, err := database.DB.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}
