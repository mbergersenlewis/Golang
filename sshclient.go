// Matthew Bergersen Lewis
// May 29, 2020
// v1

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	//"encoding/json"
	"golang.org/x/crypto/ssh"
	// Uncomment to store output in variable
	//"bytes"
)

var hostList []string

// Read from file: hostfile.txt
func loginHosts() {
	hf, _ := os.Open("hostfile.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
		// scaffolding for troubleshooting
		//       fmt.Println(hostList)
	}
}

func main() {
	loginHosts()
	//scaffolding for troubleshooting
	//for _, host := range hostList {
	//		fmt.Println(host)

	username := ""
	password := ""
	//	hostname := ""
	//	port := "22"

	// SSH client config
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Non-production only
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Loop through and connect to hosts
	for _, host := range hostList {
		targethost := host
		port := "22"
		client, err := ssh.Dial("tcp", targethost+":"+port, config)
		time.Sleep(1)
		//client, err := ssh.Dial("tcp", hostname+":"+port, config)

		//Scaffolding for troubleshooting
		//fmt.Println(client)

		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		// Create sesssion
		sess, err := client.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		defer sess.Close()

		// StdinPipe for commands
		stdin, err := sess.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}

		// Uncomment to store output in variable
		//var b bytes.Buffer
		//sess.Stdout = &amp;b
		//sess.Stderr = &amp;b

		// Enable system stdout
		// Comment these if you uncomment to store in variable
		sess.Stdout = os.Stdout
		sess.Stderr = os.Stderr

		// Start remote shell
		err = sess.Shell()
		if err != nil {
			log.Fatal(err)
		}

		// send the commands
		commands := []string{
			//"show interface terse",
			"show arp no-resolve",
			"exit",
		}
		for _, cmd := range commands {
			_, err = fmt.Fprintf(stdin, "%s\n", cmd)
			if err != nil {
				log.Fatal(err)
			}
			//		sess.Wait()
			//		sess.Close()
		}

		// Wait for sess to finish
		err = sess.Wait()
		if err != nil {
			log.Fatal(err)
		}

		// Uncomment to store in variable
		//fmt.Println(b.String())

	}
}
© 2020 GitHub, Inc.
Terms
Privacy
Security
Status
Help
Contact GitHub
Pricing
API
Training
Blog
About
