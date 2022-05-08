package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"syscall"
	"strings"

	"golang.org/x/term"
)

type UserCredentials struct {
	USERNAME  string `json:"USERNAME"`
	PASSWORD  string `json:"PASSWORD"`
}

type AuthCred struct {
	AuthParameters  struct `json:"AuthParameters"`
	AuthFlow        string `json:"AuthFlow"`
	ClientId        string `json:"ClientId"`
}

// Todo struct to map the response
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	pOne := os.Args[1]

	if pOne == "--help" || pOne == "-h" {
		fmt.Println("Flyawayhub connector, version 0.1")
		fmt.Println("Usage:  flyawayhub-connector [long option] [option] ...")
		fmt.Println("GNU long options:")
		fmt.Println("        --create-reservations, -c")
		fmt.Println("        --help, -h")
		fmt.Println("        --logout, -o")
		fmt.Println("        --login, -l [--username, -u] [--password, -p]")
		fmt.Println("        --reservations, -r")
		fmt.Println("        --schedule, -s [date] (empty will return present day)")
		fmt.Println("        --version, -v")
		fmt.Println("flyawayhub-connector home page: <https://www.github.com/jnbdz/flyawayhub-connector>")
	}

	if pOne == "--login" || pOne == "-l" {
		login()
	}

	if pOne == "--version" || pOne == "-v" {
		fmt.Println("Flyawayhub-connector, version 0.1")
	}

	if pOne == "--reservations" || pOne == "-r" {
		getReservations()
	}

	if pOne == "--schedule" || pOne == "-s" {
		getSchedule()
	}

	//get()
	//post()
	//put()
	//delete()
}

func authz() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal( err )
	}

	autzFileName := ".ConnectorAutzFlyawayhub"
	autzFilePath := dirname + autzFileName

	if _, err := os.Stat(autzFilePath); err == nil {
		content, err := ioutil.ReadFile(autzFilePath)
		bearer := string(content)
	} else if os.IsNotExist(err) {
		bearer := login()
		s := []byte(bearer)
		ioutil.WriteFile(autzFilePath, s, 0600)
	} else {
		log.Fatal( err )
	}

	return bearer
}

func login() {

	pTwo := os.Args[2]
	pThree := os.Args[3]
	if pTwo == "--username" || pTwo == "-u" {
		if pThree == "--password" || pThree == "-p" {
			
		}
	} else if pTwo == "--password" || pTwo == "-p" {
		if pThree == "--username" || pThree == "-u" {
			
		}
	} else {
		username, password, _ := credentials()
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')

	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)

	authCred := AuthCred{
			"AuthParameters",
			{},
			"AuthFlow",
			"ClientId"
		}
	jsonReq, err := json.Marshal(todo)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("user-agent", "golang application")
	req.Header.Add("foo", "bar1")
	req.Header.Add("foo", "bar2")

	response, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}

	defer response.Body.Close()

	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)


	resp, err := http.Post("")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)

}

/*func getReservations() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)

}

func get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}

func post() {
	fmt.Println("2. Performing Http Post...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}

func put() {
	fmt.Println("3. Performing Http Put...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct:\n%+v\n", todoStruct)
}

func delete() {
	fmt.Println("4. Performing Http Delete...")
	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonReq))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
}*/
