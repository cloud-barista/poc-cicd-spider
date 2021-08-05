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
			Name:                 "list cloud driver",
			EchoFunc:             "ListCloudDriver",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/driver",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"driver":[{"DriverName":"mock-unit-driver01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get cloud driver",
			EchoFunc:             "GetCloudDriver",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/driver/:DriverName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"DriverName"},
			GivenParaVals:        []string{"mock-unit-driver01"},
			GivenPostData:        "",
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
			Name:                 "list credential",
			EchoFunc:             "ListCredential",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/credential",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"credential":[{"CredentialName":"mock-unit-credential01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get credential",
			EchoFunc:             "GetCredential",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/credential/:CredentialName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"CredentialName"},
			GivenParaVals:        []string{"mock-unit-credential01"},
			GivenPostData:        "",
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
			Name:                 "list region",
			EchoFunc:             "ListRegion",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/region",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"region":[{"RegionName":"mock-unit-region01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get region",
			EchoFunc:             "GetRegion",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/region/:RegionName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"RegionName"},
			GivenParaVals:        []string{"mock-unit-region01"},
			GivenPostData:        "",
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

		tc = TestCases{
			Name:                 "list connection config",
			EchoFunc:             "ListConnectionConfig",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/connectionconfig",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get connection config",
			EchoFunc:             "GetConnectionConfig",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/connectionconfig/:ConfigName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"ConfigName"},
			GivenParaVals:        []string{"mock-unit-config01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "create vpc",
			EchoFunc:             "CreateVPC",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vpc",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "vpc-01", "IPv4_CIDR": "192.168.0.0/16", "SubnetInfoList": [ { "Name": "subnet-01", "IPv4_CIDR": "192.168.1.0/24"} ] } }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "add subnet",
			EchoFunc:             "AddSubnet",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vpc/:VPCName/subnet",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"VPCName"},
			GivenParaVals:        []string{"vpc-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "subnet-02", "IPv4_CIDR": "192.168.2.0/24" } }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list vpc",
			EchoFunc:             "ListVPC",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vpc",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"vpc":[{"IId":{"NameId":"vpc-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get vpc",
			EchoFunc:             "GetVPC",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vpc/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vpc-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "create security",
			EchoFunc:             "CreateSecurity",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/securitygroup",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "sg-01", "VPCName": "vpc-01", "SecurityRules": [ {"FromPort": "1", "ToPort" : "65535", "IPProtocol" : "tcp", "Direction" : "inbound", "CIDR" : "0.0.0.0/0"} ] } }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list security",
			EchoFunc:             "ListSecurity",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/securitygroup",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"securitygroup":[{"IId":{"NameId":"sg-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get security",
			EchoFunc:             "GetSecurity",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/securitygroup/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"sg-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "create keypair",
			EchoFunc:             "CreateKey",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/keypair",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "keypair-01" } }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list keypair",
			EchoFunc:             "ListKey",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/keypair",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"keypair":[{"IId":{"NameId":"keypair-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get keypair",
			EchoFunc:             "GetKey",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/keypair/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"keypair-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "start vm",
			EchoFunc:             "StartVM",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vm",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "vm-01", "ImageName": "mock-vmimage-01", "VPCName": "vpc-01", "SubnetName": "subnet-01", "SecurityGroupNames": [ "sg-01" ], "VMSpecName": "mock-vmspec-01", "KeyPairName": "keypair-01"} }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list vm",
			EchoFunc:             "ListVM",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vm",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"vm":[{"IId":{"NameId":"vm-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get vm",
			EchoFunc:             "GetVM",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vm/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list vm status",
			EchoFunc:             "ListVMStatus",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmstatus",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"vmstatus":[{"IId":{"NameId":"vm-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get vm status",
			EchoFunc:             "GetVMStatus",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmstatus/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Status":"Running"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "suspend vm",
			EchoFunc:             "ControlVM",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/controlvm/:Name",
			GivenQueryParams:     "?action=suspend",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Status":"Suspending"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "resume vm",
			EchoFunc:             "ControlVM",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/controlvm/:Name",
			GivenQueryParams:     "?action=resume",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Status":"Resuming"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "reboot vm",
			EchoFunc:             "ControlVM",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/controlvm/:Name",
			GivenQueryParams:     "?action=reboot",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Status":"Rebooting"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "terminate vm",
			EchoFunc:             "TerminateVM",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/vm/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Status":"Terminating"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete keypair",
			EchoFunc:             "DeleteKey",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/keypair/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"keypair-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete security",
			EchoFunc:             "DeleteSecurity",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/securitygroup/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"sg-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "remove subnet",
			EchoFunc:             "RemoveSubnet",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/vpc/:VPCName/subnet/:SubnetName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"VPCName", "SubnetName"},
			GivenParaVals:        []string{"vpc-01", "subnet-02"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete vpc",
			EchoFunc:             "DeleteVPC",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/vpc/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"vpc-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete connection config",
			EchoFunc:             "DeleteConnectionConfig",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/connectionconfig/:ConfigName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"ConfigName"},
			GivenParaVals:        []string{"mock-unit-config01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "unregister region",
			EchoFunc:             "UnRegisterRegion",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/region/:RegionName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"RegionName"},
			GivenParaVals:        []string{"mock-unit-region01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "unregister credential",
			EchoFunc:             "UnRegisterCredential",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/credential/:CredentialName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"CredentialName"},
			GivenParaVals:        []string{"mock-unit-credential01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "unregister cloud driver",
			EchoFunc:             "UnRegisterCloudDriver",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/driver/:DriverName",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"DriverName"},
			GivenParaVals:        []string{"mock-unit-driver01"},
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
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
