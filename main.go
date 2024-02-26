package main

import "github.com/pedromotita/tcp-server/cmd"

// "github.com/pedromotita/tcp-server/pkg/server"

func main() {
	cmd.Execute()

	// host := "localhost"
	// port := "9002"

	// s := server.New(host, port)
	// s.Run()

	// log.Println("starting client")

	// c := client.New()
	// conn, err := c.Connect(host, port)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	//
	// client.Write(conn, "Hi from client\n")

}
