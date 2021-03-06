package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/dnorambu/pruebas/courier"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	//Para testear en local
	//conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())

	//Para testear en MV
	conn, err := grpc.Dial("10.10.28.141:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := courier.NewCourierServiceClient(conn)

	var wg sync.WaitGroup
	var tiempo, tiempo2 int64
	fmt.Println("Ingrese el tiempo (en segundos) de espera del camion:")
	fmt.Scan(&tiempo)

	fmt.Println("Ingrese el tiempo (en segundos) de envio de paquetes:")
	fmt.Scan(&tiempo2)
	for {
		wg.Add(3)
		go cargaRetail1(c, &wg, tiempo, tiempo2)
		go cargaRetail2(c, &wg, tiempo, tiempo2)
		go cargaNormal(c, &wg, tiempo, tiempo2)
		wg.Wait()

		time.Sleep(time.Millisecond * 1500)
	}

}

func cargaRetail1(c courier.CourierServiceClient, wg *sync.WaitGroup, tiempo int64, tiempo2 int64) { // Función para camión 1
	defer wg.Done()
	queCamion := "Camion retail 1"
	message := courier.Mensaje{
		Body: "Necesito paquete.",
	}
	ret1, err := c.PedirRetail(context.Background(), &message) // Busca un retail en cola retail
	if err != nil {
		log.Fatalf("Error when calling PedirRetail: %s", err)
	}
	if ret1.Id != "" { // Caso en que si encontró retail
		time.Sleep(time.Duration(tiempo) * time.Second)            //Esperar el tiempo determinado
		ret2, err := c.PedirRetail(context.Background(), &message) // Ahora evalúa si hay un segundo retail en cola retail
		if err != nil {
			log.Fatalf("Error when calling PedirRetail: %s", err)
		}
		if ret2.Id != "" { // Encontró un segundo retail.

			if ret2.Valor > ret1.Valor { //Para ver cual intentar enviar primero por valor
				dospaquetes(c, ret2, ret1, queCamion, tiempo2)
			} else {
				dospaquetes(c, ret1, ret2, queCamion, tiempo2)
			}
			//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT RT
		} else if ret2.Id == "" { // No encontró un segundo retail.
			prio1, err := c.PedirPrioritario(context.Background(), &message) // Prueba entonces con prioritario.
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
			if prio1.Id != "" { // Si encuentra prioritario.
				if prio1.Valor > ret1.Valor { //Para ver cual intentar enviar primero por valor
					dospaquetes(c, prio1, ret1, queCamion, tiempo2)
				} else {
					dospaquetes(c, ret1, prio1, queCamion, tiempo2)
				}
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT PR
			} else if prio1.Id == "" { //No encuentra prioritario.
				unpaquete(c, ret1, queCamion, tiempo2)
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE RT
			}

		}
	} else if ret1.Id == "" { // No encontró un primer retail.
		prio1, err := c.PedirPrioritario(context.Background(), &message) // Prueba entonces buscar prioritario en cola prioritario.
		if err != nil {
			log.Fatalf("Error when calling PedirPrioritario: %s", err)
		}

		if prio1.Id != "" { // Encuentra prioritario.
			time.Sleep(time.Duration(tiempo) * time.Second)
			prio2, err := c.PedirPrioritario(context.Background(), &message) // Busca el segundo prioritario.
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
			if prio2.Id != "" { // Encuentra segundo prioritario.
				if prio2.Valor > prio1.Valor { //Para ver cual intentar enviar primero por valor
					dospaquetes(c, prio2, prio1, queCamion, tiempo2)
				} else {
					dospaquetes(c, prio1, prio2, queCamion, tiempo2)
				}
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR PR
			} else if prio2.Id == "" { // No encuentra un segundo prioritario.
				unpaquete(c, prio1, queCamion, tiempo2)
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE PR
			}
		}
	}
}

