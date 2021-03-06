﻿# gcd [![Build status](https://ci.appveyor.com/api/projects/status/n3oqlhfscxyuhvcq?svg=true)](https://ci.appveyor.com/project/tischda/gcd)

Windows utility written in [Go](https://www.golang.org) to serve arguments to a "cd" command alias.

In particular this switch:

	Use the /D switch to change current drive in addition to changing current
	directory for a drive.

I want it to be set automatically.


### Install

~~~
go install github.com/tischda/gcd@latest
~~~

### Usage

~~~
gcd <path>
~~~

Example:

~~~
c:\> gcd.exe c:\temp
c:\temp

c:\> gcd.exe e:\temp
/d e:\temp

c:\> gcd.exe c:\program files
c:\program files
~~~

To actually change the working directory, define an alias to the cd command:

~~~
ncd=FOR /F "tokens=1 delims='" %G IN ('gcd $*') do @cd %G
~~~

(I am using [cmder](https://cmder.net/))


### Known issues

This fails:

~~~
c:\> gcd.exe tips & tricks
~~~

Workarounds:

~~~
c:\> gcd.exe "tips & tricks"
c:\> gcd.exe tips ^& tricks
c:\> gcd.exe tips ? tricks
~~~

### References

* https://stackoverflow.com/questions/17026290/golang-chdir-and-stay-there-on-program-termination
* https://stackoverflow.com/questions/52435908/how-to-change-the-shells-current-working-directory-in-go
* https://stackoverflow.com/questions/53984853/change-parent-shell-directory-from-go
