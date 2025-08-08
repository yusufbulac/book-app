package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/yusufbulac/byfood-case/backend/internal/routes"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
	"github.com/yusufbulac/byfood-case/backend/pkg/middleware"
)

type okResp struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

func setupURLApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.FiberErrorHandler(),
	})
	svc := service.NewUrlService()
	v1 := app.Group("/api/v1")
	routes.RegisterUrlRoutes(v1, svc)
	return app
}

func TestURL_Transform_Canonical_Success(t *testing.T) {
	app := setupURLApp()

	body := []byte(`{"url":"https://example.com/test/?q=1#frag","operation":"canonical"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/url/transform", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var ok okResp
	assert.NoError(t, json.NewDecoder(res.Body).Decode(&ok))

	var data struct {
		ProcessedURL string `json:"processed_url"`
	}
	assert.NoError(t, json.Unmarshal(ok.Data, &data))
	assert.Equal(t, "https://example.com/test", data.ProcessedURL)
}

func TestURL_Transform_InvalidPayload(t *testing.T) {
	app := setupURLApp()

	// invalid: no url
	body := []byte(`{"operation":"canonical"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/url/transform", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}
