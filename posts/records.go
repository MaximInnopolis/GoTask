package posts

import (
	"TestTask/models"
	"TestTask/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func Records(c echo.Context) error {
	check := models.Check{}
	db := storage.GetDBInstance()

	err := c.Bind(&check)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// Пытаемся найти запись с заданными значениями
	found := db.Where("store_name = ? AND total = ? AND payment_method = ? AND tax = ?",
		check.StoreName, check.Total, check.PaymentMethod, check.Tax).First(&check)

	if found.Error != nil {
		if found.Error == gorm.ErrRecordNotFound {

			// Создаем новую запись
			if err := db.Create(&check).Error; err != nil {
				return c.String(http.StatusInternalServerError, "Error creating check")
			}

			return c.String(http.StatusOK, "This check is unique! Successfully saved in the database!")
		}

		// Если произошла другая ошибка при выполнении запроса, возвращаем ошибку сервера
		return c.String(http.StatusInternalServerError, "Error executing query")
	}
	// Запись уже существует, возвращаем соответствующий ответ
	return c.String(http.StatusInternalServerError, "This check already exists (Found by records comparing)")
}
