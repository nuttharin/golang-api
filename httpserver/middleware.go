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

// func ErrorHandlerMiddleware(languageBundles map[string]i18n.Bundle) gin.HandlerFunc {
// 	// return func(c *gin.Context) {
// 	// 	defer func() {
// 	// 		if r := recover(); r != nil {
// 	// 			c.JSON(http.StatusInternalServerError, &FailResponse{
// 	// 				Code:    http.StatusInternalServerError,
// 	// 				Status:  FailStatus,
// 	// 				Message: fmt.Sprintf("Panic: %v", r),
// 	// 			})
// 	// 		}
// 	// 	}()

// 	// 	c.Next()
// 	// 	if len(c.Errors) > 0 {

// 	// 		for _, err := range c.Errors {
// 	// 			if err.Err == gorm.ErrRecordNotFound {
// 	// 				c.JSON(http.StatusNotFound, &FailResponse{
// 	// 					Code:    http.StatusNotFound,
// 	// 					Status:  FailStatus,
// 	// 					Message: err.Err.Error(),
// 	// 				})
// 	// 			} else if validates, ok := err.Err.(validator.ValidationErrors); ok {
// 	// 				var vr []ValidationError

// 	// 				for _, validate := range validates {
// 	// 					vr = append(vr, ValidationError{
// 	// 						Field:   validate.Field(),
// 	// 						Message: validate.Error(),
// 	// 						Tag:     validate.Tag(),
// 	// 					})
// 	// 				}

// 	// 				c.JSON(http.StatusBadRequest, &FailResponse{
// 	// 					Code:    http.StatusBadRequest,
// 	// 					Status:  FailStatus,
// 	// 					Message: "Validate error, please see more on 'Errors' field",
// 	// 					Errors:  vr,
// 	// 				})
// 	// 			} else if _, ok := err.Err.(*BindError); ok {
// 	// 				c.JSON(http.StatusBadRequest, &FailResponse{
// 	// 					Code:    http.StatusBadRequest,
// 	// 					Status:  FailStatus,
// 	// 					Message: err.Err.Error(),
// 	// 				})
// 	// 			} else if e, ok := err.Err.(*apperror.CustomError); ok {
// 	// 				// for force error response cpa_odm payment (only charge api)
// 	// 				isPaymentChargeError := strings.Contains(e.ErrCode, "payment.counter_service.payment")
// 	// 				isPaymentVoidError := strings.Contains(e.ErrCode, "payment.counter_service.void")
// 	// 				if isPaymentChargeError {
// 	// 					code, message := extractPaymentErrorMessage(err.Err.Error())
// 	// 					c.JSON(e.Code, gin.H{
// 	// 						"status":  e.Code,
// 	// 						"message": message,
// 	// 						"name":    "PaymentChargeFailure",
// 	// 						"code":    code,
// 	// 						"errors":  []interface{}{},
// 	// 						"data":    nil,
// 	// 					})
// 	// 				} else if isPaymentVoidError {
// 	// 					code, message := extractPaymentErrorMessage(err.Err.Error())
// 	// 					c.JSON(e.Code, gin.H{
// 	// 						"status":  e.Code,
// 	// 						"message": message,
// 	// 						"name":    "PaymentVoidFailure",
// 	// 						"code":    code,
// 	// 						"errors":  []interface{}{},
// 	// 						"data":    nil,
// 	// 					})
// 	// 				} else {
// 	// 					c.JSON(e.Code, &FailResponse{
// 	// 						Code:    e.Code,
// 	// 						Status:  e.Status,
// 	// 						ErrCode: e.ErrCode,
// 	// 						Message: getMessage(c, err, languageBundles, e),
// 	// 						Errors:  e.Errors,
// 	// 					})
// 	// 				}
// 	// 			} else {
// 	// 				c.JSON(http.StatusInternalServerError, &FailResponse{
// 	// 					Code:    http.StatusInternalServerError,
// 	// 					Status:  FailStatus,
// 	// 					Message: err.Err.Error(),
// 	// 				})
// 	// 			}
// 	// 		}
// 	// 		c.Abort()
// 	// 	}
// 	// }
// 	return nil
// }

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
