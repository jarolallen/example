package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ip := net.ParseIP("8.8.8.8") // Replace with the IP you want to ping

	conn, err := net.DialTimeout("ip4:icmp", ip.String(), 5*time.Second) //Pings the ip address with ICMP and a 5 second time out
	if err != nil {                                                      //If the errors are not non-existent
		fmt.Println("Error pinging:", err) //Print the error
		return                             //Quit the program
	}
	defer conn.Close() //close the connection to the ip address

	// Send the ICMP echo request
	_, err = conn.Write([]byte("PING"))
	if err != nil {
		fmt.Println("Error sending ping:", err)
		return
	}

	// Measure Round Trip Time
	start := time.Now()
	_, err = conn.Read(make([]byte, 1500))
	if err != nil {
		fmt.Println("Error receiving ping reply:", err)
		return
	}
	rtt := time.Since(start)

	fmt.Printf("Ping successful! RTT: %v\n", rtt) //Print success and ping time
}
