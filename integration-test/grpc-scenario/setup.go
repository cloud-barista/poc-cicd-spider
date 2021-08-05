package grpcscenario

import (
	"context"
	"io"
	"net"
	"os"
	"time"

	"github.com/cloud-barista/poc-cicd-spider/api-runtime/grpc-runtime/config"
	"github.com/cloud-barista/poc-cicd-spider/api-runtime/grpc-runtime/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	cbstore "github.com/cloud-barista/cb-store"
	gc "github.com/cloud-barista/poc-cicd-spider/api-runtime/grpc-runtime/common"
	grpc_service "github.com/cloud-barista/poc-cicd-spider/api-runtime/grpc-runtime/service"
	pb "github.com/cloud-barista/poc-cicd-spider/api-runtime/grpc-runtime/stub/cbspider"
	sshrun "github.com/cloud-barista/poc-cicd-spider/cloud-control-manager/vm-ssh"

	"github.com/cloud-barista/poc-cicd-spider/interface/api"

	"bou.ke/monkey"
	"github.com/sirupsen/logrus"
)

type TestCases struct {
	Name                string
	Instance            interface{}
	Method              string
	Args                []interface{}
	ExpectResStartsWith string
	ExpectResContains   string
}

var (
	holdStdout *os.File     = nil
	CimApi     *api.CIMApi  = nil
	CcmApi     *api.CCMApi  = nil
	gs         *grpc.Server = nil
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}

func SetUpForGrpc() {

	holdStdout = os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull)

	cbstore.InitData()

	/**
	** Spider Grpc Server Setup
	**/
	listener := bufconn.Listen(1024 * 1024)

	monkey.Patch(gc.NewCBConnection, func(gConf *config.GrpcClientConfig) (*gc.CBConnection, io.Closer, error) {
		conn, _ := grpc.DialContext(context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(
			func(context.Context, string) (net.Conn, error) {
				return listener.Dial()
			}))
		return &gc.CBConnection{Conn: conn}, nil, nil
	})

	logger := logger.NewLogger()

	spidersrv := &config.GrpcServerConfig{
		Addr: "127.0.0.1:32048",
	}

	cbserver, closer, err := gc.NewCBServer(spidersrv)
	if err != nil {
		logger.Fatal("failed to create grpc server: ", err)
	}

	gs = cbserver.Server
	pb.RegisterCIMServer(gs, &grpc_service.CIMService{})
	pb.RegisterCCMServer(gs, &grpc_service.CCMService{})
	pb.RegisterSSHServer(gs, &grpc_service.SSHService{})

	go func() {

		if closer != nil {
			defer closer.Close()
		}

		if err := gs.Serve(listener); err != nil {
			logger.Fatal("failed to serve: ", err)
		}
	}()

	time.Sleep(time.Second * 2)

	/**
	** Spider Grpc API Setup
	**/
	CimApi = api.NewCloudInfoManager()

	err = CimApi.SetConfigPath("../conf/grpc_conf.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = CimApi.Open()
	if err != nil {
		logger.Fatal(err)
	}

	CcmApi = api.NewCloudResourceHandler()

	err = CcmApi.SetConfigPath("../conf/grpc_conf.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	err = CcmApi.Open()
	if err != nil {
		logger.Fatal(err)
	}

	/**
	** sshrun.SSHRun Patch
	**/
	monkey.Patch(sshrun.SSHRun, func(sshInfo sshrun.SSHInfo, cmd string) (string, error) {
		return cmd + " success", nil
	})
}

func TearDownForGrpc() {
	CimApi.Close()
	CcmApi.Close()
	gs.Stop()

	cbstore.InitData()

	os.Stdout = holdStdout
}
