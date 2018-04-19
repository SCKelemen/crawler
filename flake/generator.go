package flake

type Generator struct {
	workerId  int
	processId int
}

func (g Generator) Generate() {

}

func New(workerId int, processId int) Generator {
	gen := Generator{workerId: workerId}
	return gen
}
