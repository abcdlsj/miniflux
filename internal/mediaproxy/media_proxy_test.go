// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package mediaproxy // import "miniflux.app/v2/internal/mediaproxy"

import (
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"miniflux.app/v2/internal/config"
)

func TestProxyFilterWithHttpDefault(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "http-only")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="/proxy/okK5PsdNY8F082UMQEAbLPeUFfbe2WnNfInNmR9T4WA=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlLnBuZw==" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpsDefault(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "http-only")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpNever(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "none")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := input

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpsNever(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "none")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := input

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpAlways(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="/proxy/okK5PsdNY8F082UMQEAbLPeUFfbe2WnNfInNmR9T4WA=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlLnBuZw==" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpsAlways(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="/proxy/LdPNR1GBDigeeNp2ArUQRyZsVqT_PWLfHGjYFrrWWIY=/aHR0cHM6Ly93ZWJzaXRlL2ZvbGRlci9pbWFnZS5wbmc=" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestAbsoluteProxyFilterWithHttpsAlways(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithAbsoluteProxyURL(r, input)
	expected := `<p><img src="http://localhost/proxy/LdPNR1GBDigeeNp2ArUQRyZsVqT_PWLfHGjYFrrWWIY=/aHR0cHM6Ly93ZWJzaXRlL2ZvbGRlci9pbWFnZS5wbmc=" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestAbsoluteProxyFilterWithCustomPortAndSubfolderInBaseURL(t *testing.T) {
	os.Clearenv()
	os.Setenv("BASE_URL", "http://example.org:88/folder/")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	if config.Opts.BaseURL() != "http://example.org:88/folder" {
		t.Fatalf(`Unexpected base URL, got "%s"`, config.Opts.BaseURL())
	}

	if config.Opts.RootURL() != "http://example.org:88" {
		t.Fatalf(`Unexpected root URL, got "%s"`, config.Opts.RootURL())
	}

	router := mux.NewRouter()

	if config.Opts.BasePath() != "" {
		router = router.PathPrefix(config.Opts.BasePath()).Subrouter()
	}

	router.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithAbsoluteProxyURL(router, input)
	expected := `<p><img src="http://example.org:88/folder/proxy/okK5PsdNY8F082UMQEAbLPeUFfbe2WnNfInNmR9T4WA=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlLnBuZw==" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestAbsoluteProxyFilterWithHttpsAlwaysAndAudioTag(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "audio")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<audio src="https://website/folder/audio.mp3"></audio>`
	output := RewriteDocumentWithAbsoluteProxyURL(r, input)
	expected := `<audio src="http://localhost/proxy/EmBTvmU5B17wGuONkeknkptYopW_Tl6Y6_W8oYbN_Xs=/aHR0cHM6Ly93ZWJzaXRlL2ZvbGRlci9hdWRpby5tcDM="></audio>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpsAlwaysAndCustomProxyServer(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_CUSTOM_URL", "https://proxy-example/proxy")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="https://proxy-example/proxy/aHR0cHM6Ly93ZWJzaXRlL2ZvbGRlci9pbWFnZS5wbmc=" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpsAlwaysAndIncorrectCustomProxyServer(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_CUSTOM_URL", "http://:8080example.com")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestAbsoluteProxyFilterWithHttpsAlwaysAndCustomProxyServer(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_CUSTOM_URL", "https://proxy-example/proxy")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithAbsoluteProxyURL(r, input)
	expected := `<p><img src="https://proxy-example/proxy/aHR0cHM6Ly93ZWJzaXRlL2ZvbGRlci9pbWFnZS5wbmc=" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpInvalid(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "invalid")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="/proxy/okK5PsdNY8F082UMQEAbLPeUFfbe2WnNfInNmR9T4WA=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlLnBuZw==" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithHttpsInvalid(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "invalid")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)
	expected := `<p><img src="https://website/folder/image.png" alt="Test"/></p>`

	if expected != output {
		t.Errorf(`Not expected output: got %q instead of %q`, output, expected)
	}
}

func TestProxyFilterWithSrcset(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" srcset="http://website/folder/image2.png 656w, http://website/folder/image3.png 360w" alt="test"></p>`
	expected := `<p><img src="/proxy/okK5PsdNY8F082UMQEAbLPeUFfbe2WnNfInNmR9T4WA=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlLnBuZw==" srcset="/proxy/aY5Hb4urDnUCly2vTJ7ExQeeaVS-52O7kjUr2v9VrAs=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlMi5wbmc= 656w, /proxy/QgAmrJWiAud_nNAsz3F8OTxaIofwAiO36EDzH_YfMzo=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlMy5wbmc= 360w" alt="test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyFilterWithEmptySrcset(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<p><img src="http://website/folder/image.png" srcset="" alt="test"></p>`
	expected := `<p><img src="/proxy/okK5PsdNY8F082UMQEAbLPeUFfbe2WnNfInNmR9T4WA=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlLnBuZw==" srcset="" alt="test"/></p>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyFilterWithPictureSource(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<picture><source srcset="http://website/folder/image2.png 656w,   http://website/folder/image3.png 360w, https://website/some,image.png 2x"></picture>`
	expected := `<picture><source srcset="/proxy/aY5Hb4urDnUCly2vTJ7ExQeeaVS-52O7kjUr2v9VrAs=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlMi5wbmc= 656w, /proxy/QgAmrJWiAud_nNAsz3F8OTxaIofwAiO36EDzH_YfMzo=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlMy5wbmc= 360w, /proxy/ZIw0hv8WhSTls5aSqhnFaCXlUrKIqTnBRaY0-NaLnds=/aHR0cHM6Ly93ZWJzaXRlL3NvbWUsaW1hZ2UucG5n 2x"/></picture>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyFilterOnlyNonHTTPWithPictureSource(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "https")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<picture><source srcset="http://website/folder/image2.png 656w, https://website/some,image.png 2x"></picture>`
	expected := `<picture><source srcset="/proxy/aY5Hb4urDnUCly2vTJ7ExQeeaVS-52O7kjUr2v9VrAs=/aHR0cDovL3dlYnNpdGUvZm9sZGVyL2ltYWdlMi5wbmc= 656w, https://website/some,image.png 2x"/></picture>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyWithImageDataURL(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<img src="data:image/gif;base64,test">`
	expected := `<img src="data:image/gif;base64,test"/>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyWithImageSourceDataURL(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<picture><source srcset="data:image/gif;base64,test"/></picture>`
	expected := `<picture><source srcset="data:image/gif;base64,test"/></picture>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyFilterWithVideo(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "video")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<video poster="https://example.com/img.png" src="https://example.com/video.mp4"></video>`
	expected := `<video poster="/proxy/aDFfroYL57q5XsojIzATT6OYUCkuVSPXYJQAVrotnLw=/aHR0cHM6Ly9leGFtcGxlLmNvbS9pbWcucG5n" src="/proxy/0y3LR8zlx8S8qJkj1qWFOO6x3a-5yf2gLWjGIJV5yyc=/aHR0cHM6Ly9leGFtcGxlLmNvbS92aWRlby5tcDQ="></video>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyFilterVideoPoster(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<video poster="https://example.com/img.png" src="https://example.com/video.mp4"></video>`
	expected := `<video poster="/proxy/aDFfroYL57q5XsojIzATT6OYUCkuVSPXYJQAVrotnLw=/aHR0cHM6Ly9leGFtcGxlLmNvbS9pbWcucG5n" src="https://example.com/video.mp4"></video>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}

func TestProxyFilterVideoPosterOnce(t *testing.T) {
	os.Clearenv()
	os.Setenv("MEDIA_PROXY_MODE", "all")
	os.Setenv("MEDIA_PROXY_RESOURCE_TYPES", "image,video")
	os.Setenv("MEDIA_PROXY_PRIVATE_KEY", "test")

	var err error
	parser := config.NewParser()
	config.Opts, err = parser.ParseEnvironmentVariables()
	if err != nil {
		t.Fatalf(`Parsing failure: %v`, err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/proxy/{encodedDigest}/{encodedURL}", func(w http.ResponseWriter, r *http.Request) {}).Name("proxy")

	input := `<video poster="https://example.com/img.png" src="https://example.com/video.mp4"></video>`
	expected := `<video poster="/proxy/aDFfroYL57q5XsojIzATT6OYUCkuVSPXYJQAVrotnLw=/aHR0cHM6Ly9leGFtcGxlLmNvbS9pbWcucG5n" src="/proxy/0y3LR8zlx8S8qJkj1qWFOO6x3a-5yf2gLWjGIJV5yyc=/aHR0cHM6Ly9leGFtcGxlLmNvbS92aWRlby5tcDQ="></video>`
	output := RewriteDocumentWithRelativeProxyURL(r, input)

	if expected != output {
		t.Errorf(`Not expected output: got %s`, output)
	}
}
