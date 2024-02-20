package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userApp "user-app"
)

type createInput struct {
	Id        string `json:"id" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}

func (h *Handler) create(c *gin.Context) {
	var input createInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Age <= 0 || input.Age > 100 {
		newErrorResponse(c, http.StatusBadRequest, "invalid age value")
		return
	}

	recordingDate := getRecordingDate()

	user := userApp.User{
		Id:            input.Id,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Age:           input.Age,
		RecordingDate: recordingDate,
	}

	err := h.services.Create(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"record_date": recordingDate,
	})
}

type getUsersResponse struct {
	Count int64          `json:"count"`
	Data  []userApp.User `json:"users"`
}

func (h *Handler) getAll(c *gin.Context) {
	users, count, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUsersResponse{
		Count: count,
		Data:  users,
	})
}

func (h *Handler) getByDateAndAge(c *gin.Context) {
	var input userApp.SearchInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.StartDate > input.EndDate && input.EndDate != 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid date values")
		return
	}

	if input.MinAge > input.MaxAge && input.MaxAge != 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid age values")
		return
	}

	input = checkInput(input)

	users, count, err := h.services.GetByDateAndAge(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getUsersResponse{
		Count: count,
		Data:  users,
	})
}
