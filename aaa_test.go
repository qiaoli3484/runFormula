package runFormula

import "testing"

func TestAaa(t *testing.T) {

	t.Log(Run("2*9+3-2*(10-3)/14"))
	t.Log("hello world")
}

func BenchmarkAaa(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run("if(3+3>5)2*9+3-2*(10-3)/14;500")
	}
}
