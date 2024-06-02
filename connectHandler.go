package main

import (
	"net"
	"net/http"
	"os"

	"github.com/icfoxy/GoTools"
	"github.com/joho/godotenv"
)

func JoinNet(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			localAddr, err := net.ResolveTCPAddr(network, "localhost:"+os.Getenv("SendPort"))
			if err != nil {
				return nil, err
			}
			remoteAddr, err := net.ResolveTCPAddr(network, addr)
			if err != nil {
				return nil, err
			}
			return net.DialTCP(network, localAddr, remoteAddr)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get("http://" + os.Getenv("GuideNode") + "/getNodes")
	if err != nil {
		GoTools.RespondByErr(w, 801, err.Error(), "high")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 801 {
		var errData GoTools.ErrResp
		err = GoTools.GetAnyFromBody(resp.Body, &errData)
		if err != nil {
			GoTools.RespondByErr(w, 801, err.Error(), "high")
			return
		}
		GoTools.RespondByErr(w, 801, errData.ErrInfo, errData.ErrLevel)
		return
	}
	var nodes []string
	err = GoTools.GetAnyFromBody(resp.Body, &nodes)
	if err != nil {
		GoTools.RespondByErr(w, 801, err.Error(), "high")
		return
	}
	if len(nodes) == 0 {
		GoTools.RespondByErr(w, 801, "missing nodes info", "low")
		return
	}
	err = GoTools.DBPut("/"+os.Getenv("Port"), "nodes", nodes)
	if err != nil {
		GoTools.RespondByErr(w, 801, "DBPut wrong", "high")
		return
	}
	GoTools.RespondByJSON(w, 200, "Done")
}
