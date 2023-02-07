package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

func HttpRequestForTest(app *fiber.App, method, url string, reqBody, reqString any) (code int, respBody map[string]any, err error) {
	bodyJson := []byte("")
	var req *http.Request
	if reqBody != nil {
		bodyJson, err = json.Marshal(reqBody)
		if err != nil {
			return
		}
		req = httptest.NewRequest(method, url, bytes.NewReader(bodyJson))
	} else if reqString != "" {
		jsonValue, err2 := json.Marshal(reqString)
		if err2 != nil {
			return
		}
		req = httptest.NewRequest(method, url, bytes.NewBuffer(jsonValue))
		req.Header.Add("Content-Type", "application/json")
	} else {
		bodyJson, err = json.Marshal(reqBody)
		req = httptest.NewRequest(method, url, nil)
	}
	resp, err := app.Test(req, 1000)
	if err != nil {
		return
	}
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
