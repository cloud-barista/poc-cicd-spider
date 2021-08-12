package restscenario

import (
	"os"

	cbstore "github.com/cloud-barista/cb-store"
	sshrun "github.com/cloud-barista/poc-cicd-spider/cloud-control-manager/vm-ssh"
	"github.com/sirupsen/logrus"

	"bou.ke/monkey"
)

type TestCases struct {
	Name                 string
	EchoFunc             string
	HttpMethod           string
	WhenURL              string
	GivenQueryParams     string
	GivenParaNames       []string
	GivenParaVals        []string
	GivenPostData        string
	ExpectStatus         int
	ExpectBodyStartsWith string
	ExpectBodyContains   string
}

var (
	holdStdout *os.File = nil
	nullOut    *os.File = nil
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}

func SetUpForRest() {

	holdStdout = os.Stdout
	nullOut, _ := os.Open(os.DevNull)
	os.Stdout = nullOut

	cbstore.InitData()

	monkey.Patch(sshrun.SSHRun, func(sshInfo sshrun.SSHInfo, cmd string) (string, error) {
		return cmd + " success", nil
	})
}

func TearDownForRest() {

	cbstore.InitData()

	nullOut.Close()
	os.Stdout = holdStdout
}
