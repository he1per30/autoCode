package main

import (
	parse "autoCode"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	inJson       = "inJson"
	outJson      = "outJson"
	outData      = "outData"
	fileInSystem = "C:\\Users\\Yura\\Desktop\\petProject\\mapping.xlsx"
)

func main() {
	parse.Parse()
	//excel.ReadFromXSL()
}

func parseConfluence2() {
	url := "https://experauz.atlassian.net/wiki/spaces/AE/pages/141230273/Ac.BPM.LoanApp.V1#%D0%A1%D1%82%D0%B0%D1%80%D1%82-%D0%BF%D1%80%D0%BE%D1%86%D0%B5%D1%81%D1%81%D0%B0-%D0%AE%D0%9B"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Basic "+"base64 encoded username:password")
	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
