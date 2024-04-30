package interfaces

import (
	"database/sql"
	"net/http"

	"github.com/39shin52/attendanceApp/app/config"
	"github.com/39shin52/attendanceApp/app/domain/repository/transaction"
	"github.com/39shin52/attendanceApp/app/infrastructure"
	"github.com/39shin52/attendanceApp/app/interfaces/handler"
	"github.com/39shin52/attendanceApp/app/usecase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     *sql.DB
	Router *gin.Engine
}

// gin使ってみる

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (s *Server) Init() error {
	conn, err := config.CreateDBConnection()
	if err != nil {
		return err
	}

	s.db = conn
	s.Route()

	return nil
}

func (s *Server) Route() {
	// 依存性注入
	tx := transaction.NewTxRepository(s.db)

	userRepository := infrastructure.NewUserRepository(s.db)

	attendanceRepository := infrastructure.NewAttendanceRepository(s.db)
	attendanceUsecase := usecase.NewAttendanceUsecase(attendanceRepository, userRepository, *tx)
	attendanceHandler := handler.NewAttendanceHandler(*attendanceUsecase)

	s.Router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	s.Router.POST("/attendance/", attendanceHandler.RegisterAttendance)
}
