Work repository for GSOC 2016: Project - MIG Audit Module
=========================

The project is to implement a module for [MIG](http://mig.mozilla.org/) that support syscall monitoring via Audit Framework present in Linux kernel.

The module leverages [libaudit-go](https://github.com/mozilla/libaudit-go) a Go library that provides API to deal with Audit Framework and corelates messages coming from kernel into a single JSON message in a format identical to [audisp-json](https://github.com/gdestuynder/audisp-json/blob/master/messages_format.rst) and send them to web server that accepts POST requests.

NOTE: Module is only tested on amd64 architecture as libaudit-go currently only supports amd64.

## Work done during the project period:

* 	Addition of the mig-audit module, [Pull request](https://github.com/mozilla/mig/pull/253)

*	Fixing and testing the libaudit-go library, [Commits](https://github.com/mozilla/libaudit-go/commits?author=arunk-s)

*	Comparing performance of audisp-json and the Go module

	For the tests both audisp-json and go module were made blocking in nature therefore tests show high CPU usage.
	
	But for every day uses both are non-blocking in nature.

	The performance tests were done by creating a [test](vendor/mig.ninja/mig/modules/audit/audit_test.go) for the go module that takes a single file consisting of audit messages and parses, corelates and sends events to a local web server.
	
	Similarly the same file is feeded to audisp-json that also parses, corelates and sends events to the same local web server.
	
	Results:
	
	http://paste.ubuntu.com/23058381/ (with reverseMap), pprof: http://imgh.us/pprof001_5.svg 
	
	http://paste.ubuntu.com/23058390/ (without reverseMap), pprof: http://imgh.us/pprof001_6.svg

## Instructions for running the module:

First, you should have a web server running that accepts POST requests. You can get ephemeral urls that allows POST on websites like [requesb.in](http://requestb.in/).

Modify [params](main.go#L17) in the module to specify the address for the server.

Module also writes raw audit messages on `/tmp/log` so it can be checked as well to see that events are coming.
You can also use `auditctl -l` and `auditctl -s` to further check that audit settings are applied correctly.

Suggested Environment: `Ubuntu 16.04 64bit , auditd version 2.4.5`

To run the module(require super user access):
```go
go build main.go

sudo ./main
```
