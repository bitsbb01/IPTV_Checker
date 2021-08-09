package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type checker struct {
	client *http.Client
	request *http.Request
	addr string
}

func NewChecker(addr string) *checker {
	c := checker {
		addr: addr,
	}
	c.client = &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	req, err := http.NewRequest("GET", c.addr, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
		"User-Agent": []string{"Mozilla/5.0 (Windows NT 6.2; Win64; x64;) Gecko/20100101 Firefox/20.0"},
		"Connection": []string{"keep-alive"},

	}
	c.request = req
	return &c
}

func (c *checker) Do() error {
	res, err := c.client.Do(c.request)
	if err != nil {
		return errors.New(c.addr + " " + err.Error())
	}
	tp := res.Header.Get("Content-Type")
	if tp == "application/x-mpegURL" || tp == "application/x-mpegurl" || tp == "video/mp2t" || tp == "application/vnd.apple.mpegurl" || tp == "application/octet-stream" {
		return nil
	}
	return errors.New(c.addr + " unavailable now...")
}