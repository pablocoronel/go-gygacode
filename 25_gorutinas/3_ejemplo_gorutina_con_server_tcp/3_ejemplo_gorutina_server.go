package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// crea el listener para escuchar conexiones
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err) //por ejemplo conexion abortada
			continue
		}

		go manejarConexion(conn) //meneja una conexion a la vez
		// siendo una gorutina, el serve no queda bloqueado en el bucle infinito "manejarConexion", ya que continua la ejecucion del programa y por lo tanto puede aceptar una nueva conexion
	}
}

func manejarConexion(c net.Conn) {
	defer c.Close() //cierra la conexion al final

	for {
		_, err := io.WriteString(c, time.Now().Format("15:05:07\n\r"))

		if err != nil {
			return //ejemplo si el cliente se desconecta
		}

		time.Sleep(1 * time.Second)
	}
}
