package Util

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func MakeHTTPRequest(URL string, Method string, Headers map[string]string, Body interface{}) (*http.Response, error) {

	var BodyReader io.Reader

	if Body != nil {

		BodyBytes, MarshalErr := json.Marshal(Body)

		if MarshalErr != nil {

			return nil, MarshalErr

		}

		BodyReader = bytes.NewReader(BodyBytes)

	}

	Client := &http.Client{}

	Req, Error := http.NewRequest(Method, URL, BodyReader)

	if Error != nil {

		return nil, Error

	}

	for Key, Value := range Headers {

		Req.Header.Set(Key, Value)

	}

	Resp, Error := Client.Do(Req)

	if Error != nil {

		return nil, Error

	}

	return Resp, nil

}