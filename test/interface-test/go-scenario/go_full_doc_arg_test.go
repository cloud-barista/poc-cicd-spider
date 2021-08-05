package goscenario

import (
	"testing"
)

func TestGoFullDocArg(t *testing.T) {
	t.Run("go api full test for mock driver by doccument args style", func(t *testing.T) {
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
			Method:   "GetCloudDriver",
			Args: []interface{}{
				`{"DriverName":"mock-unit-driver01"}`,
			},
			ExpectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "register credential",
			Instance: CimApi,
			Method:   "CreateCredential",
			Args: []interface{}{
				`{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_full_doc"}]}`,
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
			Method:   "GetCredential",
			Args: []interface{}{
				`{"CredentialName":"mock-unit-credential01"}`,
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
			Method:   "GetRegion",
			Args: []interface{}{
				`{"RegionName":"mock-unit-region01"}`,
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
			Method:   "GetConnectionConfig",
			Args: []interface{}{
				`{"ConfigName":"mock-unit-config01"}`,
			},
			ExpectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
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
			Name:     "add subnet",
			Instance: CcmApi,
			Method:   "AddSubnet",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "VPCName": "vpc-01", "ReqInfo": { "Name": "subnet-02", "IPv4_CIDR": "192.168.2.0/24" } }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vpc",
			Instance: CcmApi,
			Method:   "ListVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vpc":[{"IId":{"NameId":"vpc-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vpc",
			Instance: CcmApi,
			Method:   "GetVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vpc-01" }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
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
			Name:     "list security",
			Instance: CcmApi,
			Method:   "ListSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"securitygroup":[{"IId":{"NameId":"sg-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get security",
			Instance: CcmApi,
			Method:   "GetSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "sg-01" }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"sg-01"`,
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
			Name:     "list keypair",
			Instance: CcmApi,
			Method:   "ListKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"keypair":[{"IId":{"NameId":"keypair-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get keypair",
			Instance: CcmApi,
			Method:   "GetKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "keypair-01" }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"keypair-01"`,
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
			Name:     "list vm",
			Instance: CcmApi,
			Method:   "ListVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vm":[{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vm",
			Instance: CcmApi,
			Method:   "GetVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vm-01" }`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "list vm status",
			Instance: CcmApi,
			Method:   "ListVMStatus",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vmstatus":[{"IId":{"NameId":"vm-01"`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "get vm status",
			Instance: CcmApi,
			Method:   "GetVMStatus",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vm-01" }`,
			},
			ExpectResStartsWith: `{"Status":"Running"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "suspend vm",
			Instance: CcmApi,
			Method:   "ControlVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vm-01", "action": "suspend" }`,
			},
			ExpectResStartsWith: `{"Status":"Suspending"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "resume vm",
			Instance: CcmApi,
			Method:   "ControlVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vm-01", "action": "resume" }`,
			},
			ExpectResStartsWith: `{"Status":"Resuming"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "reboot vm",
			Instance: CcmApi,
			Method:   "ControlVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vm-01", "action": "reboot" }`,
			},
			ExpectResStartsWith: `{"Status":"Rebooting"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "terminate vm",
			Instance: CcmApi,
			Method:   "TerminateVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vm-01", "force": "true" }`,
			},
			ExpectResStartsWith: `{"Status":"Terminating"}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete keypair",
			Instance: CcmApi,
			Method:   "DeleteKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "keypair-01", "force": "true" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete security",
			Instance: CcmApi,
			Method:   "DeleteSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "sg-01", "force": "true" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "remove subnet",
			Instance: CcmApi,
			Method:   "RemoveSubnet",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "VPCName": "vpc-01", "SubnetName": "subnet-02", "force": "true" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete vpc",
			Instance: CcmApi,
			Method:   "DeleteVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01", "Name": "vpc-01", "force": "true" }`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "delete connection config",
			Instance: CimApi,
			Method:   "DeleteConnectionConfig",
			Args: []interface{}{
				`{"ConfigName":"mock-unit-config01"}`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "unregister region",
			Instance: CimApi,
			Method:   "DeleteRegion",
			Args: []interface{}{
				`{"RegionName":"mock-unit-region01"}`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "unregister credential",
			Instance: CimApi,
			Method:   "DeleteCredential",
			Args: []interface{}{
				`{"CredentialName":"mock-unit-credential01"}`,
			},
			ExpectResStartsWith: `{"Result":true}`,
		}
		MethodTest(t, tc)

		tc = TestCases{
			Name:     "unregister cloud driver",
			Instance: CimApi,
			Method:   "DeleteCloudDriver",
			Args: []interface{}{
				`{"DriverName":"mock-unit-driver01"}`,
			},
			ExpectResStartsWith: `{"Result":true}`,
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
