package runformula

import "testing"

func TestAaa(t *testing.T) {

	t.Log(Run("(1030+977)*2000/1000000*20", 2))
	t.Log("hello world")
}

func BenchmarkAaa(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run("if(3+3>5)22*90+3-20*(10.5-3)/14;500", -1)
	}
}
