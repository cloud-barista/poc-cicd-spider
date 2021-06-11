package restscenario

import (
	"os"

	sshrun "github.com/cloud-barista/poc-cicd-spider/cloud-control-manager/vm-ssh"

	"bou.ke/monkey"
	"github.com/sirupsen/logrus"
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
}

var (
	holdStdout *os.File = nil
)

func init() {
	logrus.SetLevel(logrus.ErrorLevel)
}

func SetUpForRest() {

	holdStdout = os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull)

	os.RemoveAll("../meta_db")

	monkey.Patch(sshrun.SSHRun, func(sshInfo sshrun.SSHInfo, cmd string) (string, error) {
		return cmd + " success", nil
	})
}

func TearDownForRest() {
	os.RemoveAll("../meta_db")

	os.Stdout = holdStdout
}
