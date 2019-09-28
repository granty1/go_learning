package balance

import (
	"errors"
)

func init() {
	RegisterBalancer("round", &RoundRobin{})
}

type RoundRobin struct {
	currentIndex int64
}

func (r *RoundRobin) DoBanlance(instances []*Instance) (i *Instance, err error) {
	if len(instances) == 0 {
		err = errors.New("Instances can't be null.")
		return
	}

	len := len(instances)
	if r.currentIndex == (int64)(len) {
		r.currentIndex = 0
	}
	i = instances[r.currentIndex]
	r.currentIndex++
	return
}
