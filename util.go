package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var iptv_chan chan string

// ParseAddresses parses the m3u or m3u8 file specified by url address addr and returns a string slice containing all the IPTV addresses
func ParseAddresses(m3u string) ([]string, error) {
	f, err := os.Open(m3u)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	var channels []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "http") {
			channels = append(channels, line)
		}
	}
	return channels, nil
}

func SaveValidChannel(fd string) {
	for iptv := range iptv_chan {
		fmt.Println(iptv)//TODO
	}
}