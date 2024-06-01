package main

import (
	"net/http"

	"github.com/icfoxy/GoTools"
)

func GetNodeList(guideUrl string) (result []string, err error) {
	resp, err := http.Get(guideUrl + "/JoinNet")
	if err != nil {
		return nil, err
	}
	var nodes []string
	err = GoTools.GetAnyFromBody(resp.Body, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}
