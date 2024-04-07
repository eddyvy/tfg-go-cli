package player

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPlayers(c *fiber.Ctx) error {
	players, err := readAllPlayers()

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.JSON(&players)
}

func GetPlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	player, err := readOnePlayer(idInt)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(&player)
}

func CreatePlayer(c *fiber.Ctx) error {
	player := new(PlayerInput)

	if err := c.BodyParser(player); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid input")
	}

	newPlayer, err := createPlayer(player)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(201).JSON(&newPlayer)
}

func UpdatePlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	player := new(PlayerInput)

	if err := c.BodyParser(player); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid input")
	}

	updatedPlayer, err := updatePlayer(idInt, player)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(&updatedPlayer)
}

func DeletePlayer(c *fiber.Ctx) error {
	id := c.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	err = deletePlayer(idInt)

	if err != nil {
		log.Fatalln(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(map[string]string{"message": "Resource deleted successfully"})
}
