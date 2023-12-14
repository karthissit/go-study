# go-study


Start with :
===========
The starting point is the main() function. Triggerd from proc.main ==> a process

go mod init example.com 

go get golang.org/x/exp/constraints 

go.work -> file
===============
to work with the local modules internally within the project

go.sum:
=======
Auto-generated while receiving the external modules


Any new module creation within project :
=======================================

use: go mod init module_name
and update go.work file


Creating build file for other OS and Platforms:
===============================================

The GOARCH environment variable in Go specifies the target architecture for compilation. The common values include:

amd64: 64-bit x86 architecture, also known as x86-64 or Intel 64.
386: 32-bit x86 architecture.
arm: ARM architecture (32-bit).
arm64: ARM 64-bit architecture.
ppc64: PowerPC 64-bit architecture.
ppc64le: Little-endian PowerPC 64-bit architecture.
mips: MIPS architecture.
mipsle: Little-endian MIPS architecture.
mips64: 64-bit MIPS architecture.
mips64le: Little-endian 64-bit MIPS architecture.


eWhen cross-compiling, you set the GOARCH variable to the target architecture. For example:

"export GOARCH=arm64  # Set the target architecture to ARM 64-bit"

It's important to note that not all combinations of GOOS and GOARCH are supported for cross-compilation. The supported combinations depend on the Go compiler and the libraries available for each platform.

You can check the supported combinations by running:
   "go tool dist list"
This will display a list of supported combinations for GOOS and GOARCH in your Go installation. Always refer to the official Go documentation for the most up-to-date information on supported architectures.


The GOOS environment variable in Go specifies the target operating system for compilation. Here are some common values that can be passed to GOOS:

darwin: macOS
linux: Linux
windows: Windows
freebsd: FreeBSD
openbsd: OpenBSD
netbsd: NetBSD
dragonfly: DragonFly BSD
solaris: Solaris
plan9: Plan 9 from User Space
nacl: Native Client (deprecated)
android: Android
For example, if you want to cross-compile your Go program for Windows, you would set GOOS to "windows":

"export GOOS=windows"
Similarly, if you want to target Linux, you would set it to "linux":

"export GOOS=linux"

When cross-compiling, you typically set both GOOS and GOARCH to the target operating system and architecture. For instance, for Windows on a 64-bit architecture:

"export GOOS=windows
export GOARCH=amd64"


You can check the supported combinations of GOOS and GOARCH for your Go installation by running:

"go tool dist list"

This command will display a list of supported combinations for GOOS and GOARCH. Always refer to the official Go documentation for the most up-to-date information on supported operating systems and architectures.



example: I am trying to build the software from linux to windows
-----------------------------------------------------------------
root@karthn go-study]# export GOOS=windows

root@karthn go-study]# export GOARCH=amd64

root@karthn go-study]# go build -o gostudy.exe  ==> This will create an executable that could be run on windows machine with x64 architecture



If you want to reset the target GOOS and GOARCH:
-----------------------------------------------
unset GOOS
unset GOARCH