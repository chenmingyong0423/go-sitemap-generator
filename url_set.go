// Generated by [optioner] command-line tool; DO NOT EDIT
// If you have any questions, please create issues and submit contributions at:
// https://github.com/chenmingyong0423/go-optioner

package sitemap

import (
	"encoding/xml"
)

type UrlSet struct {
	XMLName    xml.Name `xml:"urlset"`
	Xmlns      string   `xml:"xmlns,attr"`
	XmlnsImage string   `xml:"xmlns:image,attr"`
	Urls       []URL    `xml:"url"`
}

func (urlSet *UrlSet) Marshal() ([]byte, error) {
	return xml.MarshalIndent(urlSet, "", "  ")
}

//go:generate optioner -type URL -output url_set.go -mode append
type URL struct {
	Loc        string    `xml:"loc" opt:"-"`
	LastMod    string    `xml:"lastmod,omitempty"`
	ChangeFreq string    `xml:"changefreq,omitempty"`
	Priority   float64   `xml:"priority,omitempty"`
	Image      *UrlImage `xml:"image:image,omitempty"`
}

type UrlImage struct {
	Loc string `xml:"image:loc"`
}

func NewUrlImage(ioc string) *UrlImage {
	return &UrlImage{Loc: ioc}
}

type URLOption func(*URL)

func NewURL(loc string, opts ...URLOption) *URL {
	uRL := &URL{
		Loc: loc,
	}

	for _, opt := range opts {
		opt(uRL)
	}

	return uRL
}

func WithLastMod(lastMod string) URLOption {
	return func(uRL *URL) {
		uRL.LastMod = lastMod
	}
}

func WithChangeFreq(changeFreq string) URLOption {
	return func(uRL *URL) {
		uRL.ChangeFreq = changeFreq
	}
}

func WithPriority(priority float64) URLOption {
	return func(uRL *URL) {
		uRL.Priority = priority
	}
}

func WithImage(image *UrlImage) URLOption {
	return func(uRL *URL) {
		uRL.Image = image
	}
}
