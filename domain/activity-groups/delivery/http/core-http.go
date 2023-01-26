package http

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"golang-todo/utils/config"
	"golang-todo/utils/message"
	"golang-todo/valueobject"
)

func (handler *HttpActivityGroupsHandler) GetAll(ctx *gin.Context) {

	var param = map[string]interface{}{
		"ORDER": map[string]interface{}{
			"ORDER_BY": []string{"created_at ASC"},
		},
	}

	response, err := handler.activityGroupsUsecase.GetAll(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnOk(ctx, make(map[string]interface{}), "Success", err.Error())
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnOk(ctx, response, "Success", "Success")
}

func (handler *HttpActivityGroupsHandler) GetByID(ctx *gin.Context) {
	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"activity_id": ctx.Param("id"),
		},
	}

	response, err := handler.activityGroupsUsecase.GetOne(param)

	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnNotFound(ctx, "Not Found", "Activity with ID "+ctx.Param("id")+" Not Found")
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnOk(ctx, response, "Success", "Success")
}

func (handler *HttpActivityGroupsHandler) Store(ctx *gin.Context) {
	var payload valueobject.ActivityGroups
	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		if payload.Title == "" {
			message.ReturnBadRequest(ctx, "title cannot be null", "Bad Request")
			return
		}
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}

	feedback, err := handler.activityGroupsUsecase.Store(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessInsert(ctx, feedback, "Success", "Success")
}

func (handler *HttpActivityGroupsHandler) Update(ctx *gin.Context) {
	var payload valueobject.ActivityGroups
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		if payload.Title == "" {
			message.ReturnBadRequest(ctx, "title cannot be null", "Bad Request")
			return
		}
		message.ReturnBadRequest(ctx, err.Error(), config.ERROR_BIND_JSON)
		return
	}
	payload.ID = uint64(id)

	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"activity_id": payload.ID,
		},
	}
	data, err := handler.activityGroupsUsecase.GetOne(param)
	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnNotFound(ctx, "Not Found", "Activity with ID "+ctx.Param("id")+" Not Found")
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	if payload.Email == "" {
		payload.Email = data.Email
	}

	err = handler.activityGroupsUsecase.Update(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}
	data, _ = handler.activityGroupsUsecase.GetOne(param)

	message.ReturnSuccessUpdateNew(ctx, data, "Success", "Success")
}

func (handler *HttpActivityGroupsHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var param = map[string]interface{}{
		"AND": map[string]interface{}{
			"activity_id": id,
		},
	}
	payload, err := handler.activityGroupsUsecase.GetOne(param)
	if err != nil {
		if err.Error() == config.SQL_NOT_FOUND {
			message.ReturnNotFound(ctx, "Not Found", "Activity with ID "+ctx.Param("id")+" Not Found")
			return
		}
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	err = handler.activityGroupsUsecase.Delete(payload)

	if err != nil {
		message.ReturnInternalServerError(ctx, err.Error())
		log.Println(err.Error())
		return
	}

	message.ReturnSuccessDelete(ctx, map[string]interface{}{}, "Success", "Success")
}
