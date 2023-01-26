package http

import (
	"github.com/gin-gonic/gin"

	activityGroups "golang-todo/domain/activity-groups"
)

type HttpActivityGroupsHandler struct {
	activityGroupsUsecase activityGroups.Usecase
}

func NewActivityGroupsHttpHandler(activityGroups activityGroups.Usecase, httpRouter *gin.Engine) {
	handler := &HttpActivityGroupsHandler{
		activityGroupsUsecase: activityGroups,
	}

	public := httpRouter.Group("/public/api/v1")
	public.GET("/activity-groups", handler.GetAll)
	public.GET("/activity-groups/:id", handler.GetByID)
	public.POST("/activity-groups", handler.Store)
	public.PATCH("/activity-groups/:id", handler.Update)
	public.DELETE("/activity-groups/:id", handler.Delete)
}
