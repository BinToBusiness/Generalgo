package generalgo

import (
	"regexp"
	"strconv"
)

const (

	//PersonNameValidRegex é um padrão regex para nomes de pessoas válidos.
	PersonNameValidRegex string = `^(?=.{2,40}$)(([A-zÀ-ú']{1,20} )*[A-zÀ-ú']{1,20}$)`

	//DocDatabaseIDLength é o tamanho padrão de IDs em bancos baseados em documentos.
	DocDatabaseIDLength uint8 = 24

	//MinNamesDescriptionsLength é o tamanho mínimo para nomes de pessoas.
	MinNamesDescriptionsLength uint8 = 3

	//CpfRegex é um padrão regex para formato de número de cpf.
	CpfRegexNoMask = `^[0-9]{11}$`

	//CnpjRegex é um padrão regex para formato de número de cpf.
	CnpjRegexNoMask = `^[0-9]{14}$`
)

//StringMatch valida uma string com base em um padrão regex. Retorna true se corresponde e false se não.
func StringMatch(value, pattern string) bool {
	var rgx *regexp.Regexp = regexp.MustCompile(pattern)

	return rgx.MatchString(value)
}

//CpfIsValid valida se um número de cpf é valido.
func CpfIsValid(docnum string) bool {

	return false
}

// CnpjIsValid valida se um número de CNPJ é válido.
func CnpjIsValid(docnum string) bool {

	if !StringMatch(docnum, CnpjRegexNoMask) {
		return false
	}

	var cnpj_valid string = docnum[:12]

	var base, digit uint8
	var sum uint16

CALCDIGIT:

	base, digit, sum = 2, 0, 0

	for i, j := 0, len(cnpj_valid)-1; j >= 0; i, j = i+1, j-1 {

		if base > 9 {
			base = 2
		}

		sum += uint16(uint8(cnpj_valid[j]-'0') * base)

		base++
	}

	digit = uint8(11 - (sum % 11))

	if digit >= 10 {
		cnpj_valid += "0"
	} else {
		cnpj_valid += strconv.Itoa(int(digit))
	}

	if len(cnpj_valid) < 14 {
		goto CALCDIGIT
	}

	return cnpj_valid == docnum
}
