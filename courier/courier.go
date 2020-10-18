package courier

import (
	"log"
	"math/rand"

	"golang.org/x/net/context"
)

//Server Se declara la estructura del servidor
type Server struct {
}

//Hello Implementamos la función básica de Hello para verificar la conexión
func (s *Server) Hello(ctx context.Context, message *Mensaje) (*Mensaje, error) {
	log.Printf("Received message body from client: %s", message.Body)
	return &Mensaje{Body: "Hello From the Server!"}, nil
}

//CodigoPyme Esta es la prueba inicial de mi función para devolver código de seguimiento
func (s *Server) CodigoPyme(ctx context.Context, orden *OrdenPyme) (*Codigo, error) {
	log.Printf("Orden recibida del producto: %s", orden.Id)
	return &Codigo{Cod: int64(rand.Intn(1000000000000000))}, nil
}

func (s *Server) mustEmbedUnimplementedCourierServiceServer(){}