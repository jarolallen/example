package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ip := net.ParseIP("8.8.8.8") // Takes a string and turns it into a net.IP object

	conn, err := net.DialTimeout("ip4:icmp", ip.String(), 5*time.Second) //Pings the ip address with ICMP and a 5 second time out
	if err != nil {                                                      //If the errors are not non-existent
		fmt.Println("Error pinging:", err) //Print the error
		return                             //Quit the program
	}
	defer conn.Close() //close the connection to the ip address

	fmt.Println("Ping successful!") //print success
}
