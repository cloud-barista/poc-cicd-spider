package grpcscenario

import (
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"
)

func TestGrpcApi(t *testing.T) {
	t.Run("grpc api full test for mock driver by doccument args style", func(t *testing.T) {
		setUpForGrpc()

		tc := TestCases{
			name:                "list cloud os",
			method:              "ListCloudOS",
			args:                nil,
			expectResStartsWith: `{"cloudos":[`,
		}
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "register cloud driver",
			method: "CreateCloudDriver",
			args: []interface{}{
				`{"DriverName":"mock-unit-driver01","ProviderName":"MOCK", "DriverLibFileName":"mock-driver-v1.0.so"}`,
			},
			expectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "register credential",
			method: "CreateCredential",
			args: []interface{}{
				`{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_name00"}]}`,
			},
			expectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "register region",
			method: "CreateRegion",
			args: []interface{}{
				`{"RegionName":"mock-unit-region01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"Region", "Value":"default"}]}`,
			},
			expectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "create connection config",
			method: "CreateConnectionConfig",
			args: []interface{}{
				`{"ConfigName":"mock-unit-config01","ProviderName":"MOCK", "DriverName":"mock-unit-driver01", "CredentialName":"mock-unit-credential01", "RegionName":"mock-unit-region01"}`,
			},
			expectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		CIMMethodTest(t, cim, tc)

		tearDownForGrpc()
	})

	t.Run("grpc api full test for mock driver by parameter args style", func(t *testing.T) {
		setUpForGrpc()

		tc := TestCases{
			name:   "register cloud driver",
			method: "CreateCloudDriverByParam",
			args: []interface{}{
				&api.CloudDriverReq{
					DriverName:        "mock-unit-driver01",
					ProviderName:      "MOCK",
					DriverLibFileName: "mock-driver-v1.0.so",
				},
			},
			expectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "register credential",
			method: "CreateCredentialByParam",
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
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "register region",
			method: "CreateRegionByParam",
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
		CIMMethodTest(t, cim, tc)

		tc = TestCases{
			name:   "create connection config",
			method: "CreateConnectionConfigByParam",
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
		CIMMethodTest(t, cim, tc)

		tearDownForGrpc()
	})

	t.Run("sshrun test", func(t *testing.T) {
		setUpForGrpc()

		tc := TestCases{
			name:   "ssh run by document args style",
			method: "SSHRun",
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
		CCMMethodTest(t, ccm, tc)

		tc = TestCases{
			name:   "ssh run by parameter args style",
			method: "SSHRunByParam",
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
		CCMMethodTest(t, ccm, tc)

		tearDownForGrpc()
	})
}
