package main

import (
    "fmt"
    "net/rpc"
    "os"
)

type Aluno struct {
    Nome string
    Nota float64
}

func main() {

    if len(os.Args) != 3{
        fmt.Println("Uso:", os.Args[0], " <maquina> <nome_do_aluno>")
        return
    }

    porta := 8973
    maquina := os.Args[1]
    nome := os.Args[2]

    client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%d", maquina, porta))
    if err != nil {
        fmt.Println("Erro ao conectar ao servidor:", err)
        return
    }

    var nota float64
    err = client.Call("Notas.ObtemNota", nome, &nota)
    if err != nil {
        fmt.Println("Erro ao obter nota:", err)
    } else {
        fmt.Printf("Nome: %s\n", nome)
        fmt.Printf("Nota: %.2f\n", nota)
    }

    var aluno Aluno
    err = client.Call("Notas.ObtemAluno", nome, &aluno)
    if err != nil {
        fmt.Println("Erro ao obter nota:", err)
    } else {
        fmt.Printf("Aluno: %s\n", aluno.Nome)
    }
}
