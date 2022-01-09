package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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

func lectura(linea, parametro string) {

	//Con esta funcion voy a una linea en especidico y la modifico, aca depende de las variables que defini arriba
	fileconfig := "server.properties"
	input, err := ioutil.ReadFile(fileconfig)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		if strings.Contains(line, linea) {
			lines[i] = parametro
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileconfig, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}

func setup() {
	//Cambiar Archivo Server.Propieties
	var setupmenu string

	for setupmenu != "0" {
		clear()
		fmt.Println("Setup")
		fmt.Println("1. Name")
		fmt.Println("2. Game mode")
		fmt.Println("3. Difficulty")
		fmt.Println("4. Premium mode")
		fmt.Println("5. Pvp")
		fmt.Println("6. Server port")
		fmt.Println("7. View Distance")
		fmt.Println("8. Simulation Distance")
		fmt.Println("9. Hardcore mode")
		fmt.Println("0. Menu")
		fmt.Printf("Insert an option :")
		fmt.Scanln(&setupmenu)

		switch setupmenu {
		case "1":
			clear()
			fmt.Println("Enter name server.")
			fmt.Println("Name: ")
			var nname string
			fmt.Scanln(&nname) //Aca hay que buscar la forma de que si se inserta un nombre con espacio se guarde en la variable.
			nname = "motd=" + nname
			lectura("motd=", nname)

		case "2":
			clear()
			fmt.Println("Enter gamemode.")
			fmt.Println("(spectator) (survival) (creative) (adventure) : ")
			var ngame string
			fmt.Scanln(&ngame)
			ngame = "gamemode=" + ngame
			lectura("gamemode=", ngame)
		case "3":
			clear()
			fmt.Println("Enter difficulty.")
			fmt.Println("(hard) (normal) (easy) : ")
			var ndiff string
			fmt.Scanln(&ndiff)
			ndiff = "difficulty=" + ndiff
			lectura("difficulty=", ndiff)

		case "4":
			clear()
			fmt.Println("Premium mode?")
			fmt.Println("(true) (false): ")
			var nprem string
			fmt.Scanln(&nprem)
			nprem = "online-mode=" + nprem
			lectura("online-mode=", nprem)

		case "5":
			clear()
			fmt.Println("Player vs Player.")
			fmt.Println("(true) (false): ")
			var npvp string
			fmt.Scanln(&npvp)
			npvp = "pvp=" + npvp
			lectura("pvp=", npvp)

		case "6":
			clear()
			fmt.Println("Enter port.")
			fmt.Println("(default: 25565):")
			var nport string
			fmt.Scanln(&nport)
			nport = "server-port=" + nport
			lectura("server-port=", nport)

		case "7":
			clear()
			fmt.Println("Enter view distance.")
			fmt.Println("(Chunks Number):")
			var nchun string
			fmt.Scanln(&nchun)
			nchun = "view-distance=" + nchun
			lectura("view-distance=", nchun)

		case "8":
			clear()
			fmt.Println("Enter simulation distance.")
			fmt.Println("(Chunks Number):")
			var nsimu string
			fmt.Scanln(&nsimu)
			nsimu = "simulation-distance=" + nsimu
			lectura("simulation-distance=", nsimu)

		case "9":
			clear()
			fmt.Println("Hardcore mode.")
			fmt.Println("(true) (false):")
			var nhard string
			fmt.Scanln(&nhard)
			nhard = "hardcore=" + nhard
			lectura("hardcore=", nhard)

		case "0":
		default:
			clear()
			fmt.Println(setupmenu, "is't an option, try again!")
			time.Sleep(2 * time.Second)
		}

	}
}

func server() {
	clear()
	fmt.Println("Run Server")
	time.Sleep(1 * time.Second)
	fmt.Println("Enter number in M (ex: 3072M): ")
	var cream string
	var xmx string
	var xms string
	fmt.Scan(&cream)
	xmx = "-Xmx" + cream + "M"
	xms = "-Xms" + cream + "M"
	fmt.Println("Attention!!!, you will enter on the server administration console.")
	clear()
	fmt.Println("To stop it press Crtl + C")
	time.Sleep(1 * time.Second)
	fmt.Println("-----Remember connect to ........-----")
	time.Sleep(1 * time.Second)
	fmt.Println("The server stars, in...")
	time.Sleep(1 * time.Second)
	fmt.Println("3...")
	time.Sleep(1 * time.Second)
	fmt.Println("2...")
	time.Sleep(1 * time.Second)
	fmt.Println("1...")

	cmd := exec.Command("java", xmx, xms, "-jar", "server.jar")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
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
			clear()
			fmt.Println("Task completed successfully! Going back to the main menu")
			time.Sleep(2 * time.Second)
			clear()
		case "2":
			clear()
			setup()
			clear()
			fmt.Println("Setup completed successfully! Going back to the main menu")
			time.Sleep(2 * time.Second)
			clear()
		case "3":
			server()
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
