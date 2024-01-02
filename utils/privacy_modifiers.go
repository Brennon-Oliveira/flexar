package utils

type PrivacyModifier int

var PrivacyModifiers = map[string]PrivacyModifier{
	"public":    0,
	"private":   1,
	"protected": 2,
}
