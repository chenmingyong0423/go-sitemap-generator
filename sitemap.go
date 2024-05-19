package sitemap

import (
	"encoding/xml"
	"os"
)

type Sitemap struct {
	UrlSet     UrlSet
	xmlContent []byte
	output     string
}

// NewSitemap creates a new Sitemap
// the xmlns is defaulted to https://www.sitemaps.org/schemas/sitemap/0.9
// you can change it by calling Sitemap.Xmlns Method
func NewSitemap() *Sitemap {
	return &Sitemap{
		UrlSet: UrlSet{Xmlns: "https://www.sitemaps.org/schemas/sitemap/0.9"},
	}
}

func (s *Sitemap) Xmlns(value string) *Sitemap {
	s.UrlSet.Xmlns = value
	return s
}

func (s *Sitemap) XmlnsImage(value string) *Sitemap {
	s.UrlSet.XmlnsImage = value
	return s
}

func (s *Sitemap) Url(loc string, opts ...URLOption) *Sitemap {
	s.UrlSet.Urls = append(s.UrlSet.Urls, *NewURL(loc, opts...))
	return s
}

func (s *Sitemap) Output(value string) *Sitemap {
	s.output = value
	return s
}

func (s *Sitemap) GenerateXml() error {
	// 创建文件
	file, err := os.Create(s.output)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入文件
	_, err = file.WriteString(xml.Header)
	if err != nil {
		return err
	}

	xmlContent := s.xmlContent
	if len(xmlContent) == 0 {
		xmlContent, err = s.UrlSet.Marshal()
		if err != nil {
			return err
		}
	}

	_, err = file.Write(xmlContent)
	if err != nil {
		return err
	}
	return nil
}
