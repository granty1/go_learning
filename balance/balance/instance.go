package balance

import "fmt"

type Instance struct {
	Host string
	Port int
}

func (i *Instance) String() string {
	return fmt.Sprintf("%s : %d", i.Host, i.Port)
}
