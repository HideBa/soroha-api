package handler

import (
	"net/http"

	"github.com/HideBa/soroha-api/model"
	"github.com/HideBa/soroha-api/request"
	"github.com/HideBa/soroha-api/response"
	util "github.com/HideBa/soroha-api/utils"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateExpense(c echo.Context) error {
	var expense model.Expense
	req := &request.ExpenseCreateRequest{}
	if err := req.Bind(c, &expense); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}

	expense.UserID = userIDFromToken(c)
	err := h.expenseStore.CreateExpense(&expense)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}

	return c.JSON(http.StatusCreated, response.NewExpenseResponse(c, &expense))
}
