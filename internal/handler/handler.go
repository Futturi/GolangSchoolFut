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
	//handler.GET("/") //инфа о школе, которая хранится в кэше(редисе)
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
				lesson.PUT("/", h.PutFile)
				lesson.POST("/", h.CreateHomework)
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
		lessons := user.Group("/lessons", h.CheckIdentityUser)
		{
			lessons.GET("/", h.GetAllLessonsUser)
			// 		lessons.GET("/:lesson_id", func(ctx *gin.Context) {})  // получить дз + получить сам урок
			// 		lessons.POST("/:lesson_id", func(ctx *gin.Context) {}) // отправить выполненное дз на проверку
		}
	}
	return handler
}

//TODO:
// добавить общий get запрос на котором будет инфа про школу
// инфу буду хранить в редисе(как кэш собстна)
// сделать у каждого прользователя время до которого он может испол0ьзовать
// урок, если дз по уроку не выполнено, то у юзера -сердце, если
// количество сердец < 0, юзер вылетает с курса(заблокать)
// попробовать добавить оплату, пока даже хз как это делается
// добавить логирование
// отправка писем про подтверждение email(разобраться)
