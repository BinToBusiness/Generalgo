package generalgo

//NegativeNumberError é um erro personalizado para valores menos que zero.
type NegativeNumberError struct {
	Name  string
	Value string
}

//NumberLessThanOneError é um erro personalizado para valores menos que 1.
type NumberLessThanOneError struct {
	Name  string
	Value string
}

//RangeNumberError é um erro personalizado para números que não estão entre um range.
type RangeNumberError struct {
	Name  string
	Value string
}

//EmptyError é um erro personalizado para valores vazios
type EmptyError struct {
	Name string
}

//EnumElementNotExistError é um erro personalizado elementos de enumeração que não existem.
type EnumElementNotExistError struct {
	Name  string
	Value string
}

//LengthError é um erro personalizado para valores que violam um tamanho definido.
type LengthError struct {
	Name  string
	Value string
}

//MinimumLengthError é um erro personalizado para valores que violam um tamanho mínimo.
type MinimumLengthError struct {
	Name  string
	Value string
}

//MaximumLengthError é um erro personalizado para valores que violam um tamanho mínimo.
type MaximumLengthError struct {
	Name  string
	Value string
}

//NotMatchError é um erro personalizado para valores que não correspodem a um padrão.
type NotMatchError struct {
	Name  string
	Value string
}

//DateTimeError é um erro personalizado para data/hora invalidas
type DateTimeError struct {
	Name  string
	Value string
}

//CollectionElementNotFoundError é um erro personalizado para elemenetos que não fazem parte de uma coleção.
type CollectionElementNotFoundError struct {
	Name  string
	Value string
}

/*
* Error functions
 */

//Error função retorna a descrição do erro
func (nne *NegativeNumberError) Error() string {

	if len(nne.Name) == 0 {
		nne.Name = "The value"
	}
	return nne.Name + " cannot be less than zero. Value is " + nne.Value
}

//Error função retorna a descrição do erro
func (nltoe *NumberLessThanOneError) Error() string {

	if len(nltoe.Name) == 0 {
		nltoe.Name = "The value"
	}

	return nltoe.Name + " cannot be less than one. Value is " + nltoe.Value
}

//Error função retorna a descrição do erro
func (rne *RangeNumberError) Error() string {

	if len(rne.Name) == 0 {
		rne.Name = "The value"
	}

	return rne.Name + " is outside the defined range. Value is " + rne.Value
}

//Error função retorna a descrição do erro
func (ee *EmptyError) Error() string {

	if len(ee.Name) == 0 {
		ee.Name = "The value"
	}

	return ee.Name + " cannot be empty."
}

///Error função retorna a descrição do erro
func (eenee *EnumElementNotExistError) Error() string {

	if len(eenee.Name) == 0 {
		eenee.Name = "The value"
	}

	return eenee.Name + " is not present in the enumeration. Value is " + eenee.Value
}

//Error função retorna a descrição do erro
func (le *LengthError) Error() string {

	if len(le.Name) == 0 {
		le.Name = "The value"
	}

	return le.Name + " violates the stipulated length. Value is " + le.Value
}

//Error função retorna a descrição do erro
func (mle *MinimumLengthError) Error() string {

	if len(mle.Name) == 0 {
		mle.Name = "The value"
	}

	return mle.Name + " violates the stipulated minimum length. Value is " + mle.Value
}

//Error função retorna a descrição do erro
func (mle *MaximumLengthError) Error() string {

	if len(mle.Name) == 0 {
		mle.Name = "The value"
	}

	return mle.Name + " violates the stipulated maximum length. Value is " + mle.Value
}

//Error função retorna a descrição do erro
func (nme *NotMatchError) Error() string {

	if len(nme.Name) == 0 {
		nme.Name = "The value"
	}

	return nme.Name + " not match the stipulated pattern. Value is " + nme.Value
}

//Error função retorna a descrição do erro
func (dte *DateTimeError) Error() string {

	if len(dte.Name) == 0 {
		dte.Name = "The value"
	}

	return dte.Name + " is a date/time invalid. Value is " + dte.Value
}

//Error função retorna a descrição do erro
func (cefe *CollectionElementNotFoundError) Error() string {

	if len(cefe.Name) == 0 {
		cefe.Name = "The value"
	}

	return cefe.Name + " is not present in collection. Value is " + cefe.Value
}
