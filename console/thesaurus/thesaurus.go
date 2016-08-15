package thesaurus

// Thesaurus は類義語を求めるSynonymsというinterfaceを実装している
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
