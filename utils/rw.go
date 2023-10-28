package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Rw() {
	// abrir o arquivo para leitura
	file, err := os.Open("exemplo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// criar um leitor bufio
	reader := bufio.NewReader(file)

	// ler o arquivo linha por linha
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}

	// abrir o arquivo para escrita
	file, err = os.Create("novo-arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// criar um escritor bufio
	writer := bufio.NewWriter(file)

	// escrever no arquivo
	_, err = writer.WriteString("Hello, world!\n")
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
