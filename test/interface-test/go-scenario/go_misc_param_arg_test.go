package goscenario

import (
	"testing"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"
)

func TestGoMiscParamArg(t *testing.T) {
	t.Run("go api misc test for mock driver by parameter args style", func(t *testing.T) {
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
						api.KeyValue{Key: "MockName", Value: "mock_unit_misc_param"},
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
			Name:     "create image",
			Instance: CcmApi,
			Method:   "CreateImageByParam",
			Args: []interface{}{
				&api.ImageReq{
					ConnectionName: "mock-unit-config01",
					ReqInfo: api.ImageInfo{
						Name: "cirros-0.5.1",
					},
				},
			},
			ExpectResStartsWith: `{"IId":{"NameId":"cirros-0.5.1","SystemId":"cirros-0.5.1"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list image",
			Instance: CcmApi,
			Method:   "ListImageByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"image":[{"IId":{"NameId":"mock-vmimage-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get image",
			Instance: CcmApi,
			Method:   "GetImageByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"cirros-0.5.1",
			},
			ExpectResStartsWith: `{"IId":{"NameId":"cirros-0.5.1","SystemId":"cirros-0.5.1"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vmspec",
			Instance: CcmApi,
			Method:   "ListVMSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"vmspec":[{"Region":"default"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vmspec",
			Instance: CcmApi,
			Method:   "GetVMSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"mock-vmspec-01",
			},
			ExpectResStartsWith: `{"Region":"default","Name":"mock-vmspec-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list org vmspec",
			Instance: CcmApi,
			Method:   "ListOrgVMSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: ``,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get org vmspec",
			Instance: CcmApi,
			Method:   "GetOrgVMSpecByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"org-vmspec-01",
			},
			ExpectResStartsWith: ``,
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
			Name:     "list all vpc",
			Instance: CcmApi,
			Method:   "ListAllVPCByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"vpc-01","SystemId":"vpc-01"}]`,
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
			Name:     "list all security",
			Instance: CcmApi,
			Method:   "ListAllSecurityByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"sg-01","SystemId":"vpc-01-deli-sg-01"}]`,
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
			Name:     "list all keypair",
			Instance: CcmApi,
			Method:   "ListAllKeyByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"keypair-01","SystemId":"keypair-01"}]`,
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
			Name:     "list all vm",
			Instance: CcmApi,
			Method:   "ListAllVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"vm-01","SystemId":"vm-01"}]`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "terminate csp vm",
			Instance: CcmApi,
			Method:   "TerminateCSPVMByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vm-01",
			},
			ExpectResStartsWith: `{"Status":"Terminating"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete csp keypair",
			Instance: CcmApi,
			Method:   "DeleteCSPKeyByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"keypair-01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete csp security",
			Instance: CcmApi,
			Method:   "DeleteCSPSecurityByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vpc-01-deli-sg-01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "remove csp subnet",
			Instance: CcmApi,
			Method:   "RemoveCSPSubnetByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vpc-01",
				"subnet-01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete csp vpc",
			Instance: CcmApi,
			Method:   "DeleteCSPVPCByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"vpc-01",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete image",
			Instance: CcmApi,
			Method:   "DeleteImageByParam",
			Args: []interface{}{
				"mock-unit-config01",
				"cirros-0.5.1",
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		TearDownForGrpc()
	})
}
