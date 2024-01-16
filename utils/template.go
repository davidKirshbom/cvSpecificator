package utils

import (
	"html/template"
	"strings"
)

func WithComData(pairs ...interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i := 0; i < len(pairs); i += 2 {
		result[pairs[i].(string)] = pairs[i+1]
	}
	return result
}

//currently only supports one list in the end of string
func Htmlize(str string) template.HTML {
	if strings.Contains(str, "\n-") {
	sentences := strings.Split(str, "\n-")
	htmlizedSentences := ""
	startString := ""
	for _, sentence := range sentences {
		if(strings.HasPrefix(sentence," ")){
		htmlizedSentence := "<li>" + sentence + "</li>"
		htmlizedSentences += htmlizedSentence
		}else{
			startString += sentence
		}

	}
	htmlizedStr :=startString+ "<ul>" + htmlizedSentences + "</ul>"
	return template.HTML(htmlizedStr)
}
return template.HTML(str)
}

func HtmlEmphasisWords(str string, words []string) template.HTML {
	for _, word := range words {
		str = strings.ReplaceAll(str, word, "<strong>"+word+"</strong>")
	}
	return template.HTML(str)
}