package handler

import (
	"github.com/Futturi/GolangSchoolProject/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	cfg     service.Sender
}

func NewHandler(serv *service.Service, cfg service.Sender) *Handler {
	return &Handler{service: serv, cfg: cfg}
}

func (h *Handler) InitRoutes() *gin.Engine {
	handler := gin.Default()
	handler.GET("/", h.Info)
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
			lessons.PUT("/:lesson_id", h.UpdateLesson)
			lesson := lessons.Group("/:lesson_id")
			{
				lesson.GET("/", h.GetLesson)
				lesson.POST("/", h.PutFile)
				homework := lesson.Group("/homework")
				{
					homework.POST("/", h.CheckHomework)
					homework.GET("/", h.GetHomework)
				}
			}
		}
	}
	user := handler.Group("/user")
	{
		auth := user.Group("/auth")
		{
			auth.POST("/signup", h.SignUpUser)
			auth.POST("/signin", h.SignInUser)
			auth.POST("/refresh", h.RefreshUser)
			auth.GET("/:email_token", h.CheckToken)
		}
		lessons := user.Group("/lessons", h.CheckIdentityUser, h.CheckHealth)
		{
			lessons.GET("/", h.GetAllLessonsUser)
			lessons.GET("/:lesson_id", h.GetLessonUser)
			lessons.POST("/:lesson_id", h.SolveHomework)
		}
	}
	return handler
}

// добавить логирование
// отправка писем про подтверждение email(разобраться)
