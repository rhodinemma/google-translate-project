package main

import(
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

var sourceLang string
var targetLang string
var sourceText string


func init(){
	flag.StringVar(&sourceLang, "s", "en", "Source language[en]")
	flag.StringVar(&targetLang, "t", "fr", "Target language[fr]")
	flag.StringVar(&sourceText, "st", "", "Text to translate")
}

func main(){
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	reqBody := &cli.RequestBody{
		SourceLang: sousourceLang,
		TargetLang: targtargetLang,
		SourceText: soursourceText,
	}

	cli.RequestTranslate(reqBody)
}