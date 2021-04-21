/*
/proc/meminfo下所有的信息
MemInfo()返回一个字典，包含/proc/meminfo里的字段信息,可自己解析字典计算；也可使用MemTotal、MemFree、MemAvailable、MemCached、MemUsed等函数直接获取结果
*/
package gohardwareutil

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
)
var memInfo  = map[string]float64{}
// 获取内存
func MemInfo()(map[string]float64,error){
	file := "/proc/meminfo"
	f , err := os.Open(file)
	if err != nil{
		//
		panic(err)
	}
	defer f.Close()
	memScanner := bufio.NewScanner(f)
	for memScanner.Scan(){
		line := memScanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0{
			continue
		}
		floatVal, err := strconv.ParseFloat(parts[1],64)
		if err != nil{
			//return nil,fmt.Errorf("invalid value in meminfo:%w",err)
			panic(err)
		}
		// key 去除第一个字段的 : 冒号，当作key
		key := parts[0][:len(parts[0])-1]
		memInfo[key] = floatVal
	}
	return memInfo,nil
}
// MemTotal 总内存大小
func MemTotal() float64{
	return memInfo["MemTotal"]
}
//MemFree 空闲内存大小
func MemFree()float64{
	return memInfo["MemFree"]
}
// MemAvailable 可用内存大小
func MemAvailable()float64{
	return memInfo["MemAvailable"]
}
//Buffers 大小
func MemBuffers()float64{
	return memInfo["Buffers"]
}
// Cached 大小
func MemCached()float64{
	return memInfo["Cached"]
}
// 内存使用率
func MemUsed() string{
	return fmt.Sprintf("%s%%",decimal.NewFromFloatWithExponent((memInfo["MemTotal"] - memInfo["MemFree"] - memInfo["Buffers"] - memInfo["Cached"])/memInfo["MemTotal"]*100,-2))
}