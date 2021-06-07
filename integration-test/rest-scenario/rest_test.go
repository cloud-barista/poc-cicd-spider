package restscenario

import (
	"net/http"
	"testing"
)

func TestRestApi(t *testing.T) {
	t.Run("rest api full test for mock driver", func(t *testing.T) {
		setUpForRest()

		tc := TestCases{
			name:                 "list cloud os",
			echoFunc:             "ListCloudOS",
			httpMethod:           http.MethodGet,
			whenURL:              "/spider/cloudos",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        "",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"cloudos":[`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			name:                 "register cloud driver",
			echoFunc:             "RegisterCloudDriver",
			httpMethod:           http.MethodPost,
			whenURL:              "/spider/driver",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			name:                 "register credential",
			echoFunc:             "RegisterCredential",
			httpMethod:           http.MethodPost,
			whenURL:              "/spider/credential",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_name00"}]}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			name:                 "register region",
			echoFunc:             "RegisterRegion",
			httpMethod:           http.MethodPost,
			whenURL:              "/spider/region",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			name:                 "create connection config",
			echoFunc:             "CreateConnectionConfig",
			httpMethod:           http.MethodPost,
			whenURL:              "/spider/connectionconfig",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		EchoTest(t, tc)

		tearDownForRest()
	})

	t.Run("sshrun test", func(t *testing.T) {
		setUpForRest()

		tc := TestCases{
			name:           "ssh run",
			echoFunc:       "SshRun",
			httpMethod:     http.MethodPost,
			whenURL:        "/spider/sshrun",
			givenParaNames: nil,
			givenParaVals:  nil,
			givenPostData: `{
				"PrivateKey": [],
				"UserName": "root",
				"ServerPort": "node12:22",
				"Command": "hostname"
			}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `"hostname success"`,
		}
		EchoTest(t, tc)

		tearDownForRest()
	})
}
