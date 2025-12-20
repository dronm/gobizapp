package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"
	"github.com/dronm/gobizapp/models"
	
	"github.com/dronm/gobizapp/services"
)

// ClientComponentInsert is a public method for inserting row.
func ClientComponentInsert(c *gin.Context) {
	model := models.ClientComponent{}
	if !CheckBindModel(c, "ClientComponentInsert", &model) {
		return
	}
	sess := GetSession(c, "ClientComponentInsert")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentService(database.DB, sess)
	fields, err := serv.Insert(c.Request.Context(), model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentInsert serv.Insert()", err)
		return
	}
	c.JSON(http.StatusOK, fields)
}

// ClientComponentUpdate is a public method to update models.ClientComponent by key fields.
func ClientComponentUpdate(c *gin.Context) {
	modelUpdate := models.ClientComponentUpdate{}
	if !CheckBindModel(c, "ClientComponentUpdate", &modelUpdate) {
		return
	}
	sess := GetSession(c, "ClientComponentUpdate")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentService(database.DB, sess)
	rows, err := serv.Update(c.Request.Context(), modelUpdate.Keys, modelUpdate.Model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentUpdate", err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

// ClientComponentDelete is a public mehtod for deleting row by key fields.
func ClientComponentDelete(c *gin.Context) {
	modelKeys := []models.ClientComponentKey{}
	if !CheckBindModel(c, "ClientComponentDelete", &modelKeys) {
		return
	}
	sess := GetSession(c, "ClientComponentDelete")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentService(database.DB, sess)
	rows, err := serv.Delete(c.Request.Context(), modelKeys)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentDelete serv.Delete()", err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

// ClientComponentDetail is a public method to retrieve one row from models.ClientComponent
func ClientComponentDetail(c *gin.Context) {
	id := GetDetailID(c, "ClientComponentDetail")
	if id == nil {
		return
	}
	sess := GetSession(c, "ClientComponentDetail")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentService(database.DB, sess)
	model, err := serv.FetchDetail(c.Request.Context(), *id)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentDetail", err)
		return
	}
	c.JSON(http.StatusOK, model)
}

// ClientComponentList is a public method for retrieving a list model with filters.
func ClientComponentList(c *gin.Context) {
	sess := GetSession(c, "ClientComponentList")
	if sess == nil {
		return
	}

	params, err := ParseCollectionParams(c)
	if err != nil {
		ServeError(c, http.StatusBadRequest, "ClientComponentList json.Unmarshal()", err)
		return
	}
	serv := services.NewClientComponentService(database.DB, sess)
	resList, tot, err := serv.FetchList(c.Request.Context(), params)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentList FetchList()", err)
		return
	}
	c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})
}
