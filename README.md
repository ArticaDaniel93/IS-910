# IS-910


Producer y consumer con RabbitMQ


Ruta para el directorio: C:\Program Files\Go\src\IS-910\rabbitmq-go


Descargar Docker primeramente antes de ejecutar los comandos en Visual Studio Code

#

Poner estos comandos en Visual Studio Code:


1. #docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management

2.  #docker-compose up


docker-compose queda ejecutándose en una sóla terminal de VS code


Crear una nueva terminal para correr los siguientes comandos en ese mismo orden del 2 hasta el 5


3.  #docker pull rabbitmq
4.  #go build
5.  #go run producer.go
6.  #go run consumer.go
