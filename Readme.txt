Instrucciones respecto al makefile

*Si estás en la maquina 1 de cliente (10.10.28.14), corre el comando:
    make cliente

*Si estás en la maquina 2 de camiones (10.10.28.140), corre el comando:
    make camion

*Si estás en la maquina 3 de logistica (10.10.28.141), corre el comando:
    make logistica


En caso de que las dependencias de GO no estén en orden (lo cual pasaba mucho en las VM)
Probar con los siguientes comandos

export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

Si llegan a salir errores de plugin faltante esto lo soluciona:
$ export GO111MODULE=on  # Enable module mode
$ go get github.com/golang/protobuf/protoc-gen-go \
         google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ export PATH="$PATH:$(go env GOPATH)/bin"


Consideraciones:

1. Por cada cliente, se hará envío de 1 (uno) solo csv con ordenes

2. Se están usando los csv de ejemplo (de moodle) para clientes 
retail y pyme por el momento. Si se testeará con otros, deben agregarse
en "arch/" y al correr la instancia, se debe dar el nombre en formato
<nombre>.csv

3. El envío de órdenes es bloqueante por 2 razones: evitar errores de 
escritura en el servidor; y para proveer todos los códigos de seguimiento
al cliente antes de solicitar el seguimiento.

4. No se hizo la parte de finanzas pues la maquina 4 estuvo deshabilitada
desde el jueves (15/10). No podíamos establecer conexion ssh2 y recién hoy
Lunes 20/10 se arregló. uwu

5. Perdonen el retraso de las 3 horas si es posible uwu x2



