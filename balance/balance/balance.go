package balance

type Balancer interface {
	DoBanlance(instances []*Instance) (i *Instance, err error)
}
