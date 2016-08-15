package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BigHuge はBig Huge Thesaurus APIのAPI Keyを保存する型。
type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:noun` // "encoding/json"が見つけられるよう、フィールドはエキスポートされなければならない
	Verb *words `json:verb`
}

type words struct {
	Syn []string `json:syn`
}

// Synonyms は、ある単語の類義語をBig Huge Thesaurus APIを利用して見つけるメソッド
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json")
	if err != nil {
		return syns, fmt.Errorf("BigHuge: %qの類語検索に失敗しました: %v", term, err)
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns, nil
}
