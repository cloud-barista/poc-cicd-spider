package restscenario

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	restruntime "github.com/cloud-barista/poc-cicd-spider/api-runtime/rest-runtime"
	aw "github.com/cloud-barista/poc-cicd-spider/api-runtime/rest-runtime/admin-web"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var funcs = map[string]interface{}{
	"ListCloudOS":            restruntime.ListCloudOS,
	"EndpointInfo":           restruntime.EndpointInfo,
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
	"SSHRun":                 restruntime.SSHRun,
	"Frame":                  aw.Frame,
	"Top":                    aw.Top,
	"Driver":                 aw.Driver,
	"Credential":             aw.Credential,
	"Region":                 aw.Region,
	"Connectionconfig":       aw.Connectionconfig,
	"SpiderInfo":             aw.SpiderInfo,
	"VPC":                    aw.VPC,
	"VPCMgmt":                aw.VPCMgmt,
	"SecurityGroup":          aw.SecurityGroup,
	"SecurityGroupMgmt":      aw.SecurityGroupMgmt,
	"KeyPair":                aw.KeyPair,
	"KeyPairMgmt":            aw.KeyPairMgmt,
	"VM":                     aw.VM,
	"VMMgmt":                 aw.VMMgmt,
	"VMImage":                aw.VMImage,
	"VMSpec":                 aw.VMSpec,
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

func EchoTest(t *testing.T, tc TestCases) (string, error) {

	var (
		body string = ""
		err  error  = nil
	)

	t.Run(tc.Name, func(t *testing.T) {
		e := echo.New()
		var req *http.Request = nil
		if tc.GivenPostData != "" {
			req = httptest.NewRequest(tc.HttpMethod, "/"+tc.GivenQueryParams, bytes.NewBuffer([]byte(tc.GivenPostData)))
		} else {
			req = httptest.NewRequest(tc.HttpMethod, "/"+tc.GivenQueryParams, nil)
		}
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(tc.WhenURL)
		if tc.GivenParaNames != nil {
			c.SetParamNames(tc.GivenParaNames...)
			c.SetParamValues(tc.GivenParaVals...)
		}

		res, err := Call(funcs, tc.EchoFunc, c)
		if assert.NoError(t, err) {
			if res != nil && !res[0].IsNil() {
				he, ok := res[0].Interface().(*echo.HTTPError)
				if ok { // echo.NewHTTPError() 로 에러를 리턴했을 경우
					assert.Equal(t, tc.ExpectStatus, he.Code)
					body = fmt.Sprintf("%v", he.Message)
				} else { // err 로 에러를 리턴했을 경우
					body = fmt.Sprintf("%v", res[0])
				}
				if tc.ExpectBodyStartsWith != "" {
					if !assert.True(t, strings.HasPrefix(body, tc.ExpectBodyStartsWith)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.NewHTTPError): \n"+
							"                  Expected Start With: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
				if tc.ExpectBodyContains != "" {
					if !assert.True(t, strings.Contains(body, tc.ExpectBodyContains)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.NewHTTPError): \n"+
							"                  Expected Contains: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyContains, body)
					}
				}
				if tc.ExpectBodyStartsWith == "" && tc.ExpectBodyContains == "" {
					if !assert.True(t, "" == body) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.NewHTTPError): \n"+
							"      Expected StartWith/Contains: %s\n"+
							"      Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
			} else {
				assert.Equal(t, tc.ExpectStatus, rec.Code)
				body = rec.Body.String()
				if tc.ExpectBodyStartsWith != "" {
					if !assert.True(t, strings.HasPrefix(body, tc.ExpectBodyStartsWith)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.Context): \n"+
							"                  Expected Start With: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}
				if tc.ExpectBodyContains != "" {
					if !assert.True(t, strings.Contains(body, tc.ExpectBodyContains)) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.Context): \n"+
							"                  Expected Contains: %s\n"+
							"                  Actual  : %s\n", tc.ExpectBodyContains, body)
					}
				}
				if tc.ExpectBodyStartsWith == "" && tc.ExpectBodyContains == "" {
					if !assert.True(t, "" == body) {
						fmt.Fprintf(os.Stderr, "\n                Not Equal(echo.Context): \n"+
							"      Expected StartWith/Contains: %s\n"+
							"      Actual  : %s\n", tc.ExpectBodyStartsWith, body)
					}
				}

			}
		}
	})

	return body, err
}
