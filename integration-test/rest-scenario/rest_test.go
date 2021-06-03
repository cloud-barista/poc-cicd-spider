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
			whenURL:              "/spider/cloudos",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        "",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"cloudos":[`,
		}
		EchoGetTest(t, tc)

		tc = TestCases{
			name:                 "register cloud driver",
			echoFunc:             "RegisterCloudDriver",
			whenURL:              "/spider/driver",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		EchoPostTest(t, tc)

		tc = TestCases{
			name:                 "register credential",
			echoFunc:             "RegisterCredential",
			whenURL:              "/spider/credential",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_name00"}]}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		EchoPostTest(t, tc)

		tc = TestCases{
			name:                 "register region",
			echoFunc:             "RegisterRegion",
			whenURL:              "/spider/region",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		EchoPostTest(t, tc)

		tc = TestCases{
			name:                 "create connection config",
			echoFunc:             "CreateConnectionConfig",
			whenURL:              "/spider/connectionconfig",
			givenQueryParams:     "",
			givenParaNames:       nil,
			givenParaVals:        nil,
			givenPostData:        `{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`,
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		EchoPostTest(t, tc)

		tearDownForRest()
	})

	t.Run("sshrun test", func(t *testing.T) {
		setUpForRest()

		tc := TestCases{
			name:           "ssh run",
			echoFunc:       "SshRun",
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
		EchoPostTest(t, tc)

		tearDownForRest()
	})
}
