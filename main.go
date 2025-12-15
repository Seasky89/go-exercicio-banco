package main

import (
	"fmt"

	"github.com/Seasky89/banco/clientes"
	"github.com/Seasky89/banco/contas"
	"github.com/Seasky89/banco/pagamentos"
)

func main() {
	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor",
	}

	contaDoBruno := contas.ContaCorrente{
		Titular:       clienteBruno,
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}
	contaDoBruno.Depositar(300)
	fmt.Println(contaDoBruno)

	pagamentos.PagarBoleto(&contaDoBruno, 60)
	fmt.Println(contaDoBruno)

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis)

	pagamentos.PagarBoleto(&contaDoDenis, 55)
	fmt.Println(contaDoDenis)
}
