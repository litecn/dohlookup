package main

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/miekg/dns"
)

func DohQuery(hostName, dnsServer, dnsType string) (dns.Msg, error) {
	respMsg := dns.Msg{}
	if !strings.HasSuffix(hostName, ".") {
		hostName += "."
	}
	query := dns.Msg{}
	query.SetQuestion(hostName, dns.StringToType[strings.ToUpper(dnsType)])

	msg, _ := query.Pack()
	b64 := base64.RawURLEncoding.EncodeToString(msg)
	resp, err := http.Get(dnsServer + "?dns=" + b64)
	if err != nil {
		log.Printf("Send query error, err:%v\n", err)
		return respMsg, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	respMsg.Unpack(bodyBytes)
	return respMsg, nil
}
