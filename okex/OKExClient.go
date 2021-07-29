package okex

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/okex/V3-Open-API-SDK/okex-go-sdk-api"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

/*
 http client, request, response
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

type Client struct {
	Config     okex.Config
	HttpClient *http.Client
}

type ApiMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
 Send a http request to remote server and get a response data
*/
func (client *Client) Request(method string, requestPath string,
	params, result interface{}) (response *http.Response, err error) {
	config := client.Config
	// uri
	endpoint := config.Endpoint
	if strings.HasSuffix(config.Endpoint, "/") {
		endpoint = config.Endpoint[0 : len(config.Endpoint)-1]
	}
	url := endpoint + requestPath

	// get json and bin styles request body
	var jsonBody string
	var binBody = bytes.NewReader(make([]byte, 0))
	if params != nil {
		jsonBody, binBody, err = okex.ParseRequestParams(params)
		if err != nil {
			return response, err
		}
	}

	// get a http request
	request, err := http.NewRequest(method, url, binBody)
	if err != nil {
		return response, err
	}

	// Sign and set request headers
	timestamp := okex.IsoTime()
	preHash := okex.PreHashString(timestamp, method, requestPath, jsonBody)
	sign, err := okex.HmacSha256Base64Signer(preHash, config.SecretKey)
	if err != nil {
		return response, err
	}
	Headers(request, config, timestamp, sign)

	if config.IsPrint {
		printRequest(config, request, jsonBody, preHash)
	}

	// send a request to remote server, and get a response
	response, err = client.HttpClient.Do(request)
	if err != nil {
		return response, err
	}
	defer response.Body.Close()

	// get a response results and parse
	status := response.StatusCode
	message := response.Status
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response, err
	}

	if config.IsPrint {
		printResponse(status, message, body)
	}

	responseBodyString := string(body)

	response.Header.Add(okex.ResultDataJsonString, responseBodyString)

	limit := response.Header.Get("Ok-Limit")
	if limit != "" {
		var page okex.PageResult
		page.Limit = okex.StringToInt(limit)
		from := response.Header.Get("Ok-From")
		if from != "" {
			page.From = okex.StringToInt(from)
		}
		to := response.Header.Get("Ok-To")
		if to != "" {
			page.To = okex.StringToInt(to)
		}
		pageJsonString, err := okex.Struct2JsonString(page)
		if err != nil {
			return response, err
		}
		response.Header.Add(okex.ResultPageJsonString, pageJsonString)
	}

	if status >= 200 && status < 300 {
		if body != nil && result != nil {
			err := okex.JsonBytes2Struct(body, result)
			if err != nil {
				return response, err
			}
		}
		return response, nil
	} else if status >= 400 || status <= 500 {
		errMsg := "Http error(400~500) result: status=" + okex.IntToString(status) + ", message=" + message + ", body=" + responseBodyString
		fmt.Println(errMsg)
		if body != nil {
			err := errors.New(errMsg)
			return response, err
		}
	} else {
		fmt.Println("Http error result: status=" + okex.IntToString(status) + ", message=" + message + ", body=" + responseBodyString)
		return response, errors.New(message)
	}
	return response, nil
}

func Headers(request *http.Request, config okex.Config, timestamp string, sign string) {
	request.Header.Add(okex.ACCEPT, okex.APPLICATION_JSON)
	request.Header.Add(okex.CONTENT_TYPE, okex.APPLICATION_JSON_UTF8)
	request.Header.Add(okex.COOKIE, okex.LOCALE+config.I18n)
	request.Header.Add(okex.OK_ACCESS_KEY, config.ApiKey)
	request.Header.Add(okex.OK_ACCESS_SIGN, sign)
	request.Header.Add(okex.OK_ACCESS_TIMESTAMP, timestamp)
	request.Header.Add(okex.OK_ACCESS_PASSPHRASE, config.Passphrase)
	//request.Header.Add("x-simulated-trading", "1")
}

func printRequest(config okex.Config, request *http.Request, body string, preHash string) {
	if config.SecretKey != "" {
		fmt.Println("  Secret-Key: " + config.SecretKey)
	}
	fmt.Println("  Request(" + okex.IsoTime() + "):")
	fmt.Println("\tUrl: " + request.URL.String())
	fmt.Println("\tMethod: " + strings.ToUpper(request.Method))
	if len(request.Header) > 0 {
		fmt.Println("\tHeaders: ")
		for k, v := range request.Header {
			if strings.Contains(k, "Ok-") {
				k = strings.ToUpper(k)
			}
			fmt.Println("\t\t" + k + ": " + v[0])
		}
	}
	fmt.Println("\tBody: " + body)
	if preHash != "" {
		fmt.Println("  PreHash: " + preHash)
	}
}

func printResponse(status int, message string, body []byte) {
	fmt.Println("  Response(" + okex.IsoTime() + "):")
	statusString := strconv.Itoa(status)
	message = strings.Replace(message, statusString, "", -1)
	message = strings.Trim(message, " ")
	fmt.Println("\tStatus: " + statusString)
	fmt.Println("\tMessage: " + message)
	fmt.Println("\tBody: " + string(body))
}
