package movimientos

import "fmt"

type CuentaBancaria struct {
	saldo int
}

func (c *CuentaBancaria) Cargo(monto int) {
	c.saldo += monto
	fmt.Println("la cuenta bancaria arroja un saldo de", c.saldo)
}
