package main

import (
    "fmt"
    "net/rpc"
    "os"
    "strconv"
)

func main() {

    if len(os.Args) != 3{
        fmt.Println("Uso:", os.Args[0], " <maquina> <numero>")
        return
    }

    porta := 8973
    maquina := os.Args[1]
    numero, err := strconv.Atoi(os.Args[2])

    client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        return
    }

    var resultado int
    err = client.Call("Notas.Fatorial", numero, &resultado)
    if err != nil {
        fmt.Println("Erro ao obter o fatorial:", err)
    } else {
        fmt.Printf("Fatorial de: %d é %d\n", numero, resultado)
    }
}
