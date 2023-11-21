package main

import (
	"fmt"
	"runtime"
	"time"
)

func isPrime(p int) bool {
	if p%2 == 0 {
		return false
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func contaPrimosSeq(slice []int) int {
	count := 0
	for _, v := range slice {
		if isPrime(v) {
			count++
		}
	}
	return count
}

func contaPrimosConc(slice []int) int {
	count := 0
	ch := make(chan int)
	for _, v := range slice {
		go func(v int) {
			if isPrime(v) {
				ch <- 1
			} else {
				ch <- 0
			}
		}(v)
	}
	for i := 0; i < len(slice); i++ {
		count += <-ch
	}
	return count
}

// repetir para diferentes tamanhos de primos e para diferentes numeros de processadores
func main() {
	fmt.Println("------ conta primos de tamanho 6 -------")
	slice := []int{383376390724197361, 882611655919772761, 533290385325847007, 17969611178168479, 903013501582628521, 541906710014517121,
		281512690206248899, 403936627075987639, 775148726422474717, 942319117335957539}
	// todos no slice sao primos assim, a funcao acha primo demora a computar
	start := time.Now()
	p := contaPrimosSeq(slice)
	fmt.Println(" -> sequencial ------ secs: ", time.Since(start).Seconds())
	fmt.Println(" ------ n primos : ", p)

	runtime.GOMAXPROCS(1)
	start1 := time.Now()
	p = contaPrimosConc(slice)
	fmt.Println(" -> concorrente ------ secs: ", time.Since(start1).Seconds())
	fmt.Println(" ------ n primos : ", p)

	runtime.GOMAXPROCS(2)
	start1 = time.Now()
	p = contaPrimosConc(slice)
	fmt.Println(" -> concorrente ------ secs: ", time.Since(start1).Seconds())
	fmt.Println(" ------ n primos : ", p)

	runtime.GOMAXPROCS(4)
	start1 = time.Now()
	p = contaPrimosConc(slice)
	fmt.Println(" -> concorrente ------ secs: ", time.Since(start1).Seconds())
	fmt.Println(" ------ n primos : ", p)

	runtime.GOMAXPROCS(8)
	start1 = time.Now()
	p = contaPrimosConc(slice)
	fmt.Println(" -> concorrente ------ secs: ", time.Since(start1).Seconds())
	fmt.Println(" ------ n primos : ", p)

	runtime.GOMAXPROCS(16)
	start1 = time.Now()
	p = contaPrimosConc(slice)
	fmt.Println(" -> concorrente ------ secs: ", time.Since(start1).Seconds())
	fmt.Println(" ------ n primos : ", p)
}
