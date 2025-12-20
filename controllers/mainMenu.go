package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"

	"github.com/dronm/gobizapp/models"
	"github.com/dronm/gobizapp/services"
)

func MainMenuList(c *gin.Context) {
	funcName := "MainMenuList"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	params, err := ParseCollectionParams(c)
	if err != nil {
		ServeError(c, http.StatusBadRequest, funcName+" json.Unmarshal()", err)
		return
	}
	serv := services.NewMainMenuService(database.DB, sess)
	resList, tot, err := serv.FetchList(c.Request.Context(), params)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" FetchList()", err)
		return
	}
	c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})

}

func MainMenuDetail(c *gin.Context) {
	funcName := "MainMenuDetail"
	id := GetDetailID(c, funcName)
	if id == nil {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	model, err := serv.FetchDetail(c.Request.Context(), *id)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, model)
}

func MainMenuDelete(c *gin.Context) {
	funcName := "MainMenuDelete"

	modelKeys := []models.MainMenuKey{}
	if !CheckBindModel(c, funcName, &modelKeys) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	rows, err := serv.Delete(c.Request.Context(), modelKeys)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" serv.Delete()", err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

func MainMenuUpdate(c *gin.Context) {
	funcName := "MainMenuUpdate"

	modelUpdate := models.MainMenuUpdate{}
	if !CheckBindModel(c, funcName, &modelUpdate) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	rows, err := serv.Update(c.Request.Context(), modelUpdate.Keys, modelUpdate.Model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

func MainMenuInsert(c *gin.Context) {
	funcName := "MainMenuInsert"

	model := models.MainMenu{}
	if !CheckBindModel(c, funcName, &model) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	fields, err := serv.Insert(c.Request.Context(), model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" serv.Insert()", err)
		return
	}
	c.JSON(http.StatusOK, fields)
}

func MainMenuMoveUp(c *gin.Context) {
	funcName := "MainMenuMoveUp"

	modelKey := models.MainMenuKey{}
	if !CheckBindModel(c, funcName, &modelKey) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	err := serv.MoveUp(c.Request.Context(), int(modelKey.ID.GetValue()))
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func MainMenuMoveDown(c *gin.Context) {
	funcName := "MainMenuMoveDown"

	modelKey := models.MainMenuKey{}
	if !CheckBindModel(c, funcName, &modelKey) {
		return
	}
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	err := serv.MoveDown(c.Request.Context(), int(modelKey.ID.GetValue()))
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

/*
func MenuForUser(c *gin.Context) {
	funcName := "MenuForUser"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	serv := services.NewMainMenuService(database.DB, sess)
	menu , err := serv.FetchForUser(c.Request.Context())
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" FetchForUser()", err)
		return
	}
	c.JSON(http.StatusOK, menu)

}
*/

