package http

import (
	"io/ioutil"
	h "net/http"

	"github.com/parkhomchik/fba/db"
)

type HttpManager struct {
	Manager db.DBManager
}

func (hm *HttpManager) Send(method, url, auth string) (body []byte, status int, err error) {
	req, err := h.NewRequest(method, url, nil)
	req.Header.Add("Authorization", auth)
	cl := h.Client{}
	resp, err := cl.Do(req)
	status = resp.StatusCode
	if status == 200 {
		body, err = ioutil.ReadAll(resp.Body)
	}
	return
}
