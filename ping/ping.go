package main

import (
	"bufio"
	"flag"
	"fmt"
	"os/exec"
)

var ip string

func main() {
	flag.StringVar(&ip,
		"ip",
		"8.8.8.8", //default ip address
		"IP Address",
	) //look for fed ip address
	flag.Parse()                    //parse command line
	cmd := exec.Command("ping", ip) //strings sent to cmd
	out, _ := cmd.StdoutPipe()      //output saved
	cmd.Start()                     //launches cmd

	scan := bufio.NewScanner(out) //scan output of cmd
	for scan.Scan() {
		m := scan.Text()
		fmt.Print(m) //print output of cmd
	}
	cmd.Wait()
}
