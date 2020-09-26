package handler

import (
	"net/http"

	"github.com/HideBa/soroha-api/model"
	"github.com/HideBa/soroha-api/request"
	"github.com/HideBa/soroha-api/response"
	util "github.com/HideBa/soroha-api/utils"
	"github.com/google/uuid"
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

func (h *Handler) Expenses(c echo.Context) error {
	var (
		expenses []model.Expense
		count    int
		err      error
	)

	userID := userIDFromToken(c)
	expenses, count, err = h.expenseStore.List(10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, response.NewExponseListResponse(h.userStore, userID, expenses, count))
}

func (h *Handler) UpdateExpense(c echo.Context) error {
	slugStr := c.Param("slug")
	slugUUID, err := uuid.Parse(slugStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	expense, err := h.expenseStore.GetUserExpenseBySlug(userIDFromToken(c), slugUUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	if expense == nil {
		return c.JSON(http.StatusNotFound, util.NotFound())
	}

	req := &request.ExpenseUpdateRequest{}
	req.ConvertModelToRequest(expense)

	if err := req.Bind(c, expense); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}

	if err = h.expenseStore.UpdateExpense(expense); err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}

	return c.JSON(http.StatusOK, response.NewExpenseResponse(c, expense))
}

func (h *Handler) DeleteExpense(c echo.Context) error {
	slugStr := c.Param("slug")
	slugUUID, err := uuid.Parse(slugStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	expense, err := h.expenseStore.GetUserExpenseBySlug(userIDFromToken(c), slugUUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	if expense == nil {
		return c.JSON(http.StatusNotFound, util.NotFound())
	}

	err = h.expenseStore.DeleteExpense(expense)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}
