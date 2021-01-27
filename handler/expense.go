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

	expenses, count, err = h.expenseStore.List(10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, response.ExpenseListResponse(h.userStore, expenses, count))
}

func (h *Handler) UserExpenses(c echo.Context) error {
	var (
		expenses []model.Expense
		teamName string
		count    int
		err      error
	)
	userID := userIDFromToken(c)
	teamName = c.Param("teamname")
	expenses, count, err = h.expenseStore.ListByUser(userID, 10, teamName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, response.ExpenseListResponse(h.userStore, expenses, count))
}

func (h *Handler) TeamExpenses(c echo.Context) error {
	var (
		expenses []model.Expense
		teamName string
		count    int
		err      error
	)

	teamName = c.Param("teamname")
	expenses, count, err = h.expenseStore.ListByTeam(10, teamName)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.JSON(http.StatusOK, response.ExpenseListResponse(h.userStore, expenses, count))
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

func (h *Handler) CalculateExpenses(c echo.Context) error {
	var calculations []model.Calculation
	userID := userIDFromToken(c)
	teamName := c.Param("teamname")
	team, users, err := h.userStore.TeamUsersList(teamName)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	calculations, err = h.expenseStore.CalCulateExpenses(calculations, &team, users)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	var usersCals *model.Calculation
	for _, calc := range calculations {
		if calc.UserID == userID {
			usersCals = &calc
		}
	}
	return c.JSON(http.StatusOK, response.NewSingleCalculationResponse(c, usersCals))
}

func (h *Handler) UpdateCalculation(c echo.Context) error {
	// teamName := c.Param("teamname")
	slugStr := c.Param("slug")
	slug, err := uuid.Parse(slugStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, util.NewError(err))
	}
	calculation, err := h.expenseStore.GetCalculationBySlug(slug)
	req := &request.CalculationUpdateRequest{}
	if err = req.Bind(c, calculation); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	err = h.expenseStore.UpdateCalculation(calculation)
	return c.JSON(http.StatusOK, response.NewSingleCalculationResponse(c, calculation))
}

func (h *Handler) Calculations(c echo.Context) error {
	var calculations []model.Calculation
	teamName := c.Param("name")
	err := h.expenseStore.CalculationsList(teamName, calculations)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, util.NewError(err))
	}
	return c.JSON(http.StatusOK, response.NewCalculationsListResponse(c, calculations))
}
