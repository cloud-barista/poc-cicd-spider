package rgequalscenario

import (
	"net/http"
	"testing"

	gs "github.com/cloud-barista/poc-cicd-spider/test/interface-test/go-scenario"
	rs "github.com/cloud-barista/poc-cicd-spider/test/interface-test/rest-scenario"
)

func TestPrepareFull(t *testing.T) {
	t.Run("prepare full test for mock driver", func(t *testing.T) {
		gs.SetUpForGrpc()

		tc := rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list cloud driver for rest",
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
		res, _ := rs.EchoTest(t, tc)
		RestPrepareResult["list cloud driver"] = res

		gtc := gs.TestCases{
			Name:                "list cloud driver for grpc",
			Instance:            gs.CimApi,
			Method:              "ListCloudDriver",
			Args:                nil,
			ExpectResStartsWith: `{"driver":[{"DriverName":"mock-unit-driver01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list cloud driver"] = res

		tc = rs.TestCases{
			Name:                 "register credential",
			EchoFunc:             "RegisterCredential",
			HttpMethod:           http.MethodPost,
			WhenURL:              "/spider/credential",
			GivenQueryParams:     "",
			GivenParaNames:       nil,
			GivenParaVals:        nil,
			GivenPostData:        `{"CredentialName":"mock-unit-credential01","ProviderName":"MOCK", "KeyValueInfoList": [{"Key":"MockName", "Value":"mock_unit_equal"}]}`,
			ExpectStatus:         http.StatusOK,
			ExpectBodyStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list region for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list region"] = res

		gtc = gs.TestCases{
			Name:                "list region for grpc",
			Instance:            gs.CimApi,
			Method:              "ListRegion",
			Args:                nil,
			ExpectResStartsWith: `{"region":[{"RegionName":"mock-unit-region01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list region"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list connection config for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list connection config"] = res

		gtc = gs.TestCases{
			Name:                "list connection config for grpc",
			Instance:            gs.CimApi,
			Method:              "ListConnectionConfig",
			Args:                nil,
			ExpectResStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list connection config"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list vpc for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list vpc"] = res

		gtc = gs.TestCases{
			Name:     "list vpc for grpc",
			Instance: gs.CcmApi,
			Method:   "ListVPC",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vpc":[{"IId":{"NameId":"vpc-01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list vpc"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list security for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list security"] = res

		gtc = gs.TestCases{
			Name:     "list security for grpc",
			Instance: gs.CcmApi,
			Method:   "ListSecurity",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"securitygroup":[{"IId":{"NameId":"sg-01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list security"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list keypair for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list keypair"] = res

		gtc = gs.TestCases{
			Name:     "list keypair for grpc",
			Instance: gs.CcmApi,
			Method:   "ListKey",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"keypair":[{"IId":{"NameId":"keypair-01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list keypair"] = res

		tc = rs.TestCases{
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
		rs.EchoTest(t, tc)

		tc = rs.TestCases{
			Name:                 "list vm for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list vm"] = res

		gtc = gs.TestCases{
			Name:     "list vm for grpc",
			Instance: gs.CcmApi,
			Method:   "ListVM",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vm":[{"IId":{"NameId":"vm-01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list vm"] = res

		tc = rs.TestCases{
			Name:                 "list vm status for rest",
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
		res, _ = rs.EchoTest(t, tc)
		RestPrepareResult["list vm status"] = res

		gtc = gs.TestCases{
			Name:     "list vm status for grpc",
			Instance: gs.CcmApi,
			Method:   "ListVMStatus",
			Args: []interface{}{
				`{ "ConnectionName": "mock-unit-config01" }`,
			},
			ExpectResStartsWith: `{"vmstatus":[{"IId":{"NameId":"vm-01"`,
		}
		res, _ = gs.MethodTest(t, gtc)
		GrpcPrepareResult["list vm status"] = res

		gs.TearDownForGrpc()
	})

}
