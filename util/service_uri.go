package util

import (
	"net/url"
	"strconv"
)

const (
	serviceDialTimeOut = "dial_time_out"
)

type Service struct {
	scheme      string
	addr        string
	dialTimeOut int
}

func parseService(serviceStr string) (*Service, error) {
	uri, err := url.Parse(serviceStr)
	if err != nil {
		return nil, err
	}
	addr := uri.Hostname() + ":" + uri.Port()
	if uri.Port() == "" {
		addr = uri.Hostname() + ":80"
	}

	dialTimeOut := 0
	if uri.RawQuery != "" {
		rawQueryMap, err := url.ParseQuery(uri.RawQuery)
		if err != nil {
			return nil, err
		} else {
			dialTimeOut, err = strconv.Atoi(rawQueryMap.Get(serviceDialTimeOut))
			if err != nil {
				return nil, err
			}
		}
	}
	return &Service{
		scheme:      uri.Scheme,
		addr:        addr,
		dialTimeOut: dialTimeOut,
	}, nil
}
