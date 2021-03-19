package main

import (
	"math/big"
	"testing"
)

// these are just placeholders to remind me what testing is for when i wake up lol

func BenchmarkProve(b *testing.B) {
	m := big.NewInt(1)
	r := big.NewInt(2)
	c := commit(m, r)
	var ck []CurvePoint
	ck = append(ck, CurvePoint{Group.Gx, Group.Gy})
	ck = append(ck, H)
	ca,cb,f,za,zb := Prover(ck, c, m, r)
	Verifier(ck,c,ca,cb,f,za,zb)

	// do a thing here

	//b.ResetTimer()


		//for i := 0; i < b.N; i++ {
		//	kv := vals[keys[i%len(keys)]]
		//	if trie.Prove(kv.k) == nil {
		//		b.Fatalf("nil proof for %x", kv.k)
		//	}
		//}
}

func BenchmarkVerify(b *testing.B) {

	// do some things here too
}
