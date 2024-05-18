package sitemap

type Sitemap struct {
	UrlSet UrlSet
	output string
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
	return s.UrlSet.GenerateXml(s.output)
}
