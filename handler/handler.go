package handler

import (
	"net/http"

	"example.com/test/model"
	"example.com/test/model/response"
	"example.com/test/repository"
	"example.com/test/repository/repo_impl"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	Repo repository.Repo
}

func (h *Handler) HandlerSum(c echo.Context) error {
	repo := repo_impl.NewRepo()
	h.Repo = repo
	if h.Repo == nil {
		return c.JSON(http.StatusInternalServerError, response.Noti{
			Message: "Repository is not initialized",
		})
	}
	req := model.Input{}

	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Noti{
			Message: "Cant get value",
		})
	}
	res, err := h.Repo.Sum(c.Request().Context(), req)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Noti{
			Message: "Cant get value",
		})
	}
	return c.JSON(http.StatusBadRequest, res)
}
