package main

import (
    "fmt"
    "net"
    "net/rpc"
)

// Estrutura para representar um aluno
type Aluno struct {
    Nome string
    Nota float64
}

// Estrutura para o servidor
type Notas struct {
    alunos []Aluno
}

// Método para inicializar a lista de alunos no servidor
func (s *Notas) inicializar() {
    s.alunos = []Aluno{
		{"Alexandre", 9.5},
		{"Barbara",   8.5},
		{"Joao",      6.5},
		{"Maria",     9.0},
		{"Paulo",    10.0},
		{"Pedro",     7.0},
	}
}

// Método remoto que retorna a nota de um aluno dado o seu nome
func (s *Notas) ObtemNota(nome string, nota *float64) error {
    for _, aluno := range s.alunos {
        if aluno.Nome == nome {
            *nota = aluno.Nota
            return nil
        }
    }
    return fmt.Errorf("Aluno %s não encontrado", nome)
}

func (s *Notas) Fatorial(n int, res *int) error {
    
    if(n<0){
        return fmt.Errorf("Numero invalido")
    }
    
    *res = 50;

    return nil 
}

func main() {
    porta := 8973
    servidor := new(Notas)
    servidor.inicializar()

    rpc.Register(servidor)
    l, err := net.Listen("tcp", fmt.Sprintf(":%d", porta))
    if err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
        return
    }

    for {
        fmt.Println("Servidor aguardando conexões na porta", porta)
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Erro ao aceitar conexão:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
