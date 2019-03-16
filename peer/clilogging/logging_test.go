
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456109544968192>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package clilogging

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/peer/common"
	common2 "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name      string
	args      []string
	shouldErr bool
}

func initLoggingTest(command string) (*cobra.Command, *LoggingCmdFactory) {
	adminClient := common.GetMockAdminClient(nil)
	mockCF := &LoggingCmdFactory{
		AdminClient: adminClient,
		wrapWithEnvelope: func(msg proto.Message) *common2.Envelope {
			pl := &common2.Payload{
				Data: utils.MarshalOrPanic(msg),
			}
			env := &common2.Envelope{
				Payload: utils.MarshalOrPanic(pl),
			}
			return env
		},
	}
	var cmd *cobra.Command
	if command == "getlevel" {
		cmd = getLevelCmd(mockCF)
	} else if command == "setlevel" {
		cmd = setLevelCmd(mockCF)
	} else if command == "revertlevels" {
		cmd = revertLevelsCmd(mockCF)
	} else if command == "getlogspec" {
		cmd = getLogSpecCmd(mockCF)
	} else if command == "setlogspec" {
		cmd = setLogSpecCmd(mockCF)
	} else {
//只有在下面的测试用例中出现拼写错误时才会发生
	}
	return cmd, mockCF
}

func runTests(t *testing.T, command string, tc []testCase) {
	cmd, _ := initLoggingTest(command)
	assert := assert.New(t)
	for i := 0; i < len(tc); i++ {
		t.Run(tc[i].name, func(t *testing.T) {
			cmd.SetArgs(tc[i].args)
			err := cmd.Execute()
			if tc[i].shouldErr {
				assert.NotNil(err)
			}
			if !tc[i].shouldErr {
				assert.Nil(err)
			}
		})
	}
}

//使用各种参数测试GetLevel
func TestGetLevel(t *testing.T) {
	var tc []testCase
	tc = append(tc,
		testCase{"NoParameters", []string{}, true},
		testCase{"Valid", []string{"peer"}, false},
	)
	runTests(t, "getlevel", tc)
}

//teststlevel用各种参数测试setlevel
func TestSetLevel(t *testing.T) {
	var tc []testCase
	tc = append(tc,
		testCase{"NoParameters", []string{}, true},
		testCase{"OneParameter", []string{"peer"}, true},
		testCase{"Valid", []string{"peer", "warning"}, false},
		testCase{"InvalidLevel", []string{"peer", "invalidlevel"}, true},
	)
	runTests(t, "setlevel", tc)
}

//testRevertLevels使用各种参数测试RevertLevels
func TestRevertLevels(t *testing.T) {
	var tc []testCase
	tc = append(tc,
		testCase{"Valid", []string{}, false},
		testCase{"ExtraParameter", []string{"peer"}, true},
	)
	runTests(t, "revertlevels", tc)
}

//testgetlogspec用各种参数测试getlogspec
func TestGetLogSpec(t *testing.T) {
	var tc []testCase
	tc = append(tc,
		testCase{"Valid", []string{}, false},
		testCase{"ExtraParameter", []string{"peer"}, true},
	)
	runTests(t, "getlogspec", tc)
}

//testsetlogspec用各种参数测试setlogspec
func TestSetLogSpec(t *testing.T) {
	var tc []testCase
	tc = append(tc,
		testCase{"NoParameters", []string{}, true},
		testCase{"Valid", []string{"debug"}, false},
	)
	runTests(t, "setlogspec", tc)
}

