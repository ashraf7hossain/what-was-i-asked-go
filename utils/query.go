package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	Search    string `json:"search"`
	SearchBy  string `json:"search_by"`
	SearchStr string `json:"search_str"`
	OrderBy   string `json:"order_by"`
}

func ParseQueryParams(c *gin.Context) QueryParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.DefaultQuery("search", "")
	orderBy := c.DefaultQuery("order_by", "id") // Default sorting by ID

	return QueryParams{
		Page:    page,
		Limit:   limit,
		Search:  search,
		OrderBy: orderBy,
	}
}

func Paginate(c *gin.Context, total int, data interface{}) gin.H {
	pageParams := ParseQueryParams(c)
	return gin.H{
		"data":     data,
		"page":     pageParams.Page,
		"limit":    pageParams.Limit,
		"total":    total,
		"pages":    (total + pageParams.Limit - 1) / pageParams.Limit,
		"next":     pageParams.Page + 1,
		"previous": pageParams.Page - 1,
	}
}
