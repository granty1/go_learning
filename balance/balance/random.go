package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

type RandomBalance struct {
}

func (r *RandomBalance) DoBanlance(instances []*Instance) (i *Instance, err error) {

	if len(instances) == 0 {
		err = errors.New("Instances can't be null.")
		return
	}

	len := len(instances)
	index := rand.Intn(len)

	i = instances[index]
	return
}
