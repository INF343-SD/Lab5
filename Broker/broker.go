package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "Lab5/proto"

	"google.golang.org/grpc"
)

// Fulcrum Local
var fulcrumAddresses = []string{"localhost:5001", "localhost:5002", "localhost:5003"}

// Fulcrum Remote CAMBIAR IP
// var fulcrumAddresses = []string{"ip:5001", "ip:5001", "ip:5001"}

func getRandomFulcrum() string {
	rand.Seed(time.Now().UnixNano())
	return fulcrumAddresses[rand.Intn(len(fulcrumAddresses))]
}

type server struct {
	pb.UnimplementedLaboratorio_5Server
}

func (s *server) Agregar_Base_Broker(ctx context.Context, req *pb.Ing_Broker_AgregarBase) (*pb.Broker_Ing, error) {
	address := getRandomFulcrum()
	return &pb.Broker_Ing{Direccion: address}, nil
}

func (s *server) Renombrar_Base_Broker(ctx context.Context, req *pb.Ing_Broker_RenombrarBase) (*pb.Broker_Ing, error) {
	address := getRandomFulcrum()
	return &pb.Broker_Ing{Direccion: address}, nil
}

func (s *server) Actualizar_Valor_Broker(ctx context.Context, req *pb.Ing_Broker_ActualizarValor) (*pb.Broker_Ing, error) {
	address := getRandomFulcrum()
	return &pb.Broker_Ing{Direccion: address}, nil
}

func (s *server) Borrar_Base_Broker(ctx context.Context, req *pb.Ing_Broker_BorrarBase) (*pb.Broker_Ing, error) {
	address := getRandomFulcrum()
	return &pb.Broker_Ing{Direccion: address}, nil
}

func (s *server) GetEnemigos_Broker(ctx context.Context, req *pb.Kais_Broker) (*pb.Broker_Kais, error) {
	address := getRandomFulcrum()
	return &pb.Broker_Kais{Direccion: address}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	println("Server is running on port :5000")

	s := grpc.NewServer()
	pb.RegisterLaboratorio_5Server(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
