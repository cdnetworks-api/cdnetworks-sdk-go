<p align="center">
<a href="https://www.cdnetworks.com"><img width="600" height="90" src="https://www.cdnetworks.com/wp-content/uploads/2020/11/cdnetworks-logo-svg.svg"></img></a>
</p>
<h1 align="center">Cdnetworks SDK for Go</h1>

# Table of Contents
1. [Introduction](#introduction)
2. [Installation](#installation)

# Introduction

Welcome to the Cdnetworks Developer Toolkit (SDK). This SDK is the companion development tool for the Cdnetworks OpenApi platform.

# Installation

## Environment Requirements

1. Go version 1.17 and above
2. Some products require activation in the Cdnetworks console before their interfaces can be called normally.
3. Obtain the accessKey ID and accessKey Secret from the [accessKey management](https://dash.cdnetworks.com/account/accessKey) page in the Cdnetworks console. Please keep them safe or use more secure temporary security credentials.

## Installation via go get (Recommended)

It is recommended to use Alibaba Cloud mirror for faster download:

1. Linux or MacOS:

    ```bash
    export GOPROXY=https://mirrors.aliyun.com/goproxy
    ```

2. Windows:

    ```cmd
    set GOPROXY=https://mirrors.aliyun.com/goproxy
    ```

## Install as needed (Recommended)

Note: This installation method only supports dependency management using **Go Modules** mode, i.e., the environment variable `GO111MODULE=auto` or `GO111MODULE=on`, and you have executed `go mod init xxx` in your project.

If you are using GOPATH, please refer to the next section: Complete Installation

```cmd
    go mod tidy
