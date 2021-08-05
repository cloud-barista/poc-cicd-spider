package cliscenario

import (
	"testing"
)

func TestCliFullJson(t *testing.T) {
	t.Run("command full json in/out test for mock driver", func(t *testing.T) {
		SetUpForCli()

		tc := TestCases{
			Name:                "list cloud os",
			CmdArgs:             []string{"os", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			ExpectResStartsWith: `{"cloudos":[`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "register cloud driver",
			CmdArgs: []string{"driver", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"DriverName":"mock-unit-driver01",
					"ProviderName":"MOCK", 
					"DriverLibFileName":"mock-driver-v1.0.so"
				}`,
			},
			ExpectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list cloud driver",
			CmdArgs:             []string{"driver", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			ExpectResStartsWith: `{"driver":[{"DriverName":"mock-unit-driver01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get cloud driver",
			CmdArgs:             []string{"driver", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-driver01"},
			ExpectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "register credential",
			CmdArgs: []string{"credential", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"CredentialName":"mock-unit-credential01",
					"ProviderName":"MOCK",
					"KeyValueInfoList": [
						{"Key":"MockName", "Value":"mock_unit_full_json"}
					]
				}`,
			},
			ExpectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list credential",
			CmdArgs:             []string{"credential", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			ExpectResStartsWith: `{"credential":[{"CredentialName":"mock-unit-credential01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get credential",
			CmdArgs:             []string{"credential", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-credential01"},
			ExpectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "register region",
			CmdArgs: []string{"region", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"RegionName":"mock-unit-region01",
					"ProviderName":"MOCK",
					"KeyValueInfoList": [
						{"Key":"Region", "Value":"default"}
					]
				}`,
			},
			ExpectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list region",
			CmdArgs:             []string{"region", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			ExpectResStartsWith: `{"region":[{"RegionName":"mock-unit-region01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get region",
			CmdArgs:             []string{"region", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-region01"},
			ExpectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create connection config",
			CmdArgs: []string{"connection", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConfigName":"mock-unit-config01",
					"ProviderName":"MOCK", 
					"DriverName":"mock-unit-driver01", 
					"CredentialName":"mock-unit-credential01", 
					"RegionName":"mock-unit-region01"
				}`,
			},
			ExpectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list connection config",
			CmdArgs:             []string{"connection", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			ExpectResStartsWith: `{"connectionconfig":[{"ConfigName":"mock-unit-config01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get connection config",
			CmdArgs:             []string{"connection", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-config01"},
			ExpectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create vpc",
			CmdArgs: []string{"vpc", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConnectionName":"mock-unit-config01",
					"ReqInfo": { 
						"Name": "vpc-01", 
						"IPv4_CIDR": "192.168.0.0/16", 
						"SubnetInfoList": [ 
							{ 
								"Name": "subnet-01", 
								"IPv4_CIDR": "192.168.1.0/24"
							} 
						] 
					} 
				}`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "add subnet",
			CmdArgs: []string{"vpc", "add-subnet", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConnectionName":"mock-unit-config01",
					"VPCName": "vpc-01", 
					"ReqInfo": { 
						"Name": "subnet-02", 
						"IPv4_CIDR": "192.168.2.0/24"
					} 
				}`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list vpc",
			CmdArgs:             []string{"vpc", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `{"vpc":[{"IId":{"NameId":"vpc-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get vpc",
			CmdArgs:             []string{"vpc", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "vpc-01"},
			ExpectResStartsWith: `{"IId":{"NameId":"vpc-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create security",
			CmdArgs: []string{"security", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConnectionName":"mock-unit-config01",
					"ReqInfo": { 
						"Name": "sg-01", 
						"VPCName": "vpc-01", 
						"SecurityRules": [ 
							{
								"FromPort": "1", 
								"ToPort" : "65535", 
								"IPProtocol" : "tcp", 
								"Direction" : "inbound",
								"CIDR" : "0.0.0.0/0"
							}
						] 
					} 
				}`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list security",
			CmdArgs:             []string{"security", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `{"securitygroup":[{"IId":{"NameId":"sg-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get security",
			CmdArgs:             []string{"security", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "sg-01"},
			ExpectResStartsWith: `{"IId":{"NameId":"sg-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "create keypair",
			CmdArgs: []string{"keypair", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConnectionName":"mock-unit-config01",
					"ReqInfo": { 
						"Name": "keypair-01" 
					} 
				}`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list keypair",
			CmdArgs:             []string{"keypair", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `{"keypair":[{"IId":{"NameId":"keypair-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get keypair",
			CmdArgs:             []string{"keypair", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "keypair-01"},
			ExpectResStartsWith: `{"IId":{"NameId":"keypair-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name: "start vm",
			CmdArgs: []string{"vm", "start", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConnectionName":"mock-unit-config01",
					"ReqInfo": { 
						"Name": "vm-01", 
						"ImageName": "mock-vmimage-01", 
						"VPCName": "vpc-01", 
						"SubnetName": "subnet-01", 
						"SecurityGroupNames": [ "sg-01" ], 
						"VMSpecName": "mock-vmspec-01", 
						"KeyPairName": "keypair-01"
					}
				}`,
			},
			ExpectResStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list vm",
			CmdArgs:             []string{"vm", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `{"vm":[{"IId":{"NameId":"vm-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get vm",
			CmdArgs:             []string{"vm", "get", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "vm-01"},
			ExpectResStartsWith: `{"IId":{"NameId":"vm-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "list vm status",
			CmdArgs:             []string{"vm", "liststatus", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01"},
			ExpectResStartsWith: `{"vmstatus":[{"IId":{"NameId":"vm-01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "get vm status",
			CmdArgs:             []string{"vm", "getstatus", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "vm-01"},
			ExpectResStartsWith: `{"Status":"Running"}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "terminate vm",
			CmdArgs:             []string{"vm", "terminate", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "vm-01", "--force", "true"},
			ExpectResStartsWith: `{"Status":"Terminating"}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete keypair",
			CmdArgs:             []string{"keypair", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "keypair-01", "--force", "true"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete security",
			CmdArgs:             []string{"security", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "sg-01", "--force", "true"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "remove subnet",
			CmdArgs:             []string{"vpc", "remove-subnet", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "--vname", "vpc-01", "--sname", "subnet-02", "--force", "true"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete vpc",
			CmdArgs:             []string{"vpc", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--cname", "mock-unit-config01", "-n", "vpc-01", "--force", "true"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "delete connection config",
			CmdArgs:             []string{"connection", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-config01"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "unregister region",
			CmdArgs:             []string{"region", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-region01"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "unregister credential",
			CmdArgs:             []string{"credential", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-credential01"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			Name:                "unregister cloud driver",
			CmdArgs:             []string{"driver", "delete", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "--name", "mock-unit-driver01"},
			ExpectResStartsWith: `{"Result":true}`,
		}
		SpiderCmdTest(t, tc)

		TearDownForCli()
	})
}
