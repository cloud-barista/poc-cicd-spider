package goscenario

import (
	"testing"
)

func TestGoMiscDocArg(t *testing.T) {
	t.Run("go api misc test for mock driver by document args style", func(t *testing.T) {
		SetUpForGrpc()

		tc := TestCases{
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
				`{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_misc_doc"}]}`,
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
			Name:     "create image",
			Instance: CcmApi,
			Method:   "CreateImage",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "cirros-0.5.1" } }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"cirros-0.5.1","SystemId":"cirros-0.5.1"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list image",
			Instance: CcmApi,
			Method:   "ListImage",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"image":[{"IId":{"NameId":"mock-vmimage-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get image",
			Instance: CcmApi,
			Method:   "GetImage",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "cirros-0.5.1" }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"cirros-0.5.1","SystemId":"cirros-0.5.1"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vmspec",
			Instance: CcmApi,
			Method:   "ListVMSpec",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vmspec":[{"Region":"default"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vmspec",
			Instance: CcmApi,
			Method:   "GetVMSpec",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "mock-vmspec-01" }`,
			},
			ExpectResStartsWith: `{"Region":"default","Name":"mock-vmspec-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list org vmspec",
			Instance: CcmApi,
			Method:   "ListOrgVMSpec",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: ``,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get org vmspec",
			Instance: CcmApi,
			Method:   "GetOrgVMSpec",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "org-vmspec-01" }`,
			},
			ExpectResStartsWith: ``,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create vpc",
			Instance: CcmApi,
			Method:   "CreateVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "vpc-01", "IPv4_CIDR": "192.168.0.0/16", "SubnetInfoList": [ { "Name": "subnet-01", "IPv4_CIDR": "192.168.1.0/24"} ] } }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list all vpc",
			Instance: CcmApi,
			Method:   "ListAllVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"vpc-01","SystemId":"vpc-01"}]`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create security",
			Instance: CcmApi,
			Method:   "CreateSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "sg-01", "VPCName": "vpc-01", "SecurityRules": [ {"FromPort": "1", "ToPort" : "65535", "IPProtocol" : "tcp", "Direction" : "inbound", "CIDR" : "0.0.0.0/0"} ] } }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list all security",
			Instance: CcmApi,
			Method:   "ListAllSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"sg-01","SystemId":"vpc-01-deli-sg-01"}]`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "create keypair",
			Instance: CcmApi,
			Method:   "CreateKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "keypair-01" } }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list all keypair",
			Instance: CcmApi,
			Method:   "ListAllKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"keypair-01","SystemId":"keypair-01"}]`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "start vm",
			Instance: CcmApi,
			Method:   "StartVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "vm-01", "ImageName": "mock-vmimage-01", "VPCName": "vpc-01", "SubnetName": "subnet-01", "SecurityGroupNames": [ "sg-01" ], "VMSpecName": "mock-vmspec-01", "KeyPairName": "keypair-01"} }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list all vm",
			Instance: CcmApi,
			Method:   "ListAllVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"AllList":{"MappedList":[{"NameId":"vm-01","SystemId":"vm-01"}]`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "terminate csp vm",
			Instance: CcmApi,
			Method:   "TerminateCSPVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Id": "vm-01" }`,
			},
			ExpectResStartsWith: `{"Status":"Terminating"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete csp keypair",
			Instance: CcmApi,
			Method:   "DeleteCSPKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Id": "keypair-01" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete csp security",
			Instance: CcmApi,
			Method:   "DeleteCSPSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Id": "vpc-01-deli-sg-01" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "remove csp subnet",
			Instance: CcmApi,
			Method:   "RemoveCSPSubnet",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "VPCName": "vpc-01", "Id": "subnet-01" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete csp vpc",
			Instance: CcmApi,
			Method:   "DeleteCSPVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Id": "vpc-01" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete image",
			Instance: CcmApi,
			Method:   "DeleteImage",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "cirros-0.5.1" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		TearDownForGrpc()
	})
}
