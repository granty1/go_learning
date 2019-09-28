package balance

import (
	"errors"
	"fmt"
)

type BalanceManager struct {
	allBalancer map[string]Balancer
}

var (
	manager = BalanceManager{
		allBalancer: make(map[string]Balancer),
	}
)

func (m *BalanceManager) registerBalancer(name string, b Balancer) {
	m.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) error {
	if b == nil {
		return errors.New("Balancer can't be nil.")
	}
	if !(len(name) > 0) {
		return errors.New("Key can't be nil.")
	}
	manager.registerBalancer(name, b)
	return nil
}

func GetBalance(name string, instances []*Instance) (instance *Instance, err error) {
	balancer, ok := manager.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Can't find %s balancer.", name)
		return
	}
	instance, err = balancer.DoBanlance(instances)
	return
}
