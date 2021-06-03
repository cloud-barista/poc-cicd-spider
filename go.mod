module github.com/cloud-barista/poc-cicd-spider

go 1.16

replace (
	google.golang.org/api => google.golang.org/api v0.15.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0

)

require (
	bou.ke/monkey v1.0.2 // indirect
	github.com/Azure/azure-sdk-for-go v51.2.0+incompatible
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.7
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/Microsoft/go-winio v0.4.16 // indirect

	// Do not delete. These are packages used at building NCP plugin Driver in Container.
	github.com/NaverCloudPlatform/ncloud-sdk-go-v2 v1.1.7
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.920
	github.com/appleboy/easyssh-proxy v1.3.9
	github.com/aws/aws-sdk-go v1.37.10
	github.com/bramvdbogaerde/go-scp v0.0.0-20201229172121-7a6c0268fa67
	github.com/chyeh/pubip v0.0.0-20170203095919-b7e679cf541c
	github.com/cloud-barista/cb-log v0.3.1
	github.com/cloud-barista/cb-store v0.3.14
	github.com/containerd/containerd v1.4.3 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.0.0-20200309214505-aa6a9891b09c
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-units v0.4.0 // indirect
	github.com/go-resty/resty/v2 v2.6.0 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/gophercloud/gophercloud v0.17.0
	github.com/gotestyourself/gotestyourself v1.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/labstack/echo/v4 v4.2.0
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/sftp v1.13.0 // indirect
	github.com/prometheus/client_golang v1.9.0
	github.com/sfreiberg/simplessh v0.0.0-20180301191542-495cbb862a9c
	github.com/sirupsen/logrus v1.8.1
	github.com/snowzach/rotatefilehook v0.0.0-20180327172521-2f64f265f58c
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go v1.0.150
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	golang.org/x/oauth2 v0.0.0-20210210192628-66670185b0cd
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	google.golang.org/api v0.40.0
	google.golang.org/genproto v0.0.0-20210510173355-fb37daa5cd7a // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gotest.tools v1.4.0 // indirect
// Do not delete. These are packages used at building NCP plugin Driver in Container.

)

retract (
	v0.3.12
	v0.3.11
)
