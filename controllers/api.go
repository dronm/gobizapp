package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/dronm/gobizapp/api"
	"github.com/dronm/gobizapp/database"
	"github.com/gin-gonic/gin"
)

// APIGet executes GET request. Parameters are passed in their original order
// of appearence in the URL.
func APIGet(c *gin.Context) {
	funcName := "APIGet"

	// params
	rawQuery := c.Request.URL.RawQuery
	pairs := strings.Split(rawQuery, "&")
	var params []string

	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		kv := strings.SplitN(pair, "=", 2)
		val := ""
		// key, _ = url.QueryUnescape(kv[0])
		if len(kv) > 1 {
			val, _ = url.QueryUnescape(kv[1])
		}
		params = append(params, val)
	}

	service, method, err := extractService(c)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
	}

	sess := GetSession(c, funcName)
	if sess == nil {
		ServeError(c, http.StatusMethodNotAllowed, funcName+" session not found", err)
	}

	httpRes, jsonRes, err := CallServiceMethod(c.Request.Context(), service, method , params, &api.ServiceContext{DB: database.DB, Session: sess})
	if err != nil {
		ServeError(c, httpRes, funcName, err)
	}

	c.JSON(http.StatusOK, jsonRes)
}

// APIPost handle post requests. It only deals with app/json requests.
func APIPost(c *gin.Context) {
	funcName := "APIPost"

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName+" io.ReadAll()", err)
		return
	}

	params, err := api.UnmarshalParams(bodyBytes)
	if err != nil {
		ServeError(c, http.StatusBadRequest, funcName+" "+err.Error(), err)
		return
	}

	service, method, err := extractService(c)
	if err != nil {
		ServeError(c, http.StatusInternalServerError, funcName, err)
	}

	sess := GetSession(c, funcName)
	if sess == nil {
		ServeError(c, http.StatusMethodNotAllowed, funcName+" session not found", err)
	}

	httpRes, jsonRes, err := CallServiceMethod(c.Request.Context(), service, method , params, &api.ServiceContext{DB: database.DB, Session: sess})
	if err != nil {
		ServeError(c, httpRes, funcName, err)
	}

	c.JSON(http.StatusOK, jsonRes)
}

// CallServiceMethod dynamically calls a service method with the given params. funcName is a name of the calling function
// for the log.
// It returns an http result code, json result body and error.
func CallServiceMethod(ctx context.Context, service, method string, params []string, src *api.ServiceContext) (int, any, error) {
	results, err := api.CallMethod(
		ctx,
		service,
		method,
		params,
		src,
	)
	if err != nil {
		return http.StatusBadRequest, nil, fmt.Errorf("%s.%s api.CallMethod(): %v", service, method, err)
	}

	// last result is always an error
	var resultBody any
	if len(results) > 0 {
		last := results[len(results)-1]
		errVal := results[len(results)-1]
		if last.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) && !last.IsNil() {
			err := errVal.Interface().(error)
			return http.StatusInternalServerError, nil, err
		}
		if len(results) > 1 {
			if len(results) == 1 {
				// one model
				resultBody = results[0].Interface()
			} else {
				// slice of models, minus error result
				res := make([]any, len(results)-1)
				for i := 0; i < len(results)-1; i++ {
					res[i] = results[i].Interface()
				}
				resultBody = res
			}
		}
	}

	return http.StatusOK, resultBody, nil
}

// extractService is a helper functin to retrieve
// service, method from http request.
func extractService(c *gin.Context) (service, method string, err error) {
	service = PascalCaseFromKebabCase(c.Param("service")) // kebab-cased service
	if service == "" {
		err = fmt.Errorf("service is undefined")
		return
	}

	method = PascalCaseFromKebabCase(c.Param("method")) // kebab-cased method
	if method == "" {
		err = fmt.Errorf("service method is undefined")
		return
	}

	return
}

func PascalCaseFromKebabCase(s string) string {
	if s == "" {
		return ""
	}
	var res strings.Builder
	for w := range strings.SplitSeq(s, "-") {
		if len(w) > 0 {
			res.WriteString(strings.ToUpper(w[0:1]))
			if len(w) > 1 {
				res.WriteString(w[1:])
			}
		}
	}
	return res.String()
}
