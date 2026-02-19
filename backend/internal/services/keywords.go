package services

type KeywordExtractor interface{ Extract(text string)([]string,error) }

type SimpleKeywordExtractor struct{}

func(e SimpleKeywordExtractor)Extract(t string)([]string,error){return nil,nil}