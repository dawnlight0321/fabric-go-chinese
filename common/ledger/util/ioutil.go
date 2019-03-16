
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455961565728768>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package util

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/pkg/errors"
)

var logger = flogging.MustGetLogger("kvledger.util")

//
func CreateDirIfMissing(dirPath string) (bool, error) {
//
	if !strings.HasSuffix(dirPath, "/") {
		dirPath = dirPath + "/"
	}
	logger.Debugf("CreateDirIfMissing [%s]", dirPath)
	logDirStatus("Before creating dir", dirPath)
	err := os.MkdirAll(path.Dir(dirPath), 0755)
	if err != nil {
		logger.Debugf("Error creating dir [%s]", dirPath)
		return false, errors.Wrapf(err, "error creating dir [%s]", dirPath)
	}
	logDirStatus("After creating dir", dirPath)
	return DirEmpty(dirPath)
}

//如果dirpath处的dir为空，则dir empty返回true
func DirEmpty(dirPath string) (bool, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		logger.Debugf("Error opening dir [%s]: %+v", dirPath, err)
		return false, errors.Wrapf(err, "error opening dir [%s]", dirPath)
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	err = errors.Wrapf(err, "error checking if dir [%s] is empty", dirPath)
	return false, err
}

//
//
func FileExists(filePath string) (bool, int64, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, errors.Wrapf(err, "error checking if file [%s] exists", filePath)
	}
	return true, fileInfo.Size(), nil
}

//ListSubdirs返回子目录
func ListSubdirs(dirPath string) ([]string, error) {
	subdirs := []string{}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading dir %s", dirPath)
	}
	for _, f := range files {
		if f.IsDir() {
			subdirs = append(subdirs, f.Name())
		}
	}
	return subdirs, nil
}

func logDirStatus(msg string, dirPath string) {
	exists, _, err := FileExists(dirPath)
	if err != nil {
		logger.Errorf("Error checking for dir existence")
	}
	if exists {
		logger.Debugf("%s - [%s] exists", msg, dirPath)
	} else {
		logger.Debugf("%s - [%s] does not exist", msg, dirPath)
	}
}

