#### gohardwareutil
> STATEMENT
>
Go语言业余学习者，不能保证代码的高效性；适用于linux；GO Version >= 1.7
> INFORMATION
>
- uname 信息 
- CPU信息 （/proc/cpuinfo）
- MEM信息  (/proc/meminfo) 
- DISK磁盘信息 
- 详细想看具体代码注释说明
> 示例
>
```go
package main
import (
        "github.com/dustingo/gohardwareutil"
        "fmt"
)
func main(){
        //CPU cpu:CpuUsage struct
        cpu := gohardwareutil.CpuInfo()
        cpuUsedPer := cpu.CpuUsed()
        fmt.Printf("cpuUsed:%s\n",cpuUsedPer)
        //uname uname:uname struct
        uname := gohardwareutil.Uname()
        fmt.Printf("nodeName:%s\n",uname.NodeName)
        //mem mem:map[string]float64;可自行取值计算
        mem,err :=gohardwareutil.MemInfo()
        if err != nil{
                panic(err)
        }
        fmt.Println("PageTables:",mem["PageTables"])
        memUsed := gohardwareutil.MemUsed()
        fmt.Printf("memUsed:%s\n",memUsed)
        //disk
        diskstats ,err := gohardwareutil.GetDiskStats()
        if err != nil{
                panic(err)
        }
        fmt.Println(diskstats)
        for  _,stats  := range diskstats{
                fmt.Println("MountPoint:",stats.Labels.MountPoint)
                fmt.Println("Size:",stats.Size)
                fmt.Println("Free:",stats.Free)
        }

}
```
