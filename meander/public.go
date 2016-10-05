package meander

// Facade (ファサード)はデザインパターンの一種。外部向けのビューを表す。
type Facade interface {
	Public() interface{}
}

// Public はあるオブジェクトがFacadwインターフェイスを実装しているか調べる
// されているならPublic()を返し、されていないならそのオブジェクトをそのまま返す
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok { // Public interface{}があるか？
		return p.Public()
	}
	return o
}
