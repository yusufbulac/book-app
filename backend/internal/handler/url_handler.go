package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusufbulac/byfood-case/backend/internal/dto"
	"github.com/yusufbulac/byfood-case/backend/internal/service"
	"github.com/yusufbulac/byfood-case/backend/pkg/errorhandler"
	"github.com/yusufbulac/byfood-case/backend/pkg/response"
	"github.com/yusufbulac/byfood-case/backend/pkg/validator"
)

type UrlHandler struct {
	service service.UrlService
}

func NewUrlHandler(service service.UrlService) *UrlHandler {
	return &UrlHandler{service: service}
}

// Transform godoc
// @Summary     Transform a URL
// @Description Apply canonical, redirection, or all transformations to a URL
// @Tags        url
// @Accept      json
// @Produce     json
// @Param       payload body dto.UrlProcessRequest true "URL request payload"
// @Success     200 {object} dto.UrlProcessResponse
// @Failure     400 {object} response.ErrorResponse
// @Failure     500 {object} response.ErrorResponse
// @Router      /url/transform [post]
func (h *UrlHandler) Transform(c *fiber.Ctx) error {
	var input dto.UrlProcessRequest
	if err := c.BodyParser(&input); err != nil {
		return errorhandler.BadRequest("INVALID_PAYLOAD", "Failed to parse request body")
	}

	if err := validator.ValidateStruct(input); err != nil {
		return err
	}

	result, err := h.service.ProcessURL(input)
	if err != nil {
		return err
	}

	return response.Success(c, result)
}
