package enum

type Status uint8

const (
	Created Status = iota
	Paid
	Pending
)

func (s Status) ToString() string {
	switch s {
	case Paid:
		return "paid"
	case Pending:
		return "pending"
	default:
		panic("unhandled default case")
	}

	return "unknown"
}
