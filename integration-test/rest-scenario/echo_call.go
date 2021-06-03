package restscenario

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	restruntime "github.com/cloud-barista/poc-cicd-spider/api-runtime/rest-runtime"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var funcs = map[string]interface{}{
	"ListCloudOS":            restruntime.ListCloudOS,
	"RegisterCloudDriver":    restruntime.RegisterCloudDriver,
	"ListCloudDriver":        restruntime.ListCloudDriver,
	"GetCloudDriver":         restruntime.GetCloudDriver,
	"UnRegisterCloudDriver":  restruntime.UnRegisterCloudDriver,
	"RegisterCredential":     restruntime.RegisterCredential,
	"ListCredential":         restruntime.ListCredential,
	"GetCredential":          restruntime.GetCredential,
	"UnRegisterCredential":   restruntime.UnRegisterCredential,
	"RegisterRegion":         restruntime.RegisterRegion,
	"ListRegion":             restruntime.ListRegion,
	"GetRegion":              restruntime.GetRegion,
	"UnRegisterRegion":       restruntime.UnRegisterRegion,
	"CreateConnectionConfig": restruntime.CreateConnectionConfig,
	"ListConnectionConfig":   restruntime.ListConnectionConfig,
	"GetConnectionConfig":    restruntime.GetConnectionConfig,
	"DeleteConnectionConfig": restruntime.DeleteConnectionConfig,
	"CreateImage":            restruntime.CreateImage,
	"ListImage":              restruntime.ListImage,
	"GetImage":               restruntime.GetImage,
	"DeleteImage":            restruntime.DeleteImage,
	"ListVMSpec":             restruntime.ListVMSpec,
	"GetVMSpec":              restruntime.GetVMSpec,
	"ListOrgVMSpec":          restruntime.ListOrgVMSpec,
	"GetOrgVMSpec":           restruntime.GetOrgVMSpec,
	"CreateVPC":              restruntime.CreateVPC,
	"ListVPC":                restruntime.ListVPC,
	"GetVPC":                 restruntime.GetVPC,
	"DeleteVPC":              restruntime.DeleteVPC,
	"AddSubnet":              restruntime.AddSubnet,
	"RemoveSubnet":           restruntime.RemoveSubnet,
	"RemoveCSPSubnet":        restruntime.RemoveCSPSubnet,
	"ListAllVPC":             restruntime.ListAllVPC,
	"DeleteCSPVPC":           restruntime.DeleteCSPVPC,
	"CreateSecurity":         restruntime.CreateSecurity,
	"ListSecurity":           restruntime.ListSecurity,
	"GetSecurity":            restruntime.GetSecurity,
	"DeleteSecurity":         restruntime.DeleteSecurity,
	"ListAllSecurity":        restruntime.ListAllSecurity,
	"DeleteCSPSecurity":      restruntime.DeleteCSPSecurity,
	"CreateKey":              restruntime.CreateKey,
	"ListKey":                restruntime.ListKey,
	"GetKey":                 restruntime.GetKey,
	"DeleteKey":              restruntime.DeleteKey,
	"ListAllKey":             restruntime.ListAllKey,
	"DeleteCSPKey":           restruntime.DeleteCSPKey,
	"StartVM":                restruntime.StartVM,
	"ListVM":                 restruntime.ListVM,
	"GetVM":                  restruntime.GetVM,
	"TerminateVM":            restruntime.TerminateVM,
	"ListAllVM":              restruntime.ListAllVM,
	"TerminateCSPVM":         restruntime.TerminateCSPVM,
	"ListVMStatus":           restruntime.ListVMStatus,
	"GetVMStatus":            restruntime.GetVMStatus,
	"ControlVM":              restruntime.ControlVM,
	"SshRun":                 restruntime.SshRun,
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func EchoGetTest(t *testing.T, tc TestCases) (string, error) {

	var (
		body string = ""
		err  error  = nil
	)

	t.Run(tc.name, func(t *testing.T) {
		e := echo.New()
		var req *http.Request = nil
		if tc.givenPostData != "" {
			req = httptest.NewRequest(http.MethodGet, "/"+tc.givenQueryParams, bytes.NewBuffer([]byte(tc.givenPostData)))
		} else {
			req = httptest.NewRequest(http.MethodGet, "/"+tc.givenQueryParams, nil)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(tc.whenURL)
		if tc.givenParaNames != nil {
			c.SetParamNames(tc.givenParaNames...)
			c.SetParamValues(tc.givenParaVals...)
		}

		_, err = Call(funcs, tc.echoFunc, c)
		if assert.NoError(t, err) {
			assert.Equal(t, tc.expectStatus, rec.Code)
			body = rec.Body.String()
			fmt.Printf("===== body : %s\n", body)
			if tc.expectBodyStartsWith != "" {
				assert.True(t, strings.HasPrefix(body, tc.expectBodyStartsWith))
			} else {
				assert.Equal(t, "", body)
			}
		}
	})

	return body, err
}

func EchoPostTest(t *testing.T, tc TestCases) (string, error) {

	var (
		body string = ""
		err  error  = nil
	)

	t.Run(tc.name, func(t *testing.T) {
		e := echo.New()
		var req *http.Request = nil
		if tc.givenPostData != "" {
			req = httptest.NewRequest(http.MethodPost, "/"+tc.givenQueryParams, bytes.NewBuffer([]byte(tc.givenPostData)))
		} else {
			req = httptest.NewRequest(http.MethodPost, "/"+tc.givenQueryParams, nil)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(tc.whenURL)
		if tc.givenParaNames != nil {
			c.SetParamNames(tc.givenParaNames...)
			c.SetParamValues(tc.givenParaVals...)
		}

		_, err = Call(funcs, tc.echoFunc, c)
		if assert.NoError(t, err) {
			assert.Equal(t, tc.expectStatus, rec.Code)
			body = rec.Body.String()
			fmt.Printf("===== body : %s\n", body)
			if tc.expectBodyStartsWith != "" {
				assert.True(t, strings.HasPrefix(body, tc.expectBodyStartsWith))
			} else {
				assert.Equal(t, "", body)
			}
		}
	})

	return body, err
}

func EchoDeleteTest(t *testing.T, tc TestCases) (string, error) {

	var (
		body string = ""
		err  error  = nil
	)

	t.Run(tc.name, func(t *testing.T) {
		e := echo.New()
		var req *http.Request = nil
		if tc.givenPostData != "" {
			req = httptest.NewRequest(http.MethodDelete, "/"+tc.givenQueryParams, bytes.NewBuffer([]byte(tc.givenPostData)))
		} else {
			req = httptest.NewRequest(http.MethodDelete, "/"+tc.givenQueryParams, nil)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(tc.whenURL)
		if tc.givenParaNames != nil {
			c.SetParamNames(tc.givenParaNames...)
			c.SetParamValues(tc.givenParaVals...)
		}

		_, err = Call(funcs, tc.echoFunc, c)
		if assert.NoError(t, err) {
			assert.Equal(t, tc.expectStatus, rec.Code)
			body = rec.Body.String()
			fmt.Printf("===== body : %s\n", body)
			if tc.expectBodyStartsWith != "" {
				assert.True(t, strings.HasPrefix(body, tc.expectBodyStartsWith))
			} else {
				assert.Equal(t, "", body)
			}
		}
	})

	return body, err
}