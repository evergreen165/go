package main

import (
	"fmt"

	"github.com/deatil/go-encoding/base62"
)

// base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz

func main() {

	// Encode
	src := []byte("11157")
	encoded := base62.StdEncoding.EncodeToString(src)
	fmt.Println("Encode: ", encoded)

	// Decode
	encodedStr := "11157"
	decoded, err := base62.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		fmt.Println("Erro ao decodificar: ", err)
		return
	}
	fmt.Println("Decode: ", string(decoded))

	// Mapeamento
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	fmt.Println("mapeando")
	for i := 0; i < 62; i++ {
		fmt.Printf("valor %2d -> caractere '%c'\n", i, alphabet[i])
	}

	// codificacao
	for i := 0; i < 65; i++ {
		encode := base62.StdEncoding.EncodeToString([]byte{byte(i)})
		fmt.Printf("byte %3d (%02x) -> base62: %-4s\n", i, i, encode)
	}

}
