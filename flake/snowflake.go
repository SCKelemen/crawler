package flake

const (
	TimeBitLength      = 39
	SequenceBitLength  = 8
	MachineIDBitLength = 63 - TimeBitLength - SequenceBitLength // perhaps we could use a mac address here
)

type Settings struct {
}
