package endpoints

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/MaciejTe/twitter/pkg/messenger"
	"github.com/MaciejTe/twitter/pkg/models"
)

// MessagesEndpoint aggregates all endpoint handlers into one function
func MessagesEndpoint(app fiber.Router, messenger messenger.Messenger) {
	app.Post("/messages", addMessage(messenger))
	app.Get("/messages", getMessages(messenger))
}

// addMessage is handler for inserting messages into database
func addMessage(messenger messenger.Messenger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody *models.Message
		err := c.BodyParser(&requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result, err := messenger.InsertMessage(requestBody)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		c.Status(http.StatusCreated)
		return c.JSON(&fiber.Map{
			"status": true,
			"count":  1,
			"result": result,
		})
	}
}

// getMessages is handler taking care of counting messages accoring to following query filters:
// tags - comma separated list of tags to search for
// from - start search date (RFC3339 standard, example: 2021-06-13T15:00:05Z)
// to - end search date (RFC3339 standard, example: 2021-06-13T15:00:05Z)
func getMessages(messenger messenger.Messenger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tags := c.Query("tags")
		from := c.Query("from")
		to := c.Query("to")
		filters, err := models.NewMessageFilter(tags, from, to)
		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(&fiber.Map{
				"error": err,
			})
		}
		count, err := messenger.FetchMessagesCount(*filters)
		if err != nil {
			_ = c.JSON(&fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"count":  count,
		})
	}
}
