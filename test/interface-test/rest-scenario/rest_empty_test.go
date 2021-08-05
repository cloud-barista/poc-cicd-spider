package restscenario

import (
	"net/http"
	"testing"
)

func TestRestEmpty(t *testing.T) {
	t.Run("rest api empty test for mock driver", func(t *testing.T) {
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
			Name:                 "list cloud driver",
			EchoFunc:             "ListCloudDriver",
			HttpMethod:           http.MethodGet,
			WhenURL:              "/spider/driver",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        "",
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"driver":[]}`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-driver01: does not exist!`,
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
			ExpectBodyStartsWith: `{"credential":[]}`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-credential01: does not exist!`,
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
			ExpectBodyStartsWith: `{"region":[]}`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-region01: does not exist!`,
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
			ExpectBodyStartsWith: `{"connectionconfig":[]}`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-config01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-region01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-credential01: does not exist!`,
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
			ExpectStatus:         http.StatusInternalServerError,
			ExpectBodyStartsWith: `mock-unit-driver01: does not exist!`,
		}
		EchoTest(t, tc)

		TearDownForRest()
	})

}
