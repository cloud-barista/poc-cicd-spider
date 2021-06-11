package grpcscenario

import (
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"
)

func TestGrpcFullParamArg(t *testing.T) {
	t.Run("grpc api full test for mock driver by parameter args style", func(t *testing.T) {
		SetUpForGrpc()

		tc := TestCases{
			Name:     "register cloud driver",
			Instance: CimApi,
			Method:   "CreateCloudDriverByParam",
			Args: []interface{}{
				&api.CloudDriverReq{
					DriverName:        "mock-unit-driver01",
					ProviderName:      "MOCK",
					DriverLibFileName: "mock-driver-v1.0.so",
				},
			},
			ExpectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register credential",
			Instance: CimApi,
			Method:   "CreateCredentialByParam",
			Args: []interface{}{
				&api.CredentialReq{
					CredentialName: "mock-unit-credential01",
					ProviderName:   "MOCK",
					KeyValueInfoList: []api.KeyValue{
						api.KeyValue{Key: "MockName", Value: "mock_unit_name00"},
					},
				},
			},
			ExpectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register region",
			Instance: CimApi,
			Method:   "CreateRegionByParam",
			Args: []interface{}{
				&api.RegionReq{
					RegionName:   "mock-unit-region01",
					ProviderName: "MOCK",
					KeyValueInfoList: []api.KeyValue{
						api.KeyValue{Key: "Region", Value: "default"},
					},
				},
			},
			ExpectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create connection config",
			Instance: CimApi,
			Method:   "CreateConnectionConfigByParam",
			Args: []interface{}{
				&api.ConnectionConfigReq{
					ConfigName:     "mock-unit-config01",
					ProviderName:   "MOCK",
					DriverName:     "mock-unit-driver01",
					CredentialName: "mock-unit-credential01",
					RegionName:     "mock-unit-region01",
				},
			},
			ExpectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "ssh run",
			Instance: CcmApi,
			Method:   "SSHRunByParam",
			Args: []interface{}{
				&api.SSHRUNReq{
					PrivateKey: []string{},
					UserName:   "root",
					ServerPort: "node12:22",
					Command:    "hostname",
				},
			},
			ExpectResStartsWith: `{"Result":"hostname success"}`,
		}
		MethodTest(t, tc)

		TearDownForGrpc()
	})

}
