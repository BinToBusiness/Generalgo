package generalgo

import (
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
