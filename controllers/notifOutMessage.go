package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"

	"github.com/dronm/gobizapp/models"
	"github.com/dronm/gobizapp/services"
)

func NotifOutMessageList(c *gin.Context) {
	funcName := "NotifOutMessageList"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	params, err := ParseCollectionParams(c)
	if err != nil {
		ServeError(c, http.StatusBadRequest, funcName+" json.Unmarshal()", err)
		return
	}
	serv := services.NewNotifOutMessageService(database.DB, sess)
	resList, tot, err := serv.FetchList(c.Request.Context(), params)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" FetchList()", err)
		return
	}
	c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})

}

func NotifOutMessageDetail(c *gin.Context) {
	funcName := "NotifOutMessageDetail"
	id := GetDetailID(c, funcName)
	if id == nil {
		return
	}
	sess :=GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewNotifOutMessageService(database.DB, sess)
	model, err := serv.FetchDetail(c.Request.Context(), *id)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, model)
}

