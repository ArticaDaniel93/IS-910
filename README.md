# IS-910


Producer y consumer con RabbitMQ


Ruta para el directorio: C:\Program Files\Go\src\IS-910\rabbitmq-go


Instrucciones:


Descargar Docker primeramente antes de ejecutar los comandos en Visual Studio Code


Poner estos comandos en Visual Studio Code:


1.  $docker-compose up


docker-compose queda ejecutándose en una sóla terminal de VS code


Crear una nueva terminal para correr los siguientes comandos en ese mismo orden


2.  #docker pull rabbitmq
3.  #go build
4.  $go run producer.go
5.  $go run consumer.go
