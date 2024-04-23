package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// canales
	// canla := make(chan int)
	// canla <- 15
	// valor := <- canla // recibe datos del canal

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch) // aplicar la concurrencia con go run time
	}

	//time.Sleep(1 * time.Second) // esto es impreciso
	// canales: conductos donde se usa para coordinar los goroutines

	// imprimir de manera correcta
	for i := 0; i < len(apis); i++ {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomo %v segundos!\n", elapsed.Seconds())

}

func checkAPI(api string, ch chan string) {

	if _, err := http.Get(api); err != nil {
		// fmt.Printf("ERROR: %s ! esta caido!\n", api)
		ch <- fmt.Sprintf("ERROR: %s ! esta caido!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS %s esta en funcionamiento\n", api)
}
