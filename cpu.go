/*
CPU信息
CpuUsage为CPU信息结构体，其中包含了 CPU的各个状态下的时间
CpuInfo函数返回cpu空闲时间百分比，和使用百分比；
 */
package gohardwareutil
import (
	"bufio"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
	"time"
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
func getCpuInfo() *CpuUsage{
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
// CpuInfo 返回cpu使用率和cpu空闲率
func CpuInfo() (string,string){
	//cpu0 := CpuUsage{}
	//cpu1 := CpuUsage{}
	//file := "/proc/stat"
	//f,err := os.Open(file)
	//if err != nil{
	//	panic(err)
	//}
	//defer f.Close()
	//fScanner := bufio.NewScanner(f)
	//for fScanner.Scan(){
	//	firstLine := fScanner.Text()
	//	parts := strings.Fields(firstLine)
	//	user,_:=strconv.ParseUint(parts[1],10,64)
	//	nice,_:=strconv.ParseUint(parts[2],10,64)
	//	system,_:=strconv.ParseUint(parts[3],10,64)
	//	idle,_:=strconv.ParseUint(parts[4],10,64)
	//	iowait,_:=strconv.ParseUint(parts[5],10,64)
	//	irq,_:=strconv.ParseUint(parts[6],10,64)
	//	softirq,_:=strconv.ParseUint(parts[7],10,64)
	//	steal,_:=strconv.ParseUint(parts[8],10,64)
	//	guest,_:=strconv.ParseUint(parts[9],10,64)
	//	guestnice,_:=strconv.ParseUint(parts[10],10,64)
	//	cpu0.CPU  = parts[0]
	//	cpu0.USER = user
	//	cpu0.NICE = nice
	//	cpu0.SYSTEM = system
	//	cpu0.IDLE = idle
	//	cpu0.IOWAIT = iowait
	//	cpu0.IRQ = irq
	//	cpu0.SOFTIRQ = softirq
	//	cpu0.STEAL = steal
	//	cpu0.GUEST = guest
	//	cpu0.GUESTNICE = guestnice
	//	break
	//}
	//return &cpu0,&cpu1
	cpu0 := getCpuInfo()
	time.Sleep(1 * time.Second)
	cpu1 := getCpuInfo()
	idle := cpu1.IDLE - cpu0.IDLE
	total := (cpu1.USER+cpu1.NICE+cpu1.SYSTEM+cpu1.IDLE+cpu1.IOWAIT+cpu1.IRQ+cpu1.SOFTIRQ+cpu1.STEAL) - (cpu0.USER+cpu0.NICE+cpu0.SYSTEM+cpu0.IDLE+cpu0.IOWAIT+cpu0.IRQ+cpu0.SOFTIRQ+cpu0.STEAL)
	cpuidle := decimal.NewFromFloatWithExponent(float64(idle)/float64(total)*100,-2)
	cpuused := decimal.NewFromFloatWithExponent(100 - float64(idle)/float64(total)*100,-2)
	return cpuidle.String(),cpuused.String()
}