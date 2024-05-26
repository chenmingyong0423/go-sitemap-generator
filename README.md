<h1 align="center">
  go-sitemap-generator
</h1>

[![GitHub Repo stars](https://img.shields.io/github/stars/chenmingyong0423/go-sitemap-generator)](https://github.com/chenmingyong0423/go-sitemap-generator/stargazers)
[![GitHub issues](https://img.shields.io/github/issues/chenmingyong0423/go-sitemap-generator)](https://github.com/chenmingyong0423/go-sitemap-generator/issues)
[![GitHub License](https://img.shields.io/github/license/chenmingyong0423/go-sitemap-generator)](https://github.com/chenmingyong0423/go-sitemap-generator/blob/main/LICENSE)
[![GitHub release (with filter)](https://img.shields.io/github/v/release/chenmingyong0423/go-sitemap-generator)](https://github.com/chenmingyong0423/go-sitemap-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/chenmingyong0423/go-sitemap-generator)](https://goreportcard.com/report/github.com/chenmingyong0423/go-sitemap-generator)
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)

`go-sitemap-generator` A sitemap generator in Go. It can generate a sitemap in XML format.

English | [中文简体](./README-zh_CN.md)

## Installation
```bash
go get github.com/chenmingyong0423/go-sitemap-generator
```

## Usage
```go
err := sitemap.NewSitemap().
    Url(
        "https://xxx.cn/posts/1",
        WithLastMod("2024-05-17"),
        WithChangeFreq("weekly"),
        WithPriority(1.0),
    ).
    Url("https://xxx.cn/posts/2").Output("sitemap.xml").
    GenerateXml()
```
