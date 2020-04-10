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
title = "what.Happens - a developers-only debug log package"
description = "Package what provides some handy debug-logging functions that can be enabled and disabled via build flags."
author = "Christoph Berger"
email = "info@appliedgo.net"
date = "2020-04-11"
draft = "true"
categories = ["DevOps"]
tags = ["logging", "development"]
articletypes = ["Tools And Libraries"]
vanity = "github.com/appliedgo/what"
+++

Package what provides some handy debug-logging functions that can be enabled and disabled via build flags. No more information leaks in your production code!

<!--more-->

## Why a "debug-level" log package?

### Part 1: an old blog post from Dave

Coincidentally I came across an old blog post of Dave Cheney again recently. Back in 2015, Dave took a critical look at the usefulness of log levels. Common logging packages use log levels like Info, Warning, Error, and Fatal, and some even add more, like Trace, Debug, or Critical. After discussing all these log levels one by one, he concluded that there are only two log levels you will ever need:

Debug and Info.

* Debug output is meant for developers while they are developing software
* Info output is meant for users of the software.

If you read through Dave's article, you will understand this conclusion, and maybe you even agree.

The idea of reduced log levels goes well with Go's notion of simplicity. In fact, Go's standard `log` package has no log levels at all.


### Part 2: Debug logs are for developers... only!

On a recent project, I felt the desire of spreading a wagonload of log calls across my code, to see if the code works as intended and to quickly identify a bug when the code doesn't do the right thing. (Full disclaimer: I am not a fan of debuggers. I do not want to discuss the reasons here, they are certainly highly subjective, but I know that during years of debugging my code and the code of others, I was doing quite fine without a debugger.)

However, I did not want all this debug logging code to exist in the final binary that is going to be shipped, for two simple reasons.

* **First, performance.** Every log call would need to check whether or not debug logging is enabled. Although this is just a simple `if` statement, for tons of debug log calls it can add up.
* **Second, security.** Debug logging can spill out an enormous amount of information, and some of it can be sensitive information that you'd rather not want getting exposed to prying eyes.

So the ideal debug library would provide log functions that only exist in test and staging binaries but not in production binaries.

A search on the Web did not reveal a log library with that property, so I decided to write one. And so `what` was born.


## The design of package what

### Functions

#### what.Happens


#### what.If


#### what.Is


#### what.Func


## Enabling and disabling

### Enable all


### Enable specific functions


### Enable specific packages


### Disabling for production



## The code
*/

// ## Imports and globals
package main

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/TODO:

Step 3. Run the binary.

    go run TODO:.go


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
