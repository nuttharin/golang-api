package httpserver

import (
	"strings"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"} // Allow all origins
	config.AllowHeaders = []string{"*"}

	return cors.New(config)
}

func DbTransactionMiddleware(handler func(Context), db *gorm.DB) gin.HandlerFunc {
	// Set DB transaction
	return func(ctx *gin.Context) {
		tx := db.Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// set db_tx variable
		ctx.Set("db_tx", tx)

		// before request
		ctx.Next()

		// after request
		convertToGinHandler(handler)(ctx)
		if len(ctx.Errors) > 0 {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}

func Localization(setDefaultLang bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerLang := strings.ToLower(c.GetHeader("Content-Language"))
		lang := ""

		if headerLang == "th" || headerLang == "en" {
			lang = headerLang
		} else if setDefaultLang {
			lang = "th"
		}

		c.Set("lang", lang)
		c.Next()
	}
}
