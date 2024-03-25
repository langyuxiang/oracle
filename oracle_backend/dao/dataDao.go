package dao

import (
	"Oracle/model"
	"Oracle/sdk"
	"errors"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"math"
	"math/big"
	"os/exec"
	"strconv"
)

func GradeData(nodeId int64, dataImportance float32, NodeFailureRate float32, Source float32, integrity float32, Utilization float32, result int64) (bool, error) {
	data := model.Data{}
	DB.Table("data").Where("data_source_id = ?", nodeId).First(&data)

	configs, err := conf.ParseConfigFile("./config.toml")
	if err != nil {
		logrus.Fatal(err)
		return false, err
	}
	config := &configs[0]

	client, err := client.Dial(config)
	if err != nil {
		logrus.Fatal(err)
	}
	contractAddress := common.HexToAddress("0x2bc7b60829231f8a8018ffee5f246ce9439c249f")
	oracle, err := sdk.NewOracle(contractAddress, client)
	if err != nil {
		logrus.Error(err)
		return false, err
	}

	nodeIdBi := big.NewInt(nodeId)
	dataImportancei := int64(math.Ceil(float64(dataImportance)))
	dataImportanceBi := big.NewInt(dataImportancei)
	NodeFailureRatei := int64(math.Ceil(float64(NodeFailureRate)))
	NodeFailureRateiBi := big.NewInt(NodeFailureRatei)
	Sourcei := int64(math.Ceil(float64(Source)))
	SourceBi := big.NewInt(Sourcei)
	integrityi := int64(math.Ceil(float64(integrity)))
	integrityBi := big.NewInt(integrityi)
	Utilizationi := int64(math.Ceil(float64(Utilization)))
	UtilizationBi := big.NewInt(Utilizationi)
	oracle.AddNode(client.GetTransactOpts(), big.NewInt(nodeId))
	_, _, err = oracle.Grade(client.GetTransactOpts(), nodeIdBi, dataImportanceBi, NodeFailureRateiBi, SourceBi, integrityBi, UtilizationBi, big.NewInt(result))
	if err != nil {
		logrus.Error(err)
		return false, errors.New("合约评估失败")
	}

	res, _ := oracle.GetResult(client.GetCallOpts(), nodeIdBi)
	if res == false {
		return false, errors.New("评估为不可信")
	}

	dataImportanceStr := strconv.FormatFloat(float64(dataImportance), 'f', -2, 32)
	NodeFailureRateStr := strconv.FormatFloat(float64(NodeFailureRate), 'f', -2, 32)
	SourceStr := strconv.FormatFloat(float64(Source), 'f', -2, 32)
	integrityStr := strconv.FormatFloat(float64(integrity), 'f', -2, 32)
	UtilizationStr := strconv.FormatFloat(float64(Utilization), 'f', -2, 32)
	resList := fmt.Sprintf("[%v,%v,%v,%v,%v]", dataImportanceStr, NodeFailureRateStr, SourceStr, integrityStr, UtilizationStr)
	//cmd := exec.Command("python3", "model_record/prediction_model_use.py", resList)

	//python.Initialize()
	//pdu := ImportModule("/usr/lib/go-1.18/src/Oracle/model_record", "prediction_model_use")
	//bArgs := python.PyTuple_New(1)
	//python.PyTuple_SetItem(bArgs, 0, PyStr(resList))
	//pyResult := pdu.Call(bArgs, python.Py_None)
	//fmt.Println(GoStr(pyResult))
	cmd := exec.Command("python3", "model_record/prediction_model_use.py", resList)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	//pyfunc := fmt.Sprintf("prediction(*%v)", resList)
	//cmd := exec.Command("python3", "-c", pyfunc)
	//output, err := cmd.CombinedOutput()
	//fmt.Println(string(output))
	if err != nil {
		fmt.Println(err)
		return false, errors.New("预测数据失败")
	}

	pyResult := string(output)
	if pyResult == "0" {
		oracle.DeleteInformation(&bind.TransactOpts{}, nodeIdBi)
		return false, errors.New("合约删除数据")
	} else {
		data.IsBelievable = true
		data.DataImportance = dataImportance
		data.NodeFailureRate = NodeFailureRate
		data.SourceChannel = Source
		data.Integrity = integrity
		data.UsageRate = Utilization
	}

	DB.Table("data").Where("data_source_id = ?", nodeId).Updates(&data)

	return true, nil
}

//var GoStr = python.PyString_AS_STRING
//var PyStr = python.PyString_FromString
//
//func ImportModule(dir, name string) *python.PyObject {
//	sysModule := python.PyImport_ImportModule("sys") // import sys
//	path := sysModule.GetAttrString("path")          // path = sys.path
//	python.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
//	return python.PyImport_ImportModule(name)        // return __import__(name)
//
//}

func FindDataById(dataId int) model.Data {
	data := model.Data{}
	err := DB.Table("data").Where("data_source_id = ?", dataId).First(&data).Error
	if err != nil {
		logrus.Error(err)
	}

	return data
}

func UseData(dataSourceId int, DurationTime string) error {
	data := model.Data{}
	err := DB.Table("data").Where("data_source_id = ?", dataSourceId).First(&data).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	dt, err := strconv.ParseFloat(DurationTime, 64)
	newLevel := math.Log10(dt)
	if int(newLevel) > data.BelievableLevel && int(newLevel) <= 10 {
		data.BelievableLevel = int(newLevel)
		err = DB.Table("data").Where("data_source_id = ?", dataSourceId).Updates(&data).Error
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else if int(newLevel) >= 10 {
		data.BelievableLevel = 10
		err = DB.Table("data").Where("data_source_id = ?", dataSourceId).Updates(&data).Error
		if err != nil {
			logrus.Error(err)
			return err
		}
	}

	return nil

}

func FeedBack(dataSourceId int) error {

	data := model.Data{}
	err := DB.Table("data").Where("data_source_id = ?", dataSourceId).First(&data).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	//lnx = ln2 * log2X
	data.BelievableLevel = -data.BelievableLevel
	data.DoEvilNum++
	needWorkTime := math.Log2(-float64(data.BelievableLevel)*float64(data.DoEvilNum)+1) * (24 * 60 * 60) * math.Ln2
	data.NeedWorkTime = strconv.FormatFloat(needWorkTime, 'f', 2, 64) + "s"
	err = DB.Table("data").Where("data_source_id = ?", dataSourceId).Updates(&data).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
