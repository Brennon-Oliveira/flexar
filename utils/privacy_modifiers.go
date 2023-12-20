package utils

type PrivacyModifier int

var PrivacyModifiers = map[string]int{
	"public":    0,
	"private":   1,
	"protected": 2,
}
