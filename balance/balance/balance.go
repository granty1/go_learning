package balance

type Balancer interface {
	DoBalance(instances []*Instance, key ...string) (i *Instance, err error)
}
