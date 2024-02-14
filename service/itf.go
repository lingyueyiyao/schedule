package service

type AssignItf interface {
	GetSerialNumber(BathType, int) int
	DelSerialNumber(BathType, int) error
	ShareRoom(...int) error
	CallSerialNumber(int) error
	ChangeSerialNumAward(BathType, int, int) error
	ExchangeBathType(BathType, BathType, int) (int, error)
}
