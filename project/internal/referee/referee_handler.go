package referee

import (
	"log"
	"net/http"

	"github.com/eddyvy/template/internal/parser"
	"github.com/gofiber/fiber/v2"
)

func HandleGetAll(c *fiber.Ctx) error {
	players, err := findAll()

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"message": err.Error()})
	}

	return c.JSON(&players)
}

func HandleGetOne(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := parser.StringToInt(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	player, err := findOne(id)

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	return c.JSON(&player)
}

func HandlePost(c *fiber.Ctx) error {
	player := new(Input)

	if err := c.BodyParser(player); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid input")
	}

	newPlayer, err := create(player)

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	return c.Status(201).JSON(&newPlayer)
}

func HandlePut(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := parser.StringToInt(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	player := new(Input)

	if err := c.BodyParser(player); err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid input")
	}

	updatedPlayer, err := update(id, player)

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(&updatedPlayer)
}

func HandleDelete(c *fiber.Ctx) error {
	idParam := c.Params("id")

	idInt, err := parser.StringToInt(idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid ID")
	}

	err = delete(idInt)

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"message": err.Error()})
	}

	return c.JSON(map[string]string{"message": "Resource deleted successfully"})
}
