package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

// RandomBalance is a struct
type RandomBalance struct {
}

// DoBalance to execute balance.
func (r *RandomBalance) DoBalance(instances []*Instance, key ...string) (i *Instance, err error) {

	if len(instances) == 0 {
		err = errors.New("Instances can't be null.")
		return
	}

	len := len(instances)
	index := rand.Intn(len)

	i = instances[index]
	return
}
