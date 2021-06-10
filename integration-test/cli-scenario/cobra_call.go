package cliscenario

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/cli/spider/cmd"
	"github.com/stretchr/testify/assert"
)

func SpiderCmdTest(t *testing.T, tc TestCases) (string, error) {

	var (
		res string = ""
		err error  = nil
	)

	t.Run(tc.name, func(t *testing.T) {

		spiderCmd := cmd.NewRootCmd()
		b := bytes.NewBufferString("")
		spiderCmd.SetOut(b)
		spiderCmd.SetArgs(tc.cmdArgs)
		spiderCmd.Execute()

		out, err := ioutil.ReadAll(b)

		if assert.NoError(t, err) {
			if strings.HasPrefix(string(out), "{") {
				dst := new(bytes.Buffer)
				err = json.Compact(dst, out)
				if assert.NoError(t, err) {
					res = dst.String()
				}
			} else {
				res = string(out)
			}

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
