package services

type Summarizer interface{ Summarize(text string)(string, error) }

type SimpleSummarizer struct{}

func(s SimpleSummarizer)Summarize(t string)(string,error){return "",nil}