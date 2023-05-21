package practice8_7

import (
	"fmt"
	"net/url"
	"strings"
)

type URL struct {
	value url.URL
}

func newURL(urlString string, requestUrl *URL) (*URL, error) {
	urlString = strings.TrimSpace(urlString)

	pu, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	if pu.IsAbs() {
		u := URL{value: *pu}
		return &u, nil
	}

	puPath := pu.Path
	if strings.HasPrefix(puPath, "/") {
		puPath = puPath[1:]
	}

	if strings.HasPrefix(urlString, "/") {
		pu, err := url.ParseRequestURI(fmt.Sprintf("%s://%s/%s", requestUrl.Scheme(), requestUrl.Host(), puPath))
		if err != nil {
			return nil, err
		}

		u := URL{value: *pu}
		return &u, nil
	}

	pu, err = url.ParseRequestURI(fmt.Sprintf("%s/%s", requestUrl.String(), puPath))
	if err != nil {
		return nil, err
	}

	u := URL{value: *pu}
	return &u, nil
}

func (u URL) Scheme() string {
	return u.value.Scheme
}

func (u URL) Host() string {
	return u.value.Hostname()
}

func (u URL) Path() string {
	p := u.value.Path
	p = strings.Trim(p, "/")

	return p
}

func (u URL) String() string {
	s := fmt.Sprintf("%s://%s/%s", u.Scheme(), u.Host(), u.Path())
	if strings.HasSuffix(s, "/") {
		s = s[:len(s)-1]
	}

	return s
}

func (u URL) IsTOP() bool {
	return u.Path() == ""
}

func (u URL) IsSameHost(another URL) bool {
	return u.Host() == another.Host()
}

func (u URL) Extension() string {
	path := u.Path()
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	pathList := strings.Split(path, "/")
	if len(pathList) == 0 {
		return ""
	}

	lastPath := pathList[len(pathList)-1]
	split := strings.Split(lastPath, ".")
	if len(split) <= 1 {
		return ""
	}

	return split[len(split)-1]
}
