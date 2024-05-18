// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sitemap

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSitemap(t *testing.T) {
	err := NewSitemap().
		XmlnsImage("https://www.google.com/schemas/sitemap-image/1.1").
		Url(
			"https://xxx.cn/posts/1",
			WithLastMod("2024-05-17"),
			WithChangeFreq("weekly"),
			WithPriority(1.0),
			WithImage(NewUrlImage("https://xxx.cn/posts/1.jpg")),
		).
		Url("https://xxx.cn/posts/2").Output("sitemap.xml").
		GenerateXml()
	require.NoError(t, err)

	result := `
<urlset xmlns="https://www.sitemaps.org/schemas/sitemap/0.9" xmlns:image="https://www.google.com/schemas/sitemap-image/1.1">
  <url>
    <loc>https://xxx.cn/posts/1</loc>
    <lastmod>2024-05-17</lastmod>
    <changefreq>weekly</changefreq>
    <priority>1</priority>
    <image:image>
      <image:ioc>https://xxx.cn/posts/1.jpg</image:ioc>
    </image:image>
  </url>
  <url>
    <loc>https://xxx.cn/posts/2</loc>
  </url>
</urlset>
`
	content, err := os.ReadFile("sitemap.xml")
	require.NoError(t, err)
	require.Equal(t, strings.TrimSpace(result), strings.TrimSpace(string(content)))
}
