package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"
	"github.com/dronm/gobizapp/models"
	"github.com/dronm/gobizapp/services"
)

func NotifAppList(c *gin.Context) {
	funcName := "NotifAppList"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	params, err := ParseCollectionParams(c)
	if err != nil {
		ServeError(c, http.StatusBadRequest, funcName+" json.Unmarshal()", err)
		return
	}
	serv := services.NewNotifAppService(database.DB, sess)
	resList, tot, err := serv.FetchList(c.Request.Context(), params)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" FetchList()", err)
		return
	}
	c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})

}

func NotifAppDetail(c *gin.Context) {
	funcName := "NotifAppDetail"

	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifAppService(database.DB, sess)
	model, err := serv.FetchDetail(c.Request.Context())
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, model)
}

func NotifAppUpdate(c *gin.Context) {
	funcName := "NotifAppUpdate"

	modelUpdate := models.NotifAppUpdate{}
	if CheckBindModel(c, funcName, &modelUpdate) {
		return
	}
	sess :=GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifAppService(database.DB, sess)
	rows, err := serv.Update(c.Request.Context(), modelUpdate.Model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

