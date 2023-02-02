package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func GetJsonTestRequestResponse(app *fiber.App, method string, url string, reqBody any) (code int, respBody map[string]any, err error) {
	bodyJson := []byte("")
	if reqBody != nil {
		bodyJson, err = json.Marshal(reqBody)
	}
	req := httptest.NewRequest(method, url, bytes.NewReader(bodyJson))
	resp, err := app.Test(req, 10)
	code = resp.StatusCode
	// If error we're done
	if err != nil {
		return
	}
	// If no body content, we're done
	if resp.ContentLength == 0 {
		return
	}
	bodyData := make([]byte, resp.ContentLength)
	_, _ = resp.Body.Read(bodyData)
	err = json.Unmarshal(bodyData, &respBody)
	return
}
