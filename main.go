package main

func main() {
	go func() {
		sum := 0
		for i := 0; i < 8; i++ {
			sum += i
			//time.Sleep(1 * time.Second)
		}
		println(sum)
	}()

	//println("Now goroutine = ", runtime.NumGoroutine())
	//time.Sleep(10 * time.Second)
}
