package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"

	"github.com/dronm/gobizapp/models"
	"github.com/dronm/gobizapp/services"
)

// ClientComponentSectionInsert is a public method for inserting row.
func ClientComponentSectionInsert(c *gin.Context) {
	model := models.ClientComponentSection{}
	if !CheckBindModel(c, "ClientComponentSectionInsert", &model) {
		return
	}
	sess := GetSession(c, "ClientComponentSectionInsert")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentSectionService(database.DB, sess)
	fields, err := serv.Insert(c.Request.Context(), model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentSectionInsert serv.Insert()", err)
		return
	}
	c.JSON(http.StatusOK, fields)
}

// ClientComponentSectionUpdate is a public method to update models.ClientComponentSection by key fields.
func ClientComponentSectionUpdate(c *gin.Context) {
	modelUpdate := models.ClientComponentSectionUpdate{}
	if !CheckBindModel(c, "ClientComponentSectionUpdate", &modelUpdate) {
		return
	}
	sess := GetSession(c, "ClientComponentSectionUpdate")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentSectionService(database.DB, sess)
	rows, err := serv.Update(c.Request.Context(), modelUpdate.Keys, modelUpdate.Model)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentSectionUpdate", err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

// ClientComponentSectionDelete is a public mehtod for deleting row by key fields.
func ClientComponentSectionDelete(c *gin.Context) {
	modelKeys := []models.ClientComponentSectionKey{}
	if !CheckBindModel(c, "ClientComponentSectionDelete", &modelKeys) {
		return
	}
	sess := GetSession(c, "ClientComponentSectionDelete")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentSectionService(database.DB, sess)
	rows, err := serv.Delete(c.Request.Context(), modelKeys)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentSectionDelete serv.Delete()", err)
		return
	}
	c.JSON(http.StatusOK, models.CollectionAlterResult{AffectedRows: rows})
}

// ClientComponentSectionDetail is a public method to retrieve one row from models.ClientComponentSection
func ClientComponentSectionDetail(c *gin.Context) {
	id := GetDetailID(c, "ClientComponentSectionDetail")
	if id == nil {
		return
	}
	sess := GetSession(c, "ClientComponentSectionDetail")
	if sess == nil {
		return
	}

	serv := services.NewClientComponentSectionService(database.DB, sess)
	model, err := serv.FetchDetail(c.Request.Context(), *id)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentSectionDetail", err)
		return
	}
	c.JSON(http.StatusOK, model)
}

// ClientComponentSectionList is a public method for retrieving a list model with filters.
func ClientComponentSectionList(c *gin.Context) {
	sess := GetSession(c, "ClientComponentSectionList")
	if sess == nil {
		return
	}

	params, err := ParseCollectionParams(c)
	if err != nil {
		ServeError(c, http.StatusBadRequest, "ClientComponentSectionList json.Unmarshal()", err)
		return
	}
	serv := services.NewClientComponentSectionService(database.DB, sess)
	resList, tot, err := serv.FetchList(c.Request.Context(), params)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, "ClientComponentSectionList FetchList()", err)
		return
	}
	c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})
}
