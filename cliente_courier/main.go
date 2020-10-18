package main

import (
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"io"

	"github.com/dnorambu/pruebas/courier"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error, error: ",e)
	}
}

func main() {
	var conn *grpc.ClientConn
	//Para realizar pruebas locales
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	// Descomentar para testear en MV
	//conn, err := grpc.Dial("10.10.28.141:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := courier.NewCourierServiceClient(conn)

	recordFile, err := os.Open("arch/pymes.csv") //1. Abrir el csv de ejemplo
	if err != nil {
		fmt.Println("Error al abrir el archivo: ",err)
		return
	}                             // 2. Chequear  si hubo error
	reader := csv.NewReader(recordFile)     // 3. Iniciar el reader
	//records, _ := reader.ReadAll()          // 4. Para tener todos los registros

	for i := 0; ; i++ {
		registro, err := reader.Read()
		if err == io.EOF {
			fmt.Println("Fin de archivo ",err)
			break
		}else if err != nil {
			fmt.Println("Error, error: ",err)
			return
		}
		
		valorTemp, err := strconv.Atoi(registro[2])
		prioridadTemp, err := strconv.Atoi(registro[5])
		
		orden := courier.OrdenPyme{
			Id: registro[0],
			Producto: registro[1],
			Valor: int64(valorTemp),
			Tienda: registro[3],
			Destino: registro[4],
			Prioridad: int64(prioridadTemp),
		}
		cod, err := c.CodigoPyme(context.Background(), &orden)
		if err != nil {
			fmt.Println("Error, error: ",err)
			return
		}
		log.Println("ID producto: ", registro[0], "Seguimiento: ", cod.Cod)
	}
	if recordFile.Close() != nil{
		fmt.Println("Error al cerrar")
		return
	}
	//Para testear el Hello
	message := courier.Mensaje{
		Body: "Hello from the client!",
	}

	response, err := c.Hello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}