/*
<!--
Copyright (c) 2019 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "what.Happens - a debug logging package for developers only"
description = "Package what provides some handy debug-logging functions that can be enabled and disabled via build flags."
author = "Christoph Berger"
email = "info@appliedgo.net"
date = "2020-04-12"
draft = "false"
categories = ["DevOps"]
tags = ["logging", "development"]
articletypes = ["Tools And Libraries"]
vanity = "github.com/appliedgocode/what"
+++

Package what provides some handy debug-logging functions that can be enabled and disabled via build flags. No more information leaks in your production code!

<!--more-->

## Why a "debug-level" log package?

There are two reasons for having a special debug-level log package.

### Reason 1: two log levels are all you'll ever need.

By coincidence I recently came across an old blog post of Dave Cheney again. Back in 2015, Dave took a critical look at the usefulness of log levels. Common logging packages use log levels like Info, Warning, Error, and Fatal, and some even add more, like Trace, Debug, or Critical. After discussing all these log levels one by one, Dave concluded that there are only two log levels you will ever need:

**Debug** and **Info.**

Yes, just Debug and Info. Why those two? The main distinction is the target audience.

* Debug output is meant for developers while they are developing software
* Info output is meant for users of the software.

There is no need for more levels than these. If you read through Dave's article, you will understand this conclusion, and maybe you will even agree.

The idea of reduced log levels goes well with Go's notion of simplicity. In fact, Go's standard `log` package has no log levels at all.


### Reason 2: Debug logs are for developers... only!

On a recent project, I felt the desire of spreading a wagonload of log calls across my code, to see if the code works as intended and to quickly identify a bug when the code doesn't do the right thing. (Full disclaimer: I am not a fan of debuggers. I do not want to discuss the reasons here, they are certainly highly subjective, but I know that during years of debugging my code and the code of others, I was doing quite fine without a debugger.)

However, I did not want all this debug logging code to exist in the final binary that is going to be shipped, for two simple reasons.

* **First, performance.** Every log call would need to check whether or not debug logging is enabled. Although this is just a simple `if` statement, for tons of debug log calls it can add up.
* **Second, security.** Debug logging can spill out an enormous amount of information, and some of it can be sensitive information that you'd rather not want getting exposed to prying eyes.

So the ideal debug library would provide log functions that only exist in test and staging binaries but not in production binaries.

A search on the Web did not reveal a log library with that property, so I decided to write one. And so `what` was born.



## The design of package what


What makes `what` unique is a combination of two aspects.

1. The Functions it provides are few, and their names where chosen to be meaningful and easy to remember.
2. Logging can be switched on and off for (a) specific log functions, and (b) specific packages. Imagine that you have added a lot of `what` calls in your code but for the next test you are only interested in getting debug logs from `what.Happens` and `what.Is`calls, and only for package "mydarncoolpkg" to reduce the noise, you can do both easily via build tags and an environment variable.

Let's look at both aspects in detail.

### Functions

I kept the number of available log functions small, so that the are easy to use and to remember. I added functions where they help reducing boilerplate code (such as extra if statements). Let's look at each of them in detail.

#### what.Happens

`what.Happens()` is the equivalent to `log.Printf()`. It takes the same arguments as Printf(), that is, a format string and zero or more values based on the placeholders in the format string.

Use it to see what happens in your code.

Example:

```go
what.Happens("Connecting to %s", url)
```

```
2020-04-12 01:02:03 main.main: Connecting to https://appliedgo.net/what
```


#### what.If

`what.If()` is like `what.Happens()` but only writes something if the first argument is `true`.

Use it to log something interesting or unusual.

Example:

```go
what.If(!found, "Loop finished, %s not found", searchString)
```

```
2020-04-12 01:02:03 appliedgo.net/mypackage.Search: Loop finished, enlightenment not found
```


#### what.Is

`what.Is()` dumps the contents of a variable. Especially useful for structs.

Example:

```go
client := &http.Client{
		Timeout: 10 * time.Second,
	}
	what.Is(client)
```

```
(*http.Client)(0xc000112e40)({
  Transport: (http.RoundTripper) <nil>,
  CheckRedirect: (func(*http.Request, []*http.Request) error) <nil>,
  Jar: (http.CookieJar) <nil>,
  Timeout: (time.Duration) 10s
})
```

#### what.Func

`what.Func()` prints the name of its caller, as well as the line number and file name where the caller is defined.

Example:

```go
func main() {
    what.Func()
}
```

```
2020/04/12 18:01:31 Func main.main in line 12 of file /Users/christoph/dev/go/playground/what/main.go
```

## Enabling and disabling

`what` logging is enabled at compile time, using the `-tags` flag.

### Enable all

To enable all debug logging, compile your code with `-tags what`:

```sh
go build -tags what
```

### Enable specific functions

To enable either of `Happens()`/`If()`, `Is()`, `Func()`, or `Package()`, pass the respective build tag, or a combination of tags, to `-tag`:

| Function | tag |
| -------- | --- |
| Happens, If | whathappens |
| Is | whatis |
| Func | whatfunc |
| Package | whatpackage |

Example:

```sh
go build -tags whathappens,whatis
```


### Enable specific packages

With `what` calls spread across a number of packages, you might want to get the debug logs of specific packages only, to focus on what's important. Go's build tag mechanism cannot help here, so this is done through an environment variable called "WHAT".

To enable specific packages for debug logging, set `WHAT` to a package import, or a list of package imports.

The environment variable is read only once at startup, so if you change the value while the code is running, the change will not get picked up until restart.

Examples:

```sh
export WHAT=path/to/my/pkg1,repohost.com/user/repo
```

Each entry is a package import path as you would use it in your `import` directive.


### Disabling for production

Package `what` adheres to the initially described philosophy of using only two log levels, "debug" and "info". Package `what` is the debug part. It shall help developers observe their code at runtime and track down unexpected behavior, and then just stay out of the production binary.

For any log output in production, use the standard 'log' package, and remember that your code then talks to the users of your code, not to other developers. Let them know what your code is doing at the moment, help them following progress and finding out why the application got stuck just at this point. Don't write super-detailed info to the production log, or your users will demand log levels for keeping the noise out.


## A short, complete code example
*/

// what happens?
package main

import (
	"log"
	"net/http"
	"time"

	"appliedgo.net/what"
)

func main() {
	// Some log output for the user.
	log.Println("Start")

	// Print the current function.
	what.Func()
	// Print the current package name.
	what.Package()

	// Print what's going on.
	what.Happens("Create HTTP client")
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Dump variable `client`.
	what.Is(client)

	// Print some information about `client`.
	what.Happens("Client has a %f seconds timeout", client.Timeout.Seconds())

	// Let me know that the client has no cookie jar.
	what.If(client.Jar == nil, "client has no cookie jar")

	log.Println("Done.")
}

/*
Copy & paste this code into your own local `main.go` file,

don't forget to run `go mod init main` to get the initial `go.mod` file,

then run `go build -tags what && ./main` (syntax may vary depending on your shell)

and inspect the output. Try using different tags, such as `-tags whathappens`, or add a new package and then exclude it by setting 'WHAT' accordingly.

## The code

`appliedgo.net/what` is a vanity import path. It redirects `go get` to [`github.com/appliedgocode/what`](https://github.com/appliedgocode/what). If you find an issue, please consider filing it in the repo's Issues section.


**Happy coding!**

*/
