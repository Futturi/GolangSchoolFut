package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Futturi/GolangSchoolProject/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllLessonsTeacher(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.GetAllLessonsTeacher(userid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"lessons": result,
	})
}

func (h *Handler) CreateLesson(c *gin.Context) {
	var lesson models.Lesson
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = c.BindJSON(&lesson)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.CreateLesson(userId, lesson)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": result,
	})
}

func (h *Handler) DeleteLesson(c *gin.Context) {
	user, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = h.service.DeleteLesson(user, lesson_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
func (h *Handler) GetLesson(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	result, err := h.service.GetLesson(id, lesson_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}

func (h *Handler) UpdateLesson(c *gin.Context) {
	var fil models.UpdateLesson

	id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	//TODO доделать сохранение файлов, мы должны хранить инфу о файле в бд
	// путь(dir) названиe(f)

	err = c.ShouldBindJSON(&fil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.UpdateLesson(id, lesson_id, fil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) PutFile(c *gin.Context) {
	file, handl, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	defer file.Close()
	lesson_id, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	dir := "lesson_files/" + strconv.Itoa(lesson_id)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	f, err := os.OpenFile(filepath.Join(dir, handl.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	defer f.Close()
	name := f.Name()
	err = h.service.PutFile(name, lesson_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	_, err = io.Copy(f, file)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) CheckHomework(c *gin.Context) {
	var status models.CheckHom
	url := c.Request.URL.Path
	masurl := strings.Split(url, "/")
	lesson_id, err := strconv.Atoi(masurl[3])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	teacher_id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err = c.BindJSON(&status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	err = h.service.CheckHomework(teacher_id, lesson_id, status)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *Handler) GetHomework(c *gin.Context) {
	url := c.Request.URL.Path
	masurl := strings.Split(url, "/")
	lesson_id, err := strconv.Atoi(masurl[3])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	result, err := h.service.GetHomework(lesson_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, result)
}
