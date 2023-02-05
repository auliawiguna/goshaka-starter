package api_test

import (
	controller_v1 "goshaka/app/controllers/v1"
	"goshaka/database"
	performtest "goshaka/test"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var errConnectDbOnNote = database.Connect()

// Test code can get record from database
func TestCanGetNoteIndex(t *testing.T) {
	app := fiber.New()
	defer app.Shutdown()
	app.Get("/api/v1/notes", controller_v1.NoteIndex)

	type testArg struct {
		limit string
		page  string
		sort  string
	}

	code, body, err := performtest.GetJsonTestRequestResponse(app, "GET", "/api/v1/notes", testArg{"10", "1", "ID asc"}, "")
	assert.Nil(t, err)
	assert.Equal(t, 200, code)
	assert.NotEmpty(t, body["data"])
}

// Test code can get specific record from database
func TestCanGetNoteDetail(t *testing.T) {

	app := fiber.New()
	defer app.Shutdown()
	app.Get("/api/v1/notes/:id", controller_v1.NoteShow)

	code, body, err := performtest.GetJsonTestRequestResponse(app, "GET", "/api/v1/notes/2", nil, "")
	assert.Nil(t, err)
	assert.Equal(t, 200, code)
	assert.NotEmpty(t, body["data"])
}

func TestCanPostNewNoteUsingProperPayload(t *testing.T) {

	app := fiber.New()
	defer app.Shutdown()
	app.Post("/api/v1/notes", controller_v1.NoteStore)

	payload := map[string]string{
		"title":    "John Doe",
		"subtitle": "John",
		"text":     "John",
	}

	code, body, err := performtest.GetJsonTestRequestResponse(app, "POST", "/api/v1/notes", nil, payload)
	assert.Nil(t, err)
	assert.Equal(t, 200, code)
	assert.NotEmpty(t, body["data"])
}
