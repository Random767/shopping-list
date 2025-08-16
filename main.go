package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Data struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	IsMarked bool   `json:"isMarked"`
	Priority int    `json:"priority"`
}

var shoppingList []Data

func main() {
	fmt.Println("Olá! Seja bem-vindo a lista de compras :D")
	fmt.Println("Dica: utilize help para ajuda")
	drawn()
}

func drawn() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		if !scanner.Scan() {
			fmt.Println("\nSaindo...");
			break;
		}
		var input []string = getArgs(scanner.Text())
		var command string = input[0]
		var args string
		for i := 1; i < len(input); i++ {
			if i == len(input)-1 {
				args = args + input[i]
			} else {
				args = args + input[i] + " "
			}
		}

		if command == "help" {
			var help string = `Bem vindo ao menu de ajuda, abaixo estão todos os comandos disponíveis:
 help: Veja esse menu de ajuda
 list: Retorna uma lista com todos os produtos
 add: Adicione um produto
 remove: Remove um produto
 mark: Marca o produto como comprado
 unmark: Desmarca um produto`
			fmt.Println(help)
		} else if command == "add" {
			if args == "" {
				fmt.Println("Forma de uso: add <produto>")
			} else {
				fmt.Println(createProduct(args))
			}
		} else if command == "list" {
			if len(shoppingList) == 0 {
				fmt.Println("Nenhum produto existente")
			}
			var isMarkedQuantity int = 0
			var isUnmarkedQuantity int = 0
			for i := 0; i < len(shoppingList); i++ {
				if shoppingList[i].IsMarked {
					fmt.Printf("%d. \x1b[9m%v\x1b[0m\n", i+1, shoppingList[i].Name)
					isMarkedQuantity += 1
				} else {
					fmt.Printf("%d. %v\n", i+1, shoppingList[i].Name)
				}
			}
            if len(shoppingList) == 1 {
			    isUnmarkedQuantity = len(shoppingList) - isMarkedQuantity
			    fmt.Printf("[Stat] Produtos marcados: %d\n", isMarkedQuantity)
			    fmt.Printf("[Stat] Produtos nao marcados: %d\n", isUnmarkedQuantity)
		    }
        } else if command == "remove" {
			if len(shoppingList) == 0 {
				fmt.Println("Nenhum produto foi cadastrado")
			} else {
				fmt.Println(removeProduct(args))
			}
		} else if command == "mark" {
			if args == "" {
				fmt.Println("Forma de uso: mark <produto>")
			} else {
				fmt.Println(markProduct(args))
			}
		} else if command == "unmark" {
			if args == "" {
				fmt.Println("Forma de uso: unmark <produto>")
			} else {
				fmt.Println(unmarkProduct(args))
			}
		} else if command == "stat" {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			fmt.Printf("Memória utilizada: %v bytes\n", m.Alloc)
        } else if command == "exit" {
            fmt.Println("Saindo...")
            os.Exit(0)
        } else {
            if command != "" {
                fmt.Printf("O comando \"%s\" não existe.\n", command)
            }
        }
	}
}

func getArgs(args string) []string {
	split := strings.Split(args, " ")
	return split
}

func findIndex(arr []Data, target Data) int {
	for i, v := range arr {
		if v.Name == target.Name {
			return i
		}
	}
	return -1
}

func createProduct(name string) string {
    target := Data{Name: name}
    if findIndex(shoppingList, target) != -1 {
        return fmt.Sprintf("Produto \"%s\" já existe", name)
    }
	p := Data{Name: name}
	shoppingList = append(shoppingList, p)
	return fmt.Sprintf("Produto \"%s\" adicionado", name)
}

func removeProduct(name string) string {
	target := Data{Name: name}
	num, err := strconv.Atoi(name)
	if err != nil {
		index := findIndex(shoppingList, target)
		if index == -1 {
			return fmt.Sprintf("Produto \"%s\" não foi encontrado", name)
		} else {
			shoppingList = append(shoppingList[:index], shoppingList[index+1:]...)
			return fmt.Sprintf("Produto \"%s\" removido.", name)
		}
	} else {
		if num > len(shoppingList) {
			return fmt.Sprintf("Produto com o id %d não foi encontrado", num)
		} else {
			productName := shoppingList[num-1].Name
			shoppingList = append(shoppingList[:num-1], shoppingList[num:]...)
			return fmt.Sprintf("Produto \"%v\" removido.", productName)
		}
	}
}

func markProduct(name string) string {
	target := Data{Name: name}
	num, err := strconv.Atoi(name)
	if err != nil {
		index := findIndex(shoppingList, target)
		if index == -1 {
			return fmt.Sprintf("Produto \"%s\" não foi encontrado", name)
		} else {
			shoppingList[index].IsMarked = true
			return fmt.Sprintf("Produto \"%s\" foi marcado", name)
		}
	} else {
		if num > len(shoppingList) {
			return fmt.Sprintf("Produto com o id %d não foi encontrado", num)
		} else {
			shoppingList[num-1].IsMarked = true
			return fmt.Sprintf("Produto \"%v\" foi marcado", shoppingList[num-1].Name)
		}
	}
}

func unmarkProduct(name string) string {
	target := Data{Name: name}
	num, err := strconv.Atoi(name)
	if err != nil {
		index := findIndex(shoppingList, target)
		if index == -1 {
			return fmt.Sprintf("Produto \"%s\" não foi encontrado", name)
		} else {
			shoppingList[index].IsMarked = false
			return fmt.Sprintf("Produto \"%s\" foi desmarcado", name)
		}
	} else {
		if num > len(shoppingList) {
			return fmt.Sprintf("Produto com o id %d não foi encontrado", num)
		} else {
			shoppingList[num-1].IsMarked = false
			return fmt.Sprintf("Produto \"%v\" foi desmarcado", shoppingList[num-1].Name)
		}
	}
}
