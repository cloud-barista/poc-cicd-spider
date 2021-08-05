package grpcscenario

import (
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"
)

func TestGrpcFullParamArg(t *testing.T) {
	t.Run("grpc api full test for mock driver by parameter args style", func(t *testing.T) {
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
			Name:                "list cloud driver",
			Instance:            CimApi,
			Method:              "ListCloudDriver",
			Args:                nil,
			ExpectResStartsWith: `{"driver":[{"DriverName":"mock-unit-driver01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get cloud driver",
			Instance: CimApi,
			Method:   "GetCloudDriverByParam",
			Args: []interface{}{
				"mock-unit-driver01",
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
						api.KeyValue{Key: "MockName", Value: "mock_unit_full_param"},
					},
				},
			},
			ExpectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:                "list credential",
			Instance:            CimApi,
			Method:              "ListCredential",
			Args:                nil,
			ExpectResStartsWith: `{"credential":[{"CredentialName":"mock-unit-credential01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get credential",
			Instance: CimApi,
			Method:   "GetCredentialByParam",
			Args: []interface{}{
				"mock-unit-credential01",
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
			Name:                "list region",
			Instance:            CimApi,
			Method:              "ListRegion",
			Args:                nil,
			ExpectResStartsWith: `{"region":[{"RegionName":"mock-unit-region01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get region",
			Instance: CimApi,
			Method:   "GetRegionByParam",
			Args: []interface{}{
				"mock-unit-region01",
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
			Name:                "list connection config",
			Instance:            CimApi,
			Method:              "ListConnectionConfig",
			Args:                nil,
			ExpectResStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get connection config",
			Instance: CimApi,
			Method:   "GetConnectionConfigByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create vpc",
			Instance: CcmApi,
			Method:   "CreateVPCByParam",
			Args: []interface{}{
				&api.VPCReq{
					ConnectionName: "mock-unit-config01",
					ReqInfo: api.VPCInfo{
						Name:      "vpc-01",
						IPv4_CIDR: "192.168.0.0/16",
						SubnetInfoList: &[]api.SubnetInfo{
							api.SubnetInfo{
								Name:      "subnet-01",
								IPv4_CIDR: "192.168.1.0/24",
							},
						},
					},
				},
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "add subnet",
			Instance: CcmApi,
			Method:   "AddSubnetByParam",
			Args: []interface{}{
				&api.SubnetReq{
					ConnectionName: "mock-unit-config01",
					VPCName:        "vpc-01",
					ReqInfo: api.SubnetInfo{
						Name:      "subnet-02",
						IPv4_CIDR: "192.168.2.0/24",
					},
				},
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vpc",
			Instance: CcmApi,
			Method:   "ListVPCByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"vpc":[{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vpc",
			Instance: CcmApi,
			Method:   "GetVPCByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vpc-01",
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create security",
			Instance: CcmApi,
			Method:   "CreateSecurityByParam",
			Args: []interface{}{
				&api.SecurityReq{
					ConnectionName: "mock-unit-config01",
					ReqInfo: api.SecurityInfo{
						Name:    "sg-01",
						VPCName: "vpc-01",
						SecurityRules: &[]api.SecurityRuleInfo{
							api.SecurityRuleInfo{
								FromPort:   "1",
								ToPort:     "65535",
								IPProtocol: "tcp",
								Direction:  "inbound",
								CIDR:       "0.0.0.0/0",
							},
						},
					},
				},
			},
			ExpectResStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list security",
			Instance: CcmApi,
			Method:   "ListSecurityByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"securitygroup":[{"IId":{"NameId":"sg-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get security",
			Instance: CcmApi,
			Method:   "GetSecurityByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"sg-01",
			},
			ExpectResStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create keypair",
			Instance: CcmApi,
			Method:   "CreateKeyByParam",
			Args: []interface{}{
				&api.KeyReq{
					ConnectionName: "mock-unit-config01",
					ReqInfo: api.KeyInfo{
						Name: "keypair-01",
					},
				},
			},
			ExpectResStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list keypair",
			Instance: CcmApi,
			Method:   "ListKeyByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"keypair":[{"IId":{"NameId":"keypair-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get keypair",
			Instance: CcmApi,
			Method:   "GetKeyByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"keypair-01",
			},
			ExpectResStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "start vm",
			Instance: CcmApi,
			Method:   "StartVMByParam",
			Args: []interface{}{
				&api.VMReq{
					ConnectionName: "mock-unit-config01",
					ReqInfo: api.VMInfo{
						Name:               "vm-01",
						ImageName:          "mock-vmimage-01",
						VPCName:            "vpc-01",
						SubnetName:         "subnet-01",
						SecurityGroupNames: []string{"sg-01"},
						VMSpecName:         "mock-vmspec-01",
						KeyPairName:        "keypair-01",
					},
				},
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vm",
			Instance: CcmApi,
			Method:   "ListVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"vm":[{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vm",
			Instance: CcmApi,
			Method:   "GetVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vm status",
			Instance: CcmApi,
			Method:   "ListVMStatusByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"vmstatus":[{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vm status",
			Instance: CcmApi,
			Method:   "GetVMStatusByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
			},
			ExpectResStartsWith: `{"Status":"Running"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "suspend vm",
			Instance: CcmApi,
			Method:   "ControlVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
				"suspend",
			},
			ExpectResStartsWith: `{"Status":"Suspending"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "resume vm",
			Instance: CcmApi,
			Method:   "ControlVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
				"resume",
			},
			ExpectResStartsWith: `{"Status":"Resuming"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "reboot vm",
			Instance: CcmApi,
			Method:   "ControlVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
				"reboot",
			},
			ExpectResStartsWith: `{"Status":"Rebooting"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "terminate vm",
			Instance: CcmApi,
			Method:   "TerminateVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
				"true",
			},
			ExpectResStartsWith: `{"Status":"Terminating"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete keypair",
			Instance: CcmApi,
			Method:   "DeleteKeyByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"keypair-01",
				"true",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete security",
			Instance: CcmApi,
			Method:   "DeleteSecurityByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"sg-01",
				"true",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "remove subnet",
			Instance: CcmApi,
			Method:   "RemoveSubnetByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vpc-01",
				"subnet-02",
				"true",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete vpc",
			Instance: CcmApi,
			Method:   "DeleteVPCByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vpc-01",
				"true",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete connection config",
			Instance: CimApi,
			Method:   "DeleteConnectionConfigByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "unregister region",
			Instance: CimApi,
			Method:   "DeleteRegionByParam",
			Args: []interface{}{
				"mock-unit-region01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "unregister credential",
			Instance: CimApi,
			Method:   "DeleteCredentialByParam",
			Args: []interface{}{
				"mock-unit-credential01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "unregister cloud driver",
			Instance: CimApi,
			Method:   "DeleteCloudDriverByParam",
			Args: []interface{}{
				"mock-unit-driver01",
			},
			ExpectResStartsWith: `{"Result":true}`,
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
