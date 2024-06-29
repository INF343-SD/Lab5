package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "Lab5/proto"

	"google.golang.org/grpc"
)

var broker_dir = "localhost:5000"

func getEnemigos(sv pb.Laboratorio_5Client, sector, base string) {
	resp, err := sv.GetEnemigos_Broker(context.Background(), &pb.Kais_Broker{Sector: sector, Base: base})
	if err != nil {
		log.Fatalf("No se pudo obtener la respuesta del Broker: %v", err)
	}
	fmt.Println("Servidor fulcrum: ", resp.Direccion)

	connFulcrum, err := grpc.Dial(resp.Direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Fulcrum: %v", err)
	}
	defer connFulcrum.Close()

	fulcrumClient := pb.NewLaboratorio_5Client(connFulcrum)
	time.Sleep(10 * time.Second)

	respFulcrum, err := fulcrumClient.GetEnemigos_Fulcrum(context.Background(), &pb.Kais_Fulcrum{Sector: sector, Base: base})
	if err != nil {
		log.Fatalf("No se pudo obtener la respuesta del Fulcrum: %v", err)
	}
	fmt.Println("Enemigos: ", respFulcrum.CantEnemigos)
	fmt.Println("Reloj vectorial: ", respFulcrum.Reloj)

}

func main() {
	connBroker, err := grpc.NewClient(broker_dir, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Broker: %v", err)
	}
	defer connBroker.Close()

	sv := pb.NewLaboratorio_5Client(connBroker)

	fmt.Println("Ingresa el sector: ")
	var sector, base string
	fmt.Scanln(&sector)
	fmt.Println("Ingresa la base: ")
	fmt.Scanln(&base)

	getEnemigos(sv, sector, base)
}
