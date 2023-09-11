package cli

import (
	"log"
	"net/http"
	"sync"
)

type RequestBody struct{
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup){
	client := &http.Client{}
	req, err := http.NewRequest("GET", translateUrl, nil)

	query := req.URL.Query()

	query.Add("client", "gtx")

	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	req.URL.RawQuery = query.Encode()

	if err != nil {
		 log.Fatal("There was a problem:%s", err)
	}

	client.Do(req)
}