# What's go?

Go is open source programming language that makes it easy to build simple, reliable and efficient software.

## Prerequisites

```
$ wget [go-packege](https://golang.org/dl/)
$ tar -xvf go1.10.2.linux-amd64.tar.gz
$ mv go /usr/local
$ export GOROOT=/usr/local/go
$ export GOPATH=$HOME/Apps/app1
$ export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

All above environment will be set for your current session only. To make it permanent add above commands in ~/.profile file.

## Building for the current system (Linux 64 bit)

```
$ go build hello.go
$ file hello
hello: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped
```

## Build for Windows 

```
$ GOOS=windows go build hello.go
$ file hello.exe
hello.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
```

## Platform specific code

main.go
  out, err := exec.Command("ifconfig","-a").Output

**Refactor the command line part**

main.go
  cmd := GetCommand()
  out, err := exec.Command(cmd[0],cmd[1:]...).Output

**Lunux**

main_linux.go

  func GetCommand() []string {
    return []string{"ifconfig","-a"}
  }

**Windows**

main_windows.go

  func GetCommand() []string {
    return []string {"ipconfig","/all"
  }

## Two mechanisms

* Naming the file with the build restriction
```
foo_${GOOS}.go
```
* Adding build directives to the top of the line
```
// +build linux, 386 darwin
```

# And remember folks you're now writing services, not scripts




