package restscenario

import (
	"net/http"
	"testing"
)

func TestRestFull(t *testing.T) {
	t.Run("rest api full test for mock driver", func(t *testing.T) {
		SetUpForRest()

		tc := TestCases{
			Name:                 "list cloud os",
			EchoFunc:             "ListCloudOS",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/cloudos",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"cloudos":[`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "register cloud driver",
			EchoFunc:             "RegisterCloudDriver",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/driver",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "register credential",
			EchoFunc:             "RegisterCredential",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/credential",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_full"}]}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "register region",
			EchoFunc:             "RegisterRegion",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/region",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "create connection config",
			EchoFunc:             "CreateConnectionConfig",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/connectionconfig",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

	t.Run("sshrun test", func(t *testing.T) {
		SetUpForRest()

		tc := TestCases{
			Name:           "ssh run",
			EchoFunc:       "SshRun",
			HttpMethod:     http.MethodPost,
			WhenURL:        "/spider/sshrun",
			GivenParaNames: nil,
			GivenParaVals:  nil,
			GivenPostData: `{
				"PrivateKey": [],
				"UserName": "root",
				"ServerPort": "node12:22",
				"Command": "hostname"
			}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `"hostname success"`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})
}
