package restscenario

import (
	"os"

	sshrun "github.com/cloud-barista/poc-cicd-spider/cloud-control-manager/vm-ssh"

	"bou.ke/monkey"
	"github.com/sirupsen/logrus"
)

type TestCases struct {
	name                 string
	echoFunc             string
	httpMethod           string
	whenURL              string
	givenQueryParams     string
	givenParaNames       []string
	givenParaVals        []string
	givenPostData        string
	expectStatus         int
	expectBodyStartsWith string
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
