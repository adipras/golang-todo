package message

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const ERROR_STATUS = "Layanan sedang mengalami gangguan"
const SUCCESS_UPDATE_STATUS = "Berhasil melakukan pembaharuan data"

func ReturnOk(ctx *gin.Context, data interface{}, status string, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ReturnNotFound(ctx *gin.Context, status string, message string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"status":  status,
		"message": message,
	})
}

func ReturnBadRequest(ctx *gin.Context, message string, status string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"status":  status,
		"message": message,
	})
}

func ReturnInternalServerError(ctx *gin.Context, msg ...string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message":               ERROR_STATUS,
		"message for developer": msg,
	})
}

func ReturnSuccessInsert(ctx *gin.Context, data interface{}, status string, message string) {
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ReturnSuccessDelete(ctx *gin.Context, data interface{}, status string, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ReturnSuccessUpdateNew(ctx *gin.Context, data interface{}, status string, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ReturnSuccessUpdate(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": SUCCESS_UPDATE_STATUS,
	})
}
