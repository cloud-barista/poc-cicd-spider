package restscenario

import (
	"net/http"
	"testing"
)

func TestRestMisc(t *testing.T) {
	t.Run("rest api misc test for mock driver", func(t *testing.T) {
		SetUpForRest()

		tc := TestCases{
			Name:                 "endpoint info",
			EchoFunc:             "EndpointInfo",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/endpointinfo",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: "\n  <CB-Spider> Multi-Cloud Infrastructure Federation Framework",
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
			Name:                 "register credential",
			EchoFunc:             "RegisterCredential",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/credential",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_misc"}]}`,
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
			Name:                 "create image",
			EchoFunc:             "CreateImage",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vmimage",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "cirros-0.5.1" } }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"cirros-0.5.1","SystemId":"cirros-0.5.1"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list image",
			EchoFunc:             "ListImage",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmimage",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"image":[{"IId":{"NameId":"mock-vmimage-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get image",
			EchoFunc:             "GetImage",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmimage/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"cirros-0.5.1"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"IId":{"NameId":"cirros-0.5.1","SystemId":"cirros-0.5.1"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list vmspec",
			EchoFunc:             "ListVMSpec",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmspec",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"vmspec":[{"Region":"default"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get vmspec",
			EchoFunc:             "GetVMSpec",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmspec/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"mock-vmspec-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Region":"default","Name":"mock-vmspec-01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "list org vmspec",
			EchoFunc:             "ListOrgVMSpec",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmorgspec",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: ``,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "get org vmspec",
			EchoFunc:             "GetOrgVMSpec",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/vmorgspec/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"org-vmspec-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: ``,
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
			Name:                 "list all vpc",
			EchoFunc:             "ListAllVPC",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/allvpc",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"AllList":{"MappedList":[{"NameId":"vpc-01","SystemId":"vpc-01"}]`,
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
			Name:                 "list all security",
			EchoFunc:             "ListAllSecurity",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/allsecuritygroup",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"AllList":{"MappedList":[{"NameId":"sg-01","SystemId":"vpc-01-deli-sg-01"}]`,
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
			Name:                 "list all keypair",
			EchoFunc:             "ListAllKey",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/allkeypair",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"AllList":{"MappedList":[{"NameId":"keypair-01","SystemId":"keypair-01"}]`,
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
			Name:                 "list all vm",
			EchoFunc:             "ListAllVM",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/allvm",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"AllList":{"MappedList":[{"NameId":"vm-01","SystemId":"vm-01"}]`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "terminate csp vm",
			EchoFunc:             "TerminateCSPVM",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/cspvm/:Id",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Id"},
			GivenParaVals:        []string{"vm-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Status":"Terminating"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete csp keypair",
			EchoFunc:             "DeleteCSPKey",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/cspkeypair/:Id",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Id"},
			GivenParaVals:        []string{"keypair-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete csp security",
			EchoFunc:             "DeleteCSPSecurity",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/cspsecuritygroup/:Id",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Id"},
			GivenParaVals:        []string{"vpc-01-deli-sg-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "remove csp subnet",
			EchoFunc:             "RemoveCSPSubnet",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/vpc/:VPCName/cspsubnet/:Id",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"VPCName", "Id"},
			GivenParaVals:        []string{"vpc-01", "subnet-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete csp vpc",
			EchoFunc:             "DeleteCSPVPC",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/cspvpc/:Id",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Id"},
			GivenParaVals:        []string{"vpc-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "delete image",
			EchoFunc:             "DeleteImage",
			HttpMethod:           http.MethodDelete,
			WhenURL:              "/spider/vmimage/:Name",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"Name"},
			GivenParaVals:        []string{"cirros-0.5.1"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01" }`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"Result":"true"}`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

}
