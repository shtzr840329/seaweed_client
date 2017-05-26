package filer

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/pquerna/ffjson/ffjson"
)

// Client ...
type Client struct {
	Server   string
	FileName string
}

// SubmitResult ...
type SubmitResult struct {
	FileName string `json:"fileName,omitempty"`
	FileURL  string `json:"fileUrl,omitempty"`
	Fid      string `json:"fid,omitempty"`
	Size     uint32 `json:"size,omitempty"`
	Error    string `json:"error,omitempty"`
}

// Post ...
func (c *Client) Post() (result SubmitResult) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("userfile", c.FileName)

	if err != nil {
		panic(err)
	}

	fh, err := os.Open("IOS.jpg")

	if err != nil {
		panic(err)
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		panic(err)
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(c.Server, contentType, bodyBuf)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		ffjson.Unmarshal(body, &result)
	} else {
		panic(err)
	}
	return
}

// Delete ...
func (c *Client) Delete(url string) (result SubmitResult) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	} else {
		ffjson.Unmarshal(body, &result)
	}
	return
}

// Put ...
func (c *Client) Put() {

}
