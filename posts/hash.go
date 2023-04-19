package posts

import (
	"TestTask/models"
	"TestTask/storage"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func Hash(c echo.Context) error {
	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	check := models.Check{}
	hashCheck := models.CheckHash{}
	db := storage.GetDBInstance()

	err := c.Bind(&check)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// Calculate the hash of the check
	hash, err := models.CalculateCheckHash(check)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error calculating hash for check")
	}

	println("Hash generated: ", hash)

	//Looking for same hash in table HashCheck
	found := db.Where("hash = ?", hash).First(&hashCheck)

	if found.Error != nil {
		if found.Error == gorm.ErrRecordNotFound {

			// Создаем новую запись
			if err := db.Create(&check).Error; err != nil {
				return c.String(http.StatusInternalServerError, "Error creating check")
			}

			// Создаем новый hash
			hCheck := models.CheckHash{Hash: hash, CheckID: check.ID}

			if err := db.Create(&hCheck).Error; err != nil {
				return c.String(http.StatusInternalServerError, "Error creating hash")
			}

			return c.String(http.StatusOK, "This check is unique! Successfully saved in the database!")
		}
		// Если произошла другая ошибка при выполнении запроса, возвращаем ошибку сервера
		return c.String(http.StatusInternalServerError, "Error executing query")
	}

	// Запись уже существует, возвращаем соответствующий ответ
	return c.JSON(http.StatusInternalServerError, "This check already exists (Found by hash comparing)")
}
