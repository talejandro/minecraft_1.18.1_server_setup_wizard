package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	//Limpiar consola en linux.
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func verification() {
	//Verificacion del archivo
	clear()
	fmt.Println("Checking EULA file")
	time.Sleep(1 * time.Second)
	fmt.Println("Creating and accepting EULA, please wait...")
	var ruta string
	ruta = "eula.txt"

	var exis bool
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		exis = false
	} else {
		exis = true
	}

	if exis {
		//modificamos el false por el true dentro del archivo eula.txt
		aceptar := []byte("eula=true")
		err := ioutil.WriteFile(ruta, aceptar, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		//corremos el archivo del servidor para crear el eula.txt
		cmd := exec.Command("java", "-Xmx1024M", "-Xms1024M", "-jar", "server.jar")
		cmd.Run()
		aceptar := []byte("eula=true")
		err := ioutil.WriteFile(ruta, aceptar, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	time.Sleep(1 * time.Second)
}

func setup() {

	//Cambiar Archivo Server.Propieties
}

func main() {
	welcome()
	var menu string

	for menu != "0" {
		clear()
		fmt.Println("Main menu")
		fmt.Println("1. Creation of necessary files")
		fmt.Println("2. Setup")
		fmt.Println("3. Start Server")
		fmt.Println("0. Exit")
		fmt.Println("Select an option:")
		fmt.Scan(&menu)

		switch menu {
		case "1":
			verification()
			fmt.Println("Task completed successfully! Going back to the main menu")
			time.Sleep(1 * time.Second)
			clear()
		case "2":
			clear()
			setup()
			fmt.Println("Setup completed successfully! Going back to the main menu")
			time.Sleep(1 * time.Second)
			clear()
		case "3":
		case "0":
			clear()
			fmt.Println("Bye! :)")
			time.Sleep(1 * time.Second)
		default:
			clear()
			fmt.Println(menu, "is't an option, try again!")
			time.Sleep(1 * time.Second)
		}
	}
}
