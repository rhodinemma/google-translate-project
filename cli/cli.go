package cli

import(

)

type RequestBody struct{
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody){
	client := &http.Client{}
	req

	client.Do(req)
}