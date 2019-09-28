package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/granty1/go_learning/balance/balance"
)

var wait sync.WaitGroup

type GrantBalancer struct {
}

func (g *GrantBalancer) DoBanlance(instances []*balance.Instance) (i *balance.Instance, err error) {
	return instances[0], nil
}

func main() {
	var balancer = "random"
	if len(os.Args) > 0 {
		balancer = os.Args[1]
	}

	balance.RegisterBalancer("grant", &GrantBalancer{})

	fmt.Printf("choose %s :\n", balancer)

Label:
	for {
		instance, err := balance.GetBalance(balancer, getInstances())
		if err != nil {
			fmt.Println("Error params.")
			fmt.Println("\t-random\tchoose balancer with random.")
			fmt.Println("\t-round\tchoose balancer with round.")
			break Label
		}
		fmt.Println(instance)
		time.Sleep(1 * time.Second)
	}
	// wait.Wait()
	//println("Process over :)")

}

func ExecuteBalance(balancer balance.Balancer, instances []*balance.Instance, wait *sync.WaitGroup) {
	// by round
	for i := 0; i < 20; i++ {
		i, err := balancer.DoBanlance(instances)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		v := reflect.ValueOf(balancer)
		fmt.Printf("By %s => %v\n", v.Type(), i)
	}
	wait.Done()
}

func getInstances() []*balance.Instance {
	var instances []*balance.Instance
	for i := 0; i < 20; i++ {
		instance := &balance.Instance{
			Host: fmt.Sprintf("192.168.1.%d", i),
			Port: i,
		}
		instances = append(instances, instance)
	}
	return instances
}
