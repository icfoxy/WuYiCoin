package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/icfoxy/GoTools"
	"github.com/joho/godotenv"
)

func JoinNet(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	resp, err := http.Get("http://" + os.Getenv("GuideNode") + "/getNodes")
	if err != nil {
		GoTools.RespondByErr(w, 801, err.Error(), "high")
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
	}
	for i, node := range nodes {
		err = os.Setenv("node"+fmt.Sprint(i), node)
		if err != nil {
			GoTools.RespondByErr(w, 801, err.Error(), "high")
		}
	}
	GoTools.RespondByJSON(w, 200, "Done")
}
