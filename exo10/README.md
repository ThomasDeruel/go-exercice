# Go Exercice 10 - Concurrence

1. `go func`: crée une goroutine => permet d'executer des fonctions en parallèle à partir du même espace d'adressage (processus) -> permet de paralliser des tâches.

```go
go func (){
    fmt.Println("Hey")
}
```

>Attention à la synchronisation: la fonction `main()` n'attend pas que les fonctions parallèles se terminent.

2. `type sync.WaitGroups struct {}`: permet d'attendre que tous les goroutines se terminent. Fait appel à quelques méthodes:
    -   `Add(delta int)`: incrémente le counteur de 1 
    -   `Done()`: décrémente le compteur de 1
    -   `Wait()`: bloque jusqu'à que le compteur atteigne 0

Exemple:

```go
import (
    "fmt"
    "sync"
)

var wg sync.WaitGroups
var maFunction(wg *sync.WaitGroups){
    // defer wg.Done() ou...plus bas
    //...
    //...
    // wg.Done()
}
func main(){
    wg.Add(1)
    go maFunction(&wg)
    wg.Wait()
    fmt.Println("Fin du programme")
}
```