package grpcscenario

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"
	"github.com/stretchr/testify/assert"
)

func Call(any interface{}, name string, params []interface{}) (result string, err error) {
	f := reflect.ValueOf(any).MethodByName(name)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	retVal := f.Call(in)
	result = retVal[0].String()

	dst := new(bytes.Buffer)
	err = json.Compact(dst, []byte(result))
	if err != nil {
		return
	}
	result = dst.String()

	if retVal[1].Interface() != nil {
		err = retVal[1].Interface().(error)
	}
	return
}

func CIMMethodTest(t *testing.T, cim *api.CIMApi, tc TestCases) (string, error) {

	var (
		res string = ""
		err error  = nil
	)

	t.Run(tc.name, func(t *testing.T) {

		res, err = Call(cim, tc.method, tc.args)
		if assert.NoError(t, err) {
			fmt.Printf("===== result : %s\n", res)
			if tc.expectResStartsWith != "" {
				assert.True(t, strings.HasPrefix(res, tc.expectResStartsWith))
			} else {
				assert.Equal(t, "", res)
			}
		}
	})

	return res, err

}

func CCMMethodTest(t *testing.T, ccm *api.CCMApi, tc TestCases) (string, error) {

	var (
		res string = ""
		err error  = nil
	)

	t.Run(tc.name, func(t *testing.T) {

		res, err = Call(ccm, tc.method, tc.args)
		if assert.NoError(t, err) {
			fmt.Printf("=====res : %s\n", res)
			if tc.expectResStartsWith != "" {
				assert.True(t, strings.HasPrefix(res, tc.expectResStartsWith))
			} else {
				assert.Equal(t, "", res)
			}
		}
	})

	return res, err
}
