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
	handler := gin.Default()
	handler.GET("/", h.Info) //инфа о школе, которая хранится в кэше(редисе)
	// кол-во учеников, кол-во уроков
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

//TODO:
// добавить общий get запрос на котором будет инфа про школу
// инфу буду хранить в редисе(как кэш собстна)
// добавить логирование
// отправка писем про подтверждение email(разобраться)
