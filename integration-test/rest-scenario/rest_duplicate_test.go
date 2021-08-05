package restscenario

import (
	"net/http"
	"testing"
)

func TestRestDuplicate(t *testing.T) {
	t.Run("rest api duplicate test for mock driver", func(t *testing.T) {
		SetUpForRest()

		tc := TestCases{
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
			Name:                 "register cloud driver(duplicate)",
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
			GivenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_duplicate"}]}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		EchoTest(t, tc)

		tc = TestCases{
			Name:                 "register credential(duplicate)",
			EchoFunc:             "RegisterCredential",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/credential",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_duplicate"}]}`,
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
			Name:                 "register region(duplicate)",
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
			Name:                 "create connection config(duplicate)",
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
			Name:                 "create vpc(duplicate)",
			EchoFunc:             "CreateVPC",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vpc",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "vpc-01", "IPv4_CIDR": "192.168.0.0/16", "SubnetInfoList": [ { "Name": "subnet-01", "IPv4_CIDR": "192.168.1.0/24"} ] } }`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `vpc-vpc-01 already exists!`,
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
			Name:                 "add subnet(duplicate)",
			EchoFunc:             "AddSubnet",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vpc/:VPCName/subnet",
			GivenQueryParams:     "",
			GivenParaNames:       []string{"VPCName"},
			GivenParaVals:        []string{"vpc-01"},
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "subnet-02", "IPv4_CIDR": "192.168.2.0/24" } }`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `subnet:vpc-01-subnet-02 already exists!`,
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
			Name:                 "create security(duplicate)",
			EchoFunc:             "CreateSecurity",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/securitygroup",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "sg-01", "VPCName": "vpc-01", "SecurityRules": [ {"FromPort": "1", "ToPort" : "65535", "IPProtocol" : "tcp", "Direction" : "inbound", "CIDR" : "0.0.0.0/0"} ] } }`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `sg-vpc-01-deli-sg-01 already exists!`,
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
			Name:                 "create keypair(duplicate)",
			EchoFunc:             "CreateKey",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/keypair",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "keypair-01" } }`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `keypair-01 already exists!`,
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
			Name:                 "start vm(duplicate)",
			EchoFunc:             "StartVM",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/vm",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{ "ConnectionName": "mock-unit-config01", "ReqInfo": { "Name": "vm-01", "ImageName": "mock-vmimage-01", "VPCName": "vpc-01", "SubnetName": "subnet-01", "SecurityGroupNames": [ "sg-01" ], "VMSpecName": "mock-vmspec-01", "KeyPairName": "keypair-01"} }`,
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `vm-vm-01 already exists!`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

}
