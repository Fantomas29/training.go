package main

import "fmt"

type ErrorCode int

const (
	ERR_CODE_OK ErrorCode = iota
	ERR_CODE_NOT_FOUND
	ERR_CODE_LOCKED
	ERR_CODE_GENERIC
)

func (ec ErrorCode) String() string {
	return [...]string{
		"ok",
		"not found",
		"locked",
		"generic",
	}[ec]
}

func (ec ErrorCode) IsCritical() bool {
	return ec == ERR_CODE_LOCKED || ec == ERR_CODE_NOT_FOUND
}

func IsValid(ec ErrorCode) bool {
	// a implementer si on veut tester que la valeur envoye est bien valide
	return true
}

func printErrCode(c ErrorCode) {
	fmt.Printf("code=%d, critical=%v, detail=%v\n",
		c,
		c.IsCritical(),
		c.String(),
	)
}

func main() {
	//code := ERR_CODE_LOCKED

	// si on envoie printErrCode(10) compile mais erreur à l'execution
	printErrCode(3)
}
