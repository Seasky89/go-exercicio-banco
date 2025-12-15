package pagamentos

type Conta interface {
	Sacar(valor float64) (string, float64)
}

func PagarBoleto(conta Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}
