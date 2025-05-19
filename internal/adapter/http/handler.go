package http

import (
	"net/http"
	"pongsatorn/go-http-template/internal/core/service"

	"github.com/labstack/echo/v4"
)

type HttpHandler struct {
	tplService *service.TemplateService
}

func NewHttpHandler(tplService *service.TemplateService) *HttpHandler {
	return &HttpHandler{
		tplService: tplService,
	}
}

// @Summary User Login
// @Description Logs in a user with their email and password to get an authentication token.
// @Tags auth
// @Accept json
// @Produce json
// @Param req body LoginRequest true "Login credentials"
// @Success 200 {object} domain.Token "token generated from created account"
// @Failure 400 {object} BaseResponse "Bad request (invalid email or password)"
// @Failure 401 {object} BaseResponse "Unauthorized (invalid credentials)"
// @Router /api/v1/login [post]
func (h *HttpHandler) Template(c echo.Context) error {
	type req struct {
	}

	if err := c.Bind(&req{}); err != nil {
		return c.JSON(http.StatusBadRequest, &req{})
	}

	if err := c.Validate(&req{}); err != nil {
		return c.JSON(http.StatusBadRequest, &req{})
	}

	_, err := h.tplService.Temp(c.Request().Context(), &req{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, nil)
}
