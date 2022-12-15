package utils

import (
	models "ms-scaffold/api/models/domain"
	config "ms-scaffold/config/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.NewDB()
)

func GeneratePaginationForScaffold(c *gin.Context, m models.Scaffold) models.Pagination {
	// Initializing default
	var total int64
	db.Model(&m).Count(&total)
	limit := int(total)
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
