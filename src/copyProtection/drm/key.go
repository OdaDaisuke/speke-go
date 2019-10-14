package drm

import "fmt"

type KeyLib struct {
}

func (k KeyLib) GenerateKey(kid, seed string) string {
	fmt.Println(kid, seed)
	return ""
}

func (k KeyLib) GenerateIV() string {
	return ""
}
