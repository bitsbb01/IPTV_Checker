package main

import "fmt"

var iptv_chan chan string

// ParseAddresses parses the m3u or m3u8 file specified by url address addr and returns a string slice containing all the IPTV addresses
func ParseAddresses(addr string) ([]string, error) {
	return nil,nil
}

func SaveValidChannel(fd string) {
	for iptv := range iptv_chan {
		fmt.Println(iptv)//TODO
	}
}