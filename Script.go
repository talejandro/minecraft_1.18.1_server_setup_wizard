package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func welcome() {
	//Bienvenida al Script
	fmt.Println("Welcome to the Minecraft Server setup wizard")
	time.Sleep(1 * time.Second)
	fmt.Println("By Talejandro")
	time.Sleep(1 * time.Second)
	clear()
}

func clear() {
	//Limpiar consola
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func verification() {
	//Verificacion del archivo

}

func change() {
	//Cambiar Archivo Server.Propieties
}

func main() {
	welcome()

}
