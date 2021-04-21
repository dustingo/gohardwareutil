/*
CPU信息
CpuUsage为CPU信息结构体，其中包含了 CPU的各个状态下的时间
CpuUsage有三个方法：CpuIdle,CpuTotal,CpuUsed，分别返回对应的值，
usage:
cpu := CpuInfo()
cpu.CpuIdle()
cpu.CpuTotal()
cpu.CpuUsed()
cpu.IDLE
and so on...
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
// /proc/stat中第一行，总的CPU信息
type CpuUsage struct {
	CPU string
	USER uint64
	NICE uint64
	SYSTEM uint64
	IDLE uint64
	IOWAIT uint64
	IRQ uint64
	SOFTIRQ uint64
	STEAL uint64
	GUEST uint64
	GUESTNICE uint64

}
// CpuInfo 返回cpu状态信息结构体指针，需要计算其他指标可以自行计算；计算时无论是uint 还是float 需要注意精度问题奥
func CpuInfo() *CpuUsage{
	cpu := CpuUsage{}
	file := "/proc/stat"
	f,err := os.Open(file)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	fScanner := bufio.NewScanner(f)
	for fScanner.Scan(){
		firstLine := fScanner.Text()
		parts := strings.Fields(firstLine)
		user,_:=strconv.ParseUint(parts[1],10,64)
		nice,_:=strconv.ParseUint(parts[2],10,64)
		system,_:=strconv.ParseUint(parts[3],10,64)
		idle,_:=strconv.ParseUint(parts[4],10,64)
		iowait,_:=strconv.ParseUint(parts[5],10,64)
		irq,_:=strconv.ParseUint(parts[6],10,64)
		softirq,_:=strconv.ParseUint(parts[7],10,64)
		steal,_:=strconv.ParseUint(parts[8],10,64)
		guest,_:=strconv.ParseUint(parts[9],10,64)
		guestnice,_:=strconv.ParseUint(parts[10],10,64)
		cpu.CPU  = parts[0]
		cpu.USER = user
		cpu.NICE = nice
		cpu.SYSTEM = system
		cpu.IDLE = idle
		cpu.IOWAIT = iowait
		cpu.IRQ = irq
		cpu.SOFTIRQ = softirq
		cpu.STEAL = steal
		cpu.GUEST = guest
		cpu.GUESTNICE = guestnice
		break
	}
	return &cpu
}
// 返回CPU IDLE
func (c *CpuUsage) CpuIdle() uint64{
	return c.IDLE
}
// 返回总CPU总时间
func(c *CpuUsage)CpuTotal()uint64{
	return (c.USER+c.NICE+c.SYSTEM+c.IDLE+c.IOWAIT+c.IRQ+c.SOFTIRQ+c.STEAL)
}
//返回CPU使用率
func(c *CpuUsage)CpuUsed(s string) string {
	//usedCpu := float64(c.IDLE)/float64(c.USER+c.NICE+c.SYSTEM+c.IDLE+c.IOWAIT+c.IRQ+c.SOFTIRQ+c.STEAL) * 100
	if s == "percentage"{
		return fmt.Sprintf("%v%%",decimal.NewFromFloatWithExponent(100 - float64(c.IDLE)/float64(c.USER+c.NICE+c.SYSTEM+c.IDLE+c.IOWAIT+c.IRQ+c.SOFTIRQ+c.STEAL) * 100,-2).String())
	}else{
		return fmt.Sprintf("%v",decimal.NewFromFloatWithExponent(100 - float64(c.IDLE)/float64(c.USER+c.NICE+c.SYSTEM+c.IDLE+c.IOWAIT+c.IRQ+c.SOFTIRQ+c.STEAL) * 100,-2).String())
	}

}