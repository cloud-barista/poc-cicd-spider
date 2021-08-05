package cliscenario

import (
	"testing"
)

func TestCliMiscYaml(t *testing.T) {
	t.Run("command misc yaml in/out test for mock driver", func(t *testing.T) {
		SetUpForCli()

		tc := TestCases{
			Name: "register cloud driver",
			CmdArgs: []string{"driver", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"DriverName": "mock-unit-driver01"
"ProviderName": "MOCK"
"DriverLibFileName": "mock-driver-v1.0.so"
`,
			},
			ExpectResStartsWith: `DriverName: mock-unit-driver01
ProviderName: MOCK`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "register credential",
			CmdArgs: []string{"credential", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"CredentialName": "mock-unit-credential01"
"ProviderName": "MOCK"
"KeyValueInfoList":
  - "Key": "MockName"
    "Value": "mock_unit_misc_yaml"
`,
			},
			ExpectResStartsWith: `CredentialName: mock-unit-credential01
ProviderName: MOCK`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "register region",
			CmdArgs: []string{"region", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"RegionName": "mock-unit-region01"
"ProviderName": "MOCK"
"KeyValueInfoList":
  - "Key": "Region"
    "Value": "default"
`,
			},
			ExpectResStartsWith: `RegionName: mock-unit-region01
ProviderName: MOCK`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create connection config",
			CmdArgs: []string{"connection", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"ConfigName": "mock-unit-config01"
"ProviderName": "MOCK"
"DriverName": "mock-unit-driver01"
"CredentialName": "mock-unit-credential01"
"RegionName": "mock-unit-region01"
`,
			},
			ExpectResStartsWith: `ConfigName: mock-unit-config01
ProviderName: MOCK`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create image",
			CmdArgs: []string{"image", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"ConnectionName": "mock-unit-config01"
"ReqInfo":
  "Name": "cirros-0.5.1"
`,
			},
			ExpectResStartsWith: `IId:
  NameId: cirros-0.5.1`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "list image",
			CmdArgs: []string{"image", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `image:
- IId:`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "get image",
			CmdArgs: []string{"image", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--name", "cirros-0.5.1"},
			ExpectResStartsWith: `IId:
  NameId: cirros-0.5.1`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "list vmspec",
			CmdArgs: []string{"vmspec", "list", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `vmspec:
- Region: default`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "get vmspec",
			CmdArgs: []string{"vmspec", "get", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--name", "mock-vmspec-01"},
			ExpectResStartsWith: `Region: default
Name: mock-vmspec-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list org vmspec",
			CmdArgs:             []string{"vmspec", "listorg", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: "\n",
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get org vmspec",
			CmdArgs:             []string{"vmspec", "getorg", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--name", "org-vmspec-01"},
			ExpectResStartsWith: "\n",
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create vpc",
			CmdArgs: []string{"vpc", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"ConnectionName": "mock-unit-config01"
"ReqInfo":
  "Name": "vpc-01" 
  "IPv4_CIDR": "192.168.0.0/16" 
  "SubnetInfoList": 
    - "Name": "subnet-01" 
      "IPv4_CIDR": "192.168.1.0/24"
`,
			},
			ExpectResStartsWith: `IId:
  NameId: vpc-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "list all vpc",
			CmdArgs: []string{"vpc", "listall", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `AllList:
  MappedList:
  - NameId: vpc-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create security",
			CmdArgs: []string{"security", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"ConnectionName": "mock-unit-config01"
"ReqInfo":
  "Name": "sg-01"
  "VPCName": "vpc-01"
  "SecurityRules":
    - "FromPort": "1"
      "ToPort": "65535"
      "IPProtocol": "tcp"
      "Direction": "inbound"
      "CIDR": "0.0.0.0/0"
`,
			},
			ExpectResStartsWith: `IId:
  NameId: sg-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "list all security",
			CmdArgs: []string{"security", "listall", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `AllList:
  MappedList:
  - NameId: sg-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create keypair",
			CmdArgs: []string{"keypair", "create", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"ConnectionName": "mock-unit-config01"
"ReqInfo":
  "Name": "keypair-01"
`,
			},
			ExpectResStartsWith: `IId:
  NameId: keypair-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "list all keypair",
			CmdArgs: []string{"keypair", "listall", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `AllList:
  MappedList:
  - NameId: keypair-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "start vm",
			CmdArgs: []string{"vm", "start", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "-d", `
"ConnectionName": "mock-unit-config01"
"ReqInfo":
  "Name": "vm-01"
  "ImageName": "mock-vmimage-01"
  "VPCName": "vpc-01"
  "SubnetName": "subnet-01"
  "SecurityGroupNames":
    - "sg-01" 
  "VMSpecName": "mock-vmspec-01"
  "KeyPairName": "keypair-01"
`,
			},
			ExpectResStartsWith: `IId:
  NameId: vm-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:    "list all vm",
			CmdArgs: []string{"vm", "listall", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `AllList:
  MappedList:
  - NameId: vm-01`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "terminate csp vm",
			CmdArgs:             []string{"vm", "terminatecsp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--id", "vm-01"},
			ExpectResStartsWith: `Status: Terminating`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete csp keypair",
			CmdArgs:             []string{"keypair", "deletecsp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--id", "keypair-01"},
			ExpectResStartsWith: `Result: true`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete csp security",
			CmdArgs:             []string{"security", "deletecsp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--id", "vpc-01-deli-sg-01"},
			ExpectResStartsWith: `Result: true`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "remove csp subnet",
			CmdArgs:             []string{"vpc", "removecsp-subnet", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--vname", "vpc-01", "--id", "subnet-01"},
			ExpectResStartsWith: `Result: true`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete csp vpc",
			CmdArgs:             []string{"vpc", "deletecsp", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--id", "vpc-01"},
			ExpectResStartsWith: `Result: true`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete image",
			CmdArgs:             []string{"image", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "yaml", "-o", "yaml", "--cname", "mock-unit-config01", "--name", "cirros-0.5.1"},
			ExpectResStartsWith: `Result: true`,
		}
		SpiderCmdTest(t, tc)

		TearDownForCli()
	})
}
