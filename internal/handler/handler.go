package handler

import (
	"github.com/Futturi/GolangSchoolProject/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{service: serv}
}

func (h *Handler) InitRoutes() *gin.Engine {
	handler := gin.New()
	teacher := handler.Group("/admin")
	{
		auth := teacher.Group("/auth")
		{
			auth.POST("/signup", h.SignUp)
			auth.POST("/signin", h.SingIn)
			auth.POST("/refresh", h.Refresh)
		}
		lessons := teacher.Group("/lessons", h.CheckIdentity)
		{
			lessons.GET("/", h.GetAllLessonsTeacher)
			lessons.POST("/", h.CreateLesson)
			lessons.DELETE("/:lesson_id", h.DeleteLesson)
			// 	lesson := lessons.Group("/:lesson_id")
			// 	{
			// 		lesson.GET("/", func(ctx *gin.Context) {})
			// 		lesson.POST("/", func(ctx *gin.Context) {}) // создать урок
			// 		lesson.POST("/", func(ctx *gin.Context) {}) // создать дз
			// 	}
		}
	}
	// user := handler.Group("/user")
	// {
	// 	auth := user.Group("/auth")
	// 	{
	// 		auth.POST("/signup", func(ctx *gin.Context) {})
	// 		auth.POST("/signin", func(ctx *gin.Context) {})
	// 	}
	// 	lessons := user.Group("/lessons")
	// 	{
	// 		lessons.GET("/", func(ctx *gin.Context) {})
	// 		lessons.GET("/:lesson_id", func(ctx *gin.Context) {})  // получить дз + получить сам урок
	// 		lessons.POST("/:lesson_id", func(ctx *gin.Context) {}) // отправить выполненное дз на проверку
	// 	}
	// }
	return handler
}

//TODO: добавить в бд поля для дз, выполнено оно или нет + соединить дз с
// с таблицей урока, в дз хранить чисто выполнено дз или нет