func cargaRetail2(c courier.CourierServiceClient, wg *sync.WaitGroup, tiempo, tiempo2 int64) { // Función para camión 2 (similar a la función del camión 1)
	defer wg.Done()
	queCamion := "Camion retail 2"
	message := courier.Mensaje{
		Body: "Necesito paquete.",
	}
	ret1, err := c.PedirRetail(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling PedirRetail: %s", err)
	}
	if ret1.Id != "" { //Encuentra retail
		time.Sleep(time.Duration(tiempo) * time.Second)
		ret2, err := c.PedirRetail(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling PedirRetail: %s", err)
		}
		if ret2.Id != "" {

			if ret2.Valor > ret1.Valor { //Para ver cual intentar enviar primero por valor
				dospaquetes(c, ret2, ret1, queCamion, tiempo2)
			} else {
				dospaquetes(c, ret1, ret2, queCamion, tiempo2)
			}

			//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT RT (recordar hacer comparacion de valores antes)

		} else if ret2.Id == "" {
			prio1, err := c.PedirPrioritario(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
			if prio1.Id != "" {

				if prio1.Valor > ret1.Valor { //Para ver cual intentar enviar primero por valor
					dospaquetes(c, prio1, ret1, queCamion, tiempo2)
				} else {
					dospaquetes(c, ret1, prio1, queCamion, tiempo2)
				}

				//TIRAR FUNCION YA QUE HAY 2 PAQUETES RT PR
			} else if prio1.Id == "" {
				unpaquete(c, ret1, queCamion, tiempo2)
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE RT
			}

		}
	} else if ret1.Id == "" { // No encontro paquete Retail
		prio1, err := c.PedirPrioritario(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling PedirPrioritario: %s", err)
		}
		if prio1.Id != "" { //Encuentra uno prioritario
			time.Sleep(time.Duration(tiempo) * time.Second)
			prio2, err := c.PedirPrioritario(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirPrioritario: %s", err)
			}
			if prio2.Id != "" {

				if prio2.Valor > prio1.Valor { //Para ver cual intentar enviar primero por valor
					dospaquetes(c, prio2, prio1, queCamion, tiempo2)
				} else {
					dospaquetes(c, prio1, prio2, queCamion, tiempo2)
				}
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR PR
			} else if prio2.Id == "" {
				unpaquete(c, prio1, queCamion, tiempo2)
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE PR
			}
		}
	}
}

func cargaNormal(c courier.CourierServiceClient, wg *sync.WaitGroup, tiempo, tiempo2 int64) { // Función para camión 3 (similar a las anteriores, sólo que en vez de retail -> prioritario y prioritario -> normal)
	defer wg.Done()
	queCamion := "Camion normal"
	message := courier.Mensaje{
		Body: "Necesito paquete.",
	}
	prio1, err := c.PedirPrioritario(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling PedirRetail: %s", err)
	}
	if prio1.Id != "" { //Encuentra el primer prioritario
		time.Sleep(time.Duration(tiempo) * time.Second)
		prio2, err := c.PedirPrioritario(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling PedirPrioritario: %s", err)
		}
		if prio2.Id != "" {
			if prio2.Valor > prio1.Valor { //Para ver cual intentar enviar primero por valor
				dospaquetes(c, prio2, prio1, queCamion, tiempo2)
			} else {
				dospaquetes(c, prio1, prio2, queCamion, tiempo2)
			}

			//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR PR

		} else if prio2.Id == "" {
			nor1, err := c.PedirNormal(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirNormal: %s", err)
			}
			if nor1.Id != "" {
				if nor1.Valor > prio1.Valor { //Para ver cual intentar enviar primero por valor
					dospaquetes(c, nor1, prio1, queCamion, tiempo2)
				} else {
					dospaquetes(c, prio1, nor1, queCamion, tiempo2)
				}
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES PR NR
			} else if nor1.Id == "" {
				unpaquete(c, prio1, queCamion, tiempo2)
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE PR
			}

		}
	} else if prio1.Id == "" { //No encuentra prioritario
		nor1, err := c.PedirNormal(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling PedirNormal: %s", err)
		}
		if nor1.Id != "" { //Encuentra el primer normal
			time.Sleep(time.Duration(tiempo) * time.Second)
			nor2, err := c.PedirNormal(context.Background(), &message)
			if err != nil {
				log.Fatalf("Error when calling PedirNormal: %s", err)
			}

			if nor2.Id != "" {
				if nor2.Valor > nor1.Valor { //Para ver cual intentar enviar primero por valor
					dospaquetes(c, nor2, nor1, queCamion, tiempo2)
				} else {
					dospaquetes(c, nor1, nor2, queCamion, tiempo2)
				}
				//TIRAR FUNCION YA QUE HAY 2 PAQUETES NR NR
			} else if nor2.Id == "" {
				unpaquete(c, nor1, queCamion, tiempo2)
				//TIRAR FUNCION YA QUE HAY 1 PAQUETE NR
			}
		}
	}
}

func dospaquetes(c courier.CourierServiceClient, x *courier.Paquete, y *courier.Paquete, camion string, tiempoEnv int64) { // Función cuando se entregan 2 paquetes
	trya := 1
	tryb := 1
	fechaa := "0"
	fechab := "0"
	rand.Seed(time.Now().Unix())
	for x.Estado == "En Camino" || y.Estado == "En Camino" { // Loop que sigue mientras uno de los paquetes siga en "En Camino"
		dt1 := time.Now()                        // Para determinar la fecha de entrega.
		if trya != 4 && x.Estado != "Recibido" { // Si se recibió o excedió los intentos, no se considera esta parte del loop.
			time.Sleep(time.Duration(tiempoEnv) * time.Second)
			r := 1 + rand.Intn(10) // Para estimar que el 80% de las veces se realiza la entrega.
			if r <= 8 {
				x.Estado = "Recibido"
				fechaa = dt1.Format("01-02-2006 15:04:05")
			} else {
				trya = trya + 1
				if trya == 4 {
					x.Estado = "No Recibido"
				}
			}
		}
		dt2 := time.Now()
		if tryb != 4 && y.Estado != "Recibido" { // Lo mismo pero para el segundo paquete.
			time.Sleep(time.Duration(tiempoEnv) * time.Second)
			r2 := 1 + rand.Intn(10)
			if r2 <= 8 {
				y.Estado = "Recibido"
				fechab = dt2.Format("01-02-2006 15:04:05")
			} else {
				tryb = tryb + 1
				if tryb == 4 {
					y.Estado = "No Recibido"
				}
			}
		}
	}
	resultadoa := courier.Entrega{ // Se forma el mensaje para enviar al server del resultado del primer paquete.
		Id:           x.Id,
		Dignipesos:   x.Valor,
		Estado:       x.Estado,
		Intentos:     int64(trya),
		Fechaentrega: fechaa,
		Camion:       camion,
	}

	_, err := c.ResultadoEntrega(context.Background(), &resultadoa) // Se llama a la función rpc.
	if err != nil {
		fmt.Println("Error, error. ", err)
	}

	resultadob := courier.Entrega{ // Se forma el mensaje para enviar al server del resultado del segundo paquete.
		Id:           y.Id,
		Dignipesos:   y.Valor,
		Estado:       y.Estado,
		Intentos:     int64(tryb),
		Fechaentrega: fechab,
		Camion:       camion,
	}

	_, err = c.ResultadoEntrega(context.Background(), &resultadob) // Se llama a la función rpc.
	if err != nil {
		fmt.Println("Error, error. ", err)
	}
}

func unpaquete(c courier.CourierServiceClient, x *courier.Paquete, camion string, tiempoEnv int64) { // Lo mismo del caso anterior, sólo que con un paquete.
	try := 1
	fecha := "0"
	rand.Seed(time.Now().Unix())
	for x.Estado == "En Camino" {
		time.Sleep(time.Duration(tiempoEnv) * time.Second)
		dt1 := time.Now()
		if try != 4 && x.Estado != "Recibido" {
			r := 1 + rand.Intn(10)
			if r <= 8 {
				x.Estado = "Recibido"
				fecha = dt1.Format("01-02-2006 15:04:05")
			} else {
				try = try + 1
				if try == 4 {
					x.Estado = "No Recibido"
				}
			}
		}
	}

	resultado := courier.Entrega{
		Id:           x.Id,
		Dignipesos:   x.Valor,
		Estado:       x.Estado,
		Intentos:     int64(try),
		Fechaentrega: fecha,
		Camion:       camion,
	}
	_, err := c.ResultadoEntrega(context.Background(), &resultado)
	if err != nil {
		fmt.Println("Error, error. ", err)
	}
}
