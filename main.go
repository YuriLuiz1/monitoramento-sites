package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	timeMonitorings    = 3
	sleepingMonitoring = 5
)

func main() {
	showIntroduction()

	for {
		showMenu()

		command := digitUser()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
		case 3:
			clearLogs()
		case 0:
			fmt.Println("Exiting Program...")
			os.Exit(0)
		default:
			fmt.Println("This command is unknown")
			os.Exit(-1)
		}
	}
}

func getNamePc() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Don't possible get hostname")
		return "Unknown"
	}
	return hostname
}

func showIntroduction() {
	name := getNamePc()
	var version float32 = 1.1
	fmt.Println("Welcome sr. ", name)
	fmt.Println("Version this program is ", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("3 - Clear Logs")
	fmt.Println("0 - Exit Program")
}

func digitUser() int {
	var capitDigit int
	fmt.Scan(&capitDigit)
	fmt.Println("The digit choose is: ", capitDigit)

	return capitDigit
}

func startMonitoring() {
	fmt.Println("")

	fmt.Println("Monitoring...")

	fmt.Println("")

	sites := readSitesArquives()

	for i := 0; i < timeMonitorings; i++ {
		for i, site := range sites {
			fmt.Println("Monitoring position", i, "this site is", site)
			testingSites(site)
		}
		fmt.Println("")
		time.Sleep(sleepingMonitoring * time.Second)
	}

	fmt.Println("")
}

func testingSites(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Error in program!", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site ", site, "Uploaded with successful")
		registerLogs(site, true)
	} else {
		fmt.Println("Site ", site, "Have problems! Status code: ", resp.StatusCode)
		registerLogs(site, false)
	}
}

func readSitesArquives() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error in program!", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registerLogs(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Have problem in program :", err)
	}

	arquivo.WriteString(
		time.Now().
			Format("02/01/2006 15:04:05") +
			" - " + site + " - Online: " + strconv.FormatBool(
			status,
		) + "\n",
	)

	arquivo.Close()
}

func printLogs() {
	arquivo, err := os.ReadFile("logs.txt")

	fmt.Println("")

	fmt.Println("Showing Logs...")

	fmt.Println("")

	if err != nil {
		fmt.Println("Have problem in program :", err)
	}

	fmt.Println(string(arquivo))
}

func clearLogs() {
	fmt.Println("")

	fmt.Println("Limpando Logs")

	fmt.Println("")

	logsLimpo := "logs.txt"
	arquivo, err := os.OpenFile(logsLimpo, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Have problem in program: ", err)
	}

	arquivo.Close()
}
