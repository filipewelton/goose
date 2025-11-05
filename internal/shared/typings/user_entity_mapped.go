package typings

type UserEntityMapped struct {
	ID         string `json:"id"`
	CardNumber string `json:"cardNumber"`
	Name       string `json:"name"`
	Supervisor bool   `json:"supervisor"`
}
