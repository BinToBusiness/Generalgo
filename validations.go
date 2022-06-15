package generalgo

import (
	"encoding/hex"
	"regexp"
)

const (

	//PersonNameValidRegex é um padrão regex para nomes de pessoas válidos.
	PersonNameValidRegex string = `^((\b[A-zÀ-ú']{2,40}\b)\s*){2,}$`

	//DocDatabaseIDLength é o tamanho padrão de IDs em bancos baseados em documentos.
	DocDatabaseIDLength uint8 = 24

	//MinNamesDescriptionsLength é o tamanho mínimo para nomes de pessoas.
	MinNamesDescriptionsLength uint8 = 3
)

//StringMatch valida uma string com base em um padrão regex. Retorna true se corresponde e false se não.
func StringMatch(value, pattern string) bool {
	var rgx *regexp.Regexp = regexp.MustCompile(pattern)

	return rgx.MatchString(value)
}

//ValidIDs valida se um array de bytes é uma string de ID válido
func ValidIDs(id [12]byte) bool {

	var id_s string = hex.EncodeToString(id[:])

	if len(id_s) != 24 {
		return false
	} else if _, err := hex.DecodeString(id_s); err != nil {
		return false
	}
}
