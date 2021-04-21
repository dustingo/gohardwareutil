/*
系统的基础信息；包括系统、
*/
package gohardwareutil

import (
	"bytes"
	"runtime"
	"golang.org/x/sys/unix"
)
type uname struct {
	SysName    string
	Release    string
	Version    string
	Machine    string
	NodeName   string
	DomainName string
}
func Platform() string{
	return runtime.GOOS
}
/*
uname -a
系统名；kernel版本；发行版本；系统架构；hostname；domainname；
 */
func Uname()  *uname {
	var utsname unix.Utsname
	if err := unix.Uname(&utsname); err !=nil{
		panic(err)
	}
	output := uname{
		SysName:    string(utsname.Sysname[:bytes.IndexByte(utsname.Sysname[:], 0)]),
		Release:    string(utsname.Release[:bytes.IndexByte(utsname.Release[:], 0)]),
		Version:    string(utsname.Version[:bytes.IndexByte(utsname.Version[:], 0)]),
		Machine:    string(utsname.Machine[:bytes.IndexByte(utsname.Machine[:], 0)]),
		NodeName:   string(utsname.Nodename[:bytes.IndexByte(utsname.Nodename[:], 0)]),
		DomainName: string(utsname.Domainname[:bytes.IndexByte(utsname.Domainname[:], 0)]),
	}
	return &output
}