package user

type WhitelistInclusionDTO struct {
	EmployeeId string `json:"employeeId"`
}

type UserCreationDTO struct {
	Name       string `json:"name"`
	EmployeeId string `json:"employeeId"`
	Password   string `json:"password"`
}
