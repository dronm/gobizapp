package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"

	"github.com/dronm/gobizapp/models"
	"github.com/dronm/gobizapp/services"
)

func NotifTemplateList(c *gin.Context) {
	funcName := "NotifTemplateList"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	params, err := ParseCollectionParams(c)
	if err != nil {
		ServeError(c, http.StatusBadRequest, funcName+" json.Unmarshal()", err)
		return
	}
	serv := services.NewNotifTemplateService(database.DB, sess)
	resList, tot, err := serv.FetchList(c.Request.Context(), params)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" FetchList()", err)
		return
	}
	c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})

}

func NotifTemplateDetail(c *gin.Context) {
	funcName := "NotifTemplateDetail"
	id := GetDetailID(c, funcName)
	if id == nil {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifTemplateService(database.DB, sess)
	model, err := serv.FetchDetail(c.Request.Context(), *id)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, model)
}

func NotifTemplateDelete(c *gin.Context) {
	funcName := "NotifTemplateDelete"

	modelKeys := []models.NotifTemplateKey{}
	if !CheckBindModel(c, funcName, &modelKeys) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifTemplateService(database.DB, sess)
	rows, err := serv.Delete(c.Request.Context(), modelKeys)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" serv.Delete()", err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

func NotifTemplateUpdate(c *gin.Context) {
	funcName := "NotifTemplateUpdate"

	modelUpdate := models.NotifTemplateUpdate{}
	if !CheckBindModel(c, funcName, &modelUpdate) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifTemplateService(database.DB, sess)
	rows, err := serv.Update(c.Request.Context(), modelUpdate.Keys, modelUpdate.Model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

func NotifTemplateInsert(c *gin.Context) {
	funcName := "NotifTemplateInsert"

	model := models.NotifTemplate{}
	if !CheckBindModel(c, funcName, &model) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifTemplateService(database.DB, sess)
	fields, err := serv.Insert(c.Request.Context(), model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" serv.Insert()", err)
		return
	}
	c.JSON(http.StatusOK, fields)
}

