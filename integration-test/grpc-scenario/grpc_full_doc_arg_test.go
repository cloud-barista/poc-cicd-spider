package grpcscenario

import (
	"testing"
)

func TestGrpcFullDocArg(t *testing.T) {
	t.Run("grpc api full test for mock driver by doccument args style", func(t *testing.T) {
		SetUpForGrpc()

		tc := TestCases{
			Name:                "list cloud os",
			Instance:            CimApi,
			Method:              "ListCloudOS",
			Args:                nil,
			ExpectResStartsWith: `{"cloudos":[`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register cloud driver",
			Instance: CimApi,
			Method:   "CreateCloudDriver",
			Args: []interface{}{
				`{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`,
			},
			ExpectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register credential",
			Instance: CimApi,
			Method:   "CreateCredential",
			Args: []interface{}{
				`{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_name00"}]}`,
			},
			ExpectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register region",
			Instance: CimApi,
			Method:   "CreateRegion",
			Args: []interface{}{
				`{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`,
			},
			ExpectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create connection config",
			Instance: CimApi,
			Method:   "CreateConnectionConfig",
			Args: []interface{}{
				`{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`,
			},
			ExpectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "ssh run",
			Instance: CcmApi,
			Method:   "SSHRun",
			Args: []interface{}{
				`{
					"PrivateKey": [],
					"UserName": "root",
					"ServerPort": "node12:22",
					"Command": "hostname"
				}`,
			},
			ExpectResStartsWith: `{"Result":"hostname success"}`,
		}
		MethodTest(t, tc)

		TearDownForGrpc()
	})

}
