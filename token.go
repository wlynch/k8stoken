package k8stoken

import (
	"net/http"
	"strings"

	"k8s.io/client-go/rest"
)

type fakeRT struct {
	token string
}

func (rt *fakeRT) RoundTrip(req *http.Request) (*http.Response, error) {
	rt.token = req.Header.Get("Authorization")
	return &http.Response{}, nil
}

func Token(config *rest.Config) (string, error) {
	fake := &fakeRT{}
	config.Wrap(func(rt http.RoundTripper) http.RoundTripper {
		return fake
	})
	client, err := rest.TransportFor(config)
	if err != nil {
		return "", err
	}
	if _, err = client.RoundTrip(&http.Request{}); err != nil {
		return "", err
	}

	return strings.TrimPrefix(fake.token, "Bearer "), nil
}
