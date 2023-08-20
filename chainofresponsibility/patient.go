package chainofresponsibility

type Patient struct {
	Name string
	isRegistered bool
	isDoctorChecked bool
	isMedicineProvied bool
	isPaid bool
}