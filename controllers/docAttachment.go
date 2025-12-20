package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dronm/gobizapp/database"
	"github.com/dronm/gobizapp/services"

	"github.com/dronm/gobizapp/models"
)

func DelFile(c *gin.Context) {
	funcName := "DelFile"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	model := models.DocAttachment_delete_file{}
	if err := c.ShouldBind(&model); err != nil {
		ServeError(c, http.StatusBadRequest, err.Error(), err)
		return 
	}

	serv := services.NewDocAttachmentService(database.DB, sess)
	err := serv.DelFile(c.Request.Context(), model.Ref, model.ContentId) 
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, model)
}

func GetFile(c *gin.Context) {
	funcName := "GetFile"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	model := models.DocAttachment_get_file{}
	if err := c.ShouldBind(&model); err != nil {
		ServeError(c, http.StatusBadRequest, err.Error(), err)
		return 
	}
	serv := services.NewDocAttachmentService(database.DB, sess)
	fileName, attName, err := serv.GetFile(c.Request.Context(), model.Ref, model.ContentId) 
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	var disposition string
	if model.Inline {
		disposition = "inline"
	}else {
		disposition = "attachment"
	}
	c.Header("Content-Disposition", fmt.Sprintf("%s; filename=%s", disposition, attName))
	c.File(fileName)
}

func AddFile(c *gin.Context) {
	funcName := "AddFile"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	file, fileHeader, err := c.Request.FormFile("content_data") // this field contains file data
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	defer file.Close()

	ref := models.Ref{}
	if err := json.Unmarshal([]byte(c.PostForm("ref")), &ref); err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	contentInfo := models.DocAttachmentContentInfo{}
	if err := json.Unmarshal([]byte(c.PostForm("content_info")), &contentInfo); err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	contentInfo.Name = fileHeader.Filename
	att := models.DocAttachment{
		Ref:         ref,
		ContentInfo: contentInfo,
	}

	serv := services.NewDocAttachmentService(database.DB, sess)
	model, err := serv.AddFile(c.Request.Context(), file, att)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
		return
	}
	c.JSON(http.StatusOK, model)
}

