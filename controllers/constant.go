package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dronm/gobizapp/errs"
	"github.com/dronm/gobizapp/models"
	"github.com/dronm/gobizapp/services"

	"github.com/gin-gonic/gin"
)

const errConstantUndefined = "constant is undefined"

func ConstantSet(c *gin.Context) {
	funcName := "ConstantSet"

	constID := c.Param("id")
	if constID == "" {
		ServeError(c, http.StatusInternalServerError, funcName, fmt.Errorf("constant ID is undefined"))
		return
	}

	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	modelVal := models.ConstantSet{}
	if !CheckBindModel(c, funcName, &modelVal) {
		return
	}

	if err := services.ConstantServiceInstance.Update(c.Request.Context(), constID, modelVal.Val.GetValue()); err != nil {
		ServeError(c, http.StatusBadRequest, funcName, err)
		return
	}
}

// ConstantGet retrieves a list of constants defined by parameter 'ids'.
// ids is an array of strings.
func ConstantGet(c *gin.Context) {
	funcName := "ConstantGet"

	constIDStr := c.Query("ids")
	if constIDStr == "" {
		ServeError(c, http.StatusBadRequest, funcName,
			errs.NewPublicError("VALIDATION_FAILED"),
		)
		return
	}
	constDecodedID, err := url.QueryUnescape(constIDStr)
	if err != nil {
		ServeError(c, http.StatusBadRequest, 
			fmt.Sprintf("%s url.QueryUnescape(): %v",funcName, err),
			errs.NewPublicError("VALIDATION_FAILED"),
		)
		return
	}
	var constIDList []string
	if err := json.Unmarshal([]byte(constDecodedID), &constIDList); err != nil {
		ServeError(c, http.StatusBadRequest,
			fmt.Sprintf("%s json.Unmarshal(): %v",funcName, err),
			errs.NewPublicError("VALIDATION_FAILED"),
		)
		return
	}

	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	// userLogin := models.UserLogin{}
	// if err := sess.Get("user", &userLogin); err != nil {
	// 	ServeError(c, http.StatusInternalServerError,
	// 		fmt.Sprintf("%s sess.Get(): %v",funcName, err),
	// 		err,
	// 	)
	// 	return
	// }
	// if userLogin.RoleID == "" {
	// 	ServeError(c, http.StatusBadRequest,
	// 		fmt.Sprintf("%s RoleId is not set",funcName),
	// 		errs.NewPublicError("NOT_LOGGED"),
	// 	)
	// 	return
	// }

	models, err := services.ConstantServiceInstance.FetchValues(c.Request.Context(), constIDList, UserRoleID)
	if err != nil {
		ServeError(c, http.StatusBadRequest, funcName, err)
		return
	}

	c.JSON(http.StatusOK, models)
}

func ConstantList(c *gin.Context) {
	funcName := "ConstantList"
	sess := GetSession(c, funcName)
	if sess == nil {
		return
	}

	// params, err := ParseCollectionParams(c)
	// if err != nil {
	// 	ServeError(c, http.StatusBadRequest, funcName+" json.Unmarshal()", err)
	// 	return
	// }
	// serv := services.NewConstantService(database.DB, sess)
	// resList, tot, err := serv.FetchList(c.Request.Context(), params)
	// if err != nil {
	// 	ServeError(c, http.StatusInternalServerError, funcName+" FetchList()", err)
	// 	return
	// }
	// c.JSON(http.StatusOK, &models.Collection{Data: resList, Agg: tot})

}


