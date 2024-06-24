package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"strings"
)

func Post(url string, result interface{}, body ...map[string]interface{}) error {
	client := resty.New().R()
	if len(body) > 0 {
		client = client.SetBody(body[0])
	}

	resp, err := client.
		SetHeader("Accept", "application/json").
		SetResult(result).
		Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("http status: %d", resp.StatusCode())
	}
	return nil
}

func Get(url string, result interface{}, query ...map[string]string) error {
	client := resty.New().R()
	if len(query) > 0 {
		client = client.SetQueryParams(query[0])
	}

	resp, err := client.
		SetHeader("Accept", "application/json").
		SetResult(result).
		Get(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("http status: %d", resp.StatusCode())
	}
	return nil
}

func Call(method, url string, headers map[string][]string, requestPayloads ...func(*resty.Request) error) (*resty.Response, error) {
	client := resty.New().R()
	for _, payload := range requestPayloads {
		if err := payload(client); err != nil {
			return nil, err
		}
	}
	resp, err := client.SetHeaderMultiValues(headers).Execute(method, url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("status code : %d", resp.StatusCode())
	}

	return resp, nil
}

func GinProxy(c *gin.Context) error {
	req := c.Request

	proxyAddr := Path2ProxyAddr[req.URL.Path]
	if proxyAddr == "" {
		return errors.New("GinProxy proxyAddr not found")
	}
	target := fmt.Sprintf("%s%s", proxyAddr, req.URL.Path)

	resp, err := Call(req.Method, target, c.Request.Header, func(request *resty.Request) error {
		if strings.ToUpper(req.Method) == "POST" {
			postBody := make(map[string]interface{})
			if err := c.Bind(&postBody); err != nil {
				return errors.Wrap(err, "Bind err")
			}
			request.SetBody(postBody)
		}
		return nil
	})

	if err != nil {
		return errors.Wrap(err, "Call err")
	}

	_, err = c.Writer.Write(resp.Body())
	return err
}
