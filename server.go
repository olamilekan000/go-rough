package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
	"github.com/qinains/fastergoding"

	"go-ml-app/src/helpers"
)

func main() {

	fastergoding.Run()

	engine := handlebars.New("./src/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	todosm := []string{
		"Wash",
		"Read Manga",
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"data": todosm,
		})
	})

	app.Post("/add-task", func(c *fiber.Ctx) error {

		val := c.FormValue("task")

		helpers.AddTasks(&todosm, val)

		return c.Render("add", fiber.Map{
			"data": todosm,
		})
	})

	app.Get("/delete/:index", func(c *fiber.Ctx) error {

		val, _ := strconv.Atoi(c.Params("index"))

		todosm = helpers.Remove(todosm, val)

		return c.Redirect("/")
	})

	log.Fatal(app.Listen(":4000"))

}
