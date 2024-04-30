package handler

import (
	"net/http"

	"github.com/39shin52/attendanceApp/app/interfaces/request"
	"github.com/39shin52/attendanceApp/app/usecase"
	"github.com/gin-gonic/gin"
)

type AttendanceHandler struct {
	attendanceUsecase usecase.AttendanceUsecase
}

func NewAttendanceHandler(attendanceUsecase usecase.AttendanceUsecase) *AttendanceHandler {
	return &AttendanceHandler{attendanceUsecase: attendanceUsecase}
}

func (ah *AttendanceHandler) RegisterAttendance(c *gin.Context) {
	newAttendanceRequest := new(request.AttendanceRegistrationRequest)
	if err := c.Bind(&newAttendanceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err := ah.attendanceUsecase.RegisterAttendance(c.Request.Context(), newAttendanceRequest.ID, newAttendanceRequest.Name, newAttendanceRequest.Time); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
