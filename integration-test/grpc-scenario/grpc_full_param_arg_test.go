package grpcscenario

import (
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"
)

func TestGrpcFullParamArg(t *testing.T) {
	t.Run("grpc api full test for mock driver by parameter args style", func(t *testing.T) {
		setUpForGrpc()

		tc := TestCases{
			name:     "register cloud driver",
			instance: cimApi,
			method:   "CreateCloudDriverByParam",
			args: []interface{}{
				&api.CloudDriverReq{
					DriverName:        "mock-unit-driver01",
					ProviderName:      "MOCK",
					DriverLibFileName: "mock-driver-v1.0.so",
				},
			},
			expectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "register credential",
			instance: cimApi,
			method:   "CreateCredentialByParam",
			args: []interface{}{
				&api.CredentialReq{
					CredentialName: "mock-unit-credential01",
					ProviderName:   "MOCK",
					KeyValueInfoList: []api.KeyValue{
						api.KeyValue{Key: "MockName", Value: "mock_unit_name00"},
					},
				},
			},
			expectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "register region",
			instance: cimApi,
			method:   "CreateRegionByParam",
			args: []interface{}{
				&api.RegionReq{
					RegionName:   "mock-unit-region01",
					ProviderName: "MOCK",
					KeyValueInfoList: []api.KeyValue{
						api.KeyValue{Key: "Region", Value: "default"},
					},
				},
			},
			expectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "create connection config",
			instance: cimApi,
			method:   "CreateConnectionConfigByParam",
			args: []interface{}{
				&api.ConnectionConfigReq{
					ConfigName:     "mock-unit-config01",
					ProviderName:   "MOCK",
					DriverName:     "mock-unit-driver01",
					CredentialName: "mock-unit-credential01",
					RegionName:     "mock-unit-region01",
				},
			},
			expectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "ssh run",
			instance: ccmApi,
			method:   "SSHRunByParam",
			args: []interface{}{
				&api.SSHRUNReq{
					PrivateKey: []string{},
					UserName:   "root",
					ServerPort: "node12:22",
					Command:    "hostname",
				},
			},
			expectResStartsWith: `{"Result":"hostname success"}`,
		}
		MethodTest(t, tc)

		tearDownForGrpc()
	})

}
