package Generalgo

import (
	"regexp"
	"strconv"
)

const (

	//PersonNameValidRegex é um padrão regex para nomes de pessoas válidos.
	PersonNameValidRegex string = `^(([A-Za-záàâãéèêëíïóôõöúüÿçñÁÀÂÃÉÈËÍÏÓÔÕÖÚÜÇÑ]{1,20} )*[A-Za-záàâãéèêëíïóôõöúüÿçñÁÀÂÃÉÈËÍÏÓÔÕÖÚÜÇÑ]{1,20}$)`

	//MaxNamesLength é o tamanho máximo para nomes em geral.
	MaxNamesLength uint8 = 40

	//DocDatabaseIDLength é o tamanho padrão de IDs em bancos baseados em documentos.
	DocDatabaseIDLength uint8 = 24

	//MinNamesDescriptionsLength é o tamanho mínimo para nomes de pessoas.
	MinNamesDescriptionsLength uint8 = 3

	//CpfRegex é um padrão regex para formato de número de cpf.
	CpfRegexNoMask = `^[0-9]{11}$`

	//CnpjRegex é um padrão regex para formato de número de cpf.
	CnpjRegexNoMask = `^[0-9]{14}$`
)

//StringMatch valida uma string com base em padrões regex. Retorna true se corresponde e false se não.
func StringMatch(value string, patterns ...string) bool {

	var rgx *regexp.Regexp
	var ret bool = true

	for i := 0; i < len(patterns); i++ {
		rgx = regexp.MustCompile(patterns[i])

		if !rgx.MatchString(value) {
			ret = false
			break
		}
	}
	return ret
}

//cpfIsValid valida se um número de cpf é valido.
func cpfIsValid(docnum string) bool {

	var cpf_valid string = docnum[:9]
	var base, digit uint8
	var sum uint16

CALCDIGIT:

	base, sum = uint8(len(cpf_valid)+1), 0

	for _, d := range cpf_valid {
		sum += uint16(uint8(d-'0') * base)

		base--
	}

	digit = uint8(11 - (sum % 11))

	if digit >= 10 {
		cpf_valid += "0"
	} else {
		cpf_valid += strconv.Itoa(int(digit))
	}

	if len(cpf_valid) < 11 {
		goto CALCDIGIT
	}

	return cpf_valid == docnum
}

// cnpjIsValid valida se um cnpj é válido.
func cnpjIsValid(docnum string) bool {

	var cnpj_valid string = docnum[:12]

	var base, digit uint8
	var sum uint16

CALCDIGIT:

	base, sum = 2, 0

	for j := len(cnpj_valid) - 1; j >= 0; j-- {

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

// CpfOrCnpjIsValid valida se um número de CNPJ ou CPF é válido.
func CpfOrCnpjIsValid(docnum string) bool {
	if !StringMatch(docnum, CnpjRegexNoMask) {
		if StringMatch(docnum, CpfRegexNoMask) {
			return cpfIsValid(docnum)
		}
		return false
	} else {
		return cnpjIsValid(docnum)
	}
}
