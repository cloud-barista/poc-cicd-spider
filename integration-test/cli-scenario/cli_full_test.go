package cliscenario

import (
	"testing"
)

func TestCliFull(t *testing.T) {
	t.Run("command full test for mock driver", func(t *testing.T) {
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
			Name: "register credential",
			CmdArgs: []string{"credential", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"CredentialName":"mock-unit-credential01",
					"ProviderName":"MOCK",
					"KeyValueInfoList": [
						{"Key":"MockName", "Value":"mock_unit_name00"}
					]
				}`,
			},
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

		TearDownForCli()
	})
}
