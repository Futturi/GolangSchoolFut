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
				//TODO доделать логику добавления файла
				lesson.POST("/", h.CreateHomework) // создать дз
				//lesson.POST("/", func(ctx *gin.Context) {}) // создать дз
			}
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

//TODO: Добавить в таблицу про дз + в структуру по дз
// поля связанные с выполнением дз(бул переменная)
// добавить общий get запрос на котором будет инфа про школу
// инфу буду хранить в редисе(как кэш собстна)
// сделать у каждого прользователя время до которого он может испол0ьзовать
// урок, если дз по уроку не выполнено, то у юзера -сердце, если
// количество сердец < 0, юзер вылетает с курса(заблокать)
// попробовать добавить оплату, пока даже хз как это делается
// мб добавить микросервис oauth для понимания микр арх и свзяи двух микр
// k8s и тд
//добавить логирование
// отправка писем про подтверждение email(разобраться)
