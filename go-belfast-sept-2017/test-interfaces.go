import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

type mockClient struct{}

func (mc *mockClient) Get(url string) (resp *http.Response, err error) {
	htmlBody, _ := ioutil.ReadFile("testdata/some_web_page.html")
	r := &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(htmlBody)),
	}
	return r, nil
}
