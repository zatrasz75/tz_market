package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zatrasz75/tz_market/configs"
	"zatrasz75/tz_market/internal/models"
	"zatrasz75/tz_market/internal/repository"
	"zatrasz75/tz_market/pkg/logger"
)

func RegisterHomeHandlers(mainGroup *gin.RouterGroup, l logger.LoggersInterface, repo *repository.Store, cfg *configs.Config) {
	en := &Server{l: l, repo: repo, cfg: cfg}

	mainGroup.POST("/building", en.saveBuildingHandler)
	mainGroup.GET("/building", en.searchBuildingHandler)
}

// saveBuildingHandler godoc
//
// @Tags Building
// @Summary Создайте новой записи
// @Description Принимает обязательные поля name , city , year_built , а так же не обязательные floors.
// @Accept  json
// @Produce  json
// @Param   user body models.Building true "Данные здания"
// @Success 201 {string} string "запись успешно добавлена"
// @Failure 400 {string} string "обязательные поля отсутствуют"
// @Failure 409 {string} string "название уже существует"
// @Failure 500 {string} string "Ошибка при добавлении данных"
// @Router /en/building [post]
func (s *Server) saveBuildingHandler(c *gin.Context) {
	var b models.Building

	err := json.NewDecoder(c.Request.Body).Decode(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не удалось проанализировать запрос JSON"})
		s.l.Error("не удалось проанализировать запрос JSON", err)
		return
	}

	if b.Name == "" || b.City == "" || b.YearBuilt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "обязательные поля отсутствуют"})
		s.l.Debug("обязательные поля отсутствуют")
		return
	}

	exists, err := s.repo.CheckBuilding(c, b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при проверке данных"})
		s.l.Error("ошибка при проверке данных", err)
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "название уже существует"})
		s.l.Error("название уже существует", err)
		return
	}

	err = s.repo.SaveBuilding(c, b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при добавлении данных"})
		s.l.Error("ошибка при добавлении данных", err)
		return
	}

	// Ответ при успешном создании
	c.JSON(http.StatusCreated, gin.H{"message": "запись успешно добавлена"})
	s.l.Info("запись добавлена успешно:", b)
}

// godoc godoc
//
// @Tags Building
// @Summary Возвращает список строений, с возможностью фильтрации по городу, году и кол-ву этажей(параметры не обязательные)
// @Description Принимает обязательные city, year_built, а так же не обязательные floors.
// @Accept  json
// @Produce  json
// @Param city query string false "по названию города"
// @Param year_built query integer false "по году сдачи"
// @Param floors query integer false "кол-ву этажей"
// @Success 200 {array} models.Building "Список домов"
// @Failure 400 {string} string "обязательные поля отсутствуют"
// @Failure 500 {string} string "Ошибка при получении данных"
// @Router /en/building [get]
func (s *Server) searchBuildingHandler(c *gin.Context) {
	city := c.Query("city")
	year := c.Query("year_built")
	floors := c.Query("floors")

	if city == "" || year == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "обязательные поля отсутствуют"})
		s.l.Debug("обязательные поля отсутствуют")
		return
	}

	// Преобразование year и floors в int
	yearInt, _ := strconv.Atoi(year)
	floorsInt, _ := strconv.Atoi(floors)

	// Получаем список строений из базы данных с учетом фильтров
	buildings, err := s.repo.GetBuildings(c.Request.Context(), city, yearInt, floorsInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
		s.l.Error("Ошибка при получении данных", err)
		return
	}

	// Возвращаем результаты в формате JSON
	c.JSON(http.StatusOK, buildings)
	s.l.Info("Запрос на поиск зданий выполнен успешно")
}
