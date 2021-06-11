package grpcscenario

import (
	"testing"
)

func TestGrpcFullDocArg(t *testing.T) {
	t.Run("grpc api full test for mock driver by doccument args style", func(t *testing.T) {
		SetUpForGrpc()

		tc := TestCases{
			name:                "list cloud os",
			instance:            cimApi,
			method:              "ListCloudOS",
			args:                nil,
			expectResStartsWith: `{"cloudos":[`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "register cloud driver",
			instance: cimApi,
			method:   "CreateCloudDriver",
			args: []interface{}{
				`{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`,
			},
			expectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "register credential",
			instance: cimApi,
			method:   "CreateCredential",
			args: []interface{}{
				`{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_name00"}]}`,
			},
			expectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "register region",
			instance: cimApi,
			method:   "CreateRegion",
			args: []interface{}{
				`{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`,
			},
			expectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "create connection config",
			instance: cimApi,
			method:   "CreateConnectionConfig",
			args: []interface{}{
				`{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`,
			},
			expectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			name:     "ssh run",
			instance: ccmApi,
			method:   "SSHRun",
			args: []interface{}{
				`{
					"PrivateKey": [],
					"UserName": "root",
					"ServerPort": "node12:22",
					"Command": "hostname"
				}`,
			},
			expectResStartsWith: `{"Result":"hostname success"}`,
		}
		MethodTest(t, tc)

		TearDownForGrpc()
	})

}
