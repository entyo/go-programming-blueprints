package meander

import "strings"

// Cost は価格帯を擬似enumで表現するための型
type Cost int8

const (
	// iota の最初の値は0、これを捨てる
	_ Cost = iota
	// Cost1 == 1
	Cost1
	// Cost2 == 2
	Cost2
	// Cost3 == 3
	Cost3
	// Cost4 == 4
	Cost4
	// Cost5 == 5
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (l Cost) String() string {
	for s, v := range costStrings {
		if l == v {
			return s
		}
	}
	return "不正な値です"
}

// ParseCost は渡された文字列に対応したコストの値を返す
func ParseCost(s string) Cost {
	return costStrings[s]
}

// CostRange はCost型を使って価格帯を表す
type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

// ParseCostRange は価格帯を表す文字列( i.e.) $$...$$$$$ )をパースする"
func ParseCostRange(s string) *CostRange {
	segs := strings.Split(s, "...")
	return &CostRange{
		From: ParseCost(segs[0]),
		To:   ParseCost(segs[1]),
	}
}
