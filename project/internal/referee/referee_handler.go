package referee

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetReferees(c *fiber.Ctx) error {
	referees, err := readAllReferees()

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.JSON(&referees)
}

func GetReferee(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	referee, err := readOneReferee(idInt)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(&referee)
}

func CreateReferee(c *fiber.Ctx) error {
	referee := new(RefereeInput)

	if err := c.BodyParser(referee); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid input")
	}

	newReferee, err := createReferee(referee)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(201).JSON(&newReferee)
}

func UpdateReferee(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	referee := new(RefereeInput)

	if err := c.BodyParser(referee); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid input")
	}

	updatedReferee, err := updateReferee(idInt, referee)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(&updatedReferee)
}

func DeleteReferee(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	err = deleteReferee(idInt)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(map[string]string{"message": "Resource deleted successfully"})
}
