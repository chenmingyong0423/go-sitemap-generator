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
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSitemap(t *testing.T) {
	err := NewSitemap().Url(
		"https://chenmingyong.cn/posts/1",
		WithLastMod("2024-05-17"), WithChangeFreq("weekly"), WithPriority(1.0)).
		Url("https://chenmingyong.cn/posts/2").Output("sitemap.xml").GenerateXml()
	require.NoError(t, err)
	_, err = os.ReadFile("sitemap.xml")
	require.NoError(t, err)

}
