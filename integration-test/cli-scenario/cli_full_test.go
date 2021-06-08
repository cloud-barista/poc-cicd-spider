package cliscenario

import (
	"testing"
)

func TestCliFull(t *testing.T) {
	t.Run("command full test for mock driver", func(t *testing.T) {
		setUpForCli()

		tc := TestCases{
			name:                "list cloud os",
			cmdArgs:             []string{"os", "list", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json"},
			expectResStartsWith: `{"cloudos":[`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			name: "register cloud driver",
			cmdArgs: []string{"driver", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"DriverName":"mock-unit-driver01",
					"ProviderName":"MOCK", 
					"DriverLibFileName":"mock-driver-v1.0.so"
				}`,
			},
			expectResStartsWith: `{"DriverName":"mock-unit-driver01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			name: "register credential",
			cmdArgs: []string{"credential", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"CredentialName":"mock-unit-credential01",
					"ProviderName":"MOCK",
					"KeyValueInfoList": [
						{"Key":"MockName", "Value":"mock_unit_name00"}
					]
				}`,
			},
			expectResStartsWith: `{"CredentialName":"mock-unit-credential01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			name: "register region",
			cmdArgs: []string{"region", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"RegionName":"mock-unit-region01",
					"ProviderName":"MOCK",
					"KeyValueInfoList": [
						{"Key":"Region", "Value":"default"}
					]
				}`,
			},
			expectResStartsWith: `{"RegionName":"mock-unit-region01"`,
		}
		SpiderCmdTest(t, tc)

		tc = TestCases{
			name: "create connection config",
			cmdArgs: []string{"connection", "create", "--config", "../conf/grpc_conf.yaml", "-i", "json", "-o", "json", "-d",
				`{
					"ConfigName":"mock-unit-config01",
					"ProviderName":"MOCK", 
					"DriverName":"mock-unit-driver01", 
					"CredentialName":"mock-unit-credential01", 
					"RegionName":"mock-unit-region01"
				}`,
			},
			expectResStartsWith: `{"ConfigName":"mock-unit-config01"`,
		}
		SpiderCmdTest(t, tc)

		tearDownForCli()
	})
}
