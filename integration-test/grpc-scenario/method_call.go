package grpcscenario

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

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

	if strings.HasPrefix(result, "{") {
		dst := new(bytes.Buffer)
		err = json.Compact(dst, []byte(result))
		if err != nil {
			return
		}
		result = dst.String()
	}

	if retVal[1].Interface() != nil {
		err = retVal[1].Interface().(error)
	}
	return
}

func MethodTest(t *testing.T, tc TestCases) (string, error) {

	var (
		res string = ""
		err error  = nil
	)

	t.Run(tc.Name, func(t *testing.T) {

		res, err = Call(tc.Instance, tc.Method, tc.Args)
		if assert.NoError(t, err) {
			if tc.ExpectResStartsWith != "" {
				if !assert.True(t, strings.HasPrefix(res, tc.ExpectResStartsWith)) {
					fmt.Fprintf(os.Stderr, "\n                Not Equal: \n"+
						"                  Expected Start With: %s\n"+
						"                  Actual  : %s\n", tc.ExpectResStartsWith, res)
				}
			} else {
				if !assert.Equal(t, "", res) {
					fmt.Fprintf(os.Stderr, "\n                Not Equal: \n"+
						"      Expected Start With: %s\n"+
						"      Actual  : %s\n", tc.ExpectResStartsWith, res)
				}
			}
		}
	})

	return res, err

}
