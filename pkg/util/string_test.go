package util

import (
	"testing"
)

func TestRandomHexString(t *testing.T) {
    want_len := 10
    str := RandHexString(want_len)

    for i := 0; i < want_len; i++ { 
        ch := str[i]
        if ch < '0' || ch > '9' && ch < 'A' || ch > 'F' && ch < 'a' || ch > 'f'{ 
            t.Fatalf("String %s contains a not hex value at index %d", str, i)
        } 
    }
 
    str_len := len(str)
    if str_len != want_len {
        t.Fatalf("Invalid string lenght: %d ", str_len)
    }
}

func BenchmarkRandomHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
        RandHexString(10)
    }
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
        RandString(10)
    }
}
