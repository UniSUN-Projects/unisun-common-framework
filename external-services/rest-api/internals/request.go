package internals

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/UniSUN-Projects/unisun-common-framework/external-services/rest-api/constant"
	"github.com/UniSUN-Projects/unisun-common-framework/external-services/rest-api/models"
)

type OptionsHandle struct {
	Method string
	Url    string
	Body   []byte
	Option models.Option
}

func NewOptionsHandle(_method string, _url string, _body []byte) *OptionsHandle {
	return &OptionsHandle{
		Method: _method,
		Url:    _url,
		Body:   _body,
		Option: models.Config.AppConig.Authorization,
	}
}

func (option *OptionsHandle) Call(model any) error {
	var body *bytes.Reader = nil
	if checkBody(option.Method) {
		body = bytes.NewReader(option.Body)
	}
	var req *http.Request = nil
	var err error = nil
	if req, err = http.NewRequest(option.Method, option.Url, body); err != nil {
		return err
	}
	client := mapHeader(req, option.Option)
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return err
	} else {
		if resBody, err := io.ReadAll(res.Body); err != nil {
			return err
		} else {
			if err := json.Unmarshal([]byte(resBody), model); err != nil {
				return err
			}
		}
	}
	defer res.Body.Close()
	return nil
}

func mapHeader(req *http.Request, option models.Option) *http.Client {
	if option.Header.Authorization.Type != "" && option.Header.Authorization.Token != "" {
		req.Header.Add(constant.HEADER_AUTHORIZATION, strings.Join([]string{option.Header.Authorization.Type, option.Header.Authorization.Token}, " "))
	}
	if option.Header.ContentType != "" {
		req.Header.Add(constant.HEADER_CONTENT_TYPE, option.Header.ContentType)
	}
	if option.Header.UserAgent != "" {
		req.Header.Set(constant.HEADER_USER_AGENT, option.Header.UserAgent)
	}
	client := &http.Client{
		Timeout: time.Second * time.Duration(option.Timeout),
	}
	return client
}

func checkBody(method string) bool {
	switch method {
	case http.MethodGet:
		return false
	case http.MethodPost:
		return true
	case http.MethodDelete:
		return false
	case http.MethodPut:
		return true
	default:
		return false
	}
}
