package main


type EmployeeNext struct {
	Name,Position string
	AnnualIncome float64
}

const (
	Developer=iota
	Manager
)

func NewEmployeeNext(role int) *EmployeeNext {
	switch role {
	case Developer:
		return &EmployeeNext{}
	default:
		panic("error role not defined")
	}
}

func main() {

}
