package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"

	pb "Lab5/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedLaboratorio_5Server
}

func (s *server) Agregar_Base_Fulcrum(ctx context.Context, req *pb.Ing_Fulcrum_AgregarBase) (*pb.Fulcrum_Ing, error) {
	EscribirLog("AgregarBase", req.Base, req.Sector)
	return &pb.Fulcrum_Ing{Reloj: "Reloj vectorial"}, nil
}

func (s *server) Renombrar_Base_Fulcrum(ctx context.Context, req *pb.Ing_Fulcrum_RenombrarBase) (*pb.Fulcrum_Ing, error) {
	EscribirLog("RenombrarBase", req.Base, req.Sector, req.NuevoNombre)
	return &pb.Fulcrum_Ing{Reloj: "Reloj vectorial"}, nil
}

func (s *server) Actualizar_Valor_Fulcrum(ctx context.Context, req *pb.Ing_Fulcrum_ActualizarValor) (*pb.Fulcrum_Ing, error) {
	EscribirLog("ActualizarValor", req.Base, req.Sector, fmt.Sprint(req.NuevoValor))
	return &pb.Fulcrum_Ing{Reloj: "Reloj vectorial"}, nil
}

func (s *server) Borrar_Base_Fulcrum(ctx context.Context, req *pb.Ing_Fulcrum_BorrarBase) (*pb.Fulcrum_Ing, error) {
	EscribirLog("BorrarBase", req.Base, req.Sector)
	return &pb.Fulcrum_Ing{Reloj: "Reloj vectorial"}, nil
}

func (s *server) GetEnemigos_Fulcrum(ctx context.Context, req *pb.Kais_Fulcrum) (*pb.Fulcrum_Kais, error) {
	EscribirLog("GetEnemigos", req.Base, req.Sector)
	return &pb.Fulcrum_Kais{CantEnemigos: 10, Reloj: "Reloj vectorial"}, nil
}

// Crea o actualiza un archivo de texto para cada sector y base con la cantidad de soldados
func AlmacenarSoldados(sector, base string, cantidad int32) {
	file := fmt.Sprintf("%s_%s.txt", sector, base)
	err := ioutil.WriteFile(file, []byte(fmt.Sprint(cantidad)), 0644)
	if err != nil {
		fmt.Println("No se pudo almacenar la cantidad de soldados")
	}
	fmt.Printf("Almacenado %d soldados en %s_%s.txt\n", cantidad, sector, base)
}

func EscribirLog(accion, base, sector string, valor ...string) {
	log := fmt.Sprintf("Accion: %s, Base: %s, Sector: %s, Valor: %s\n", accion, base, sector, valor)
	err := ioutil.WriteFile("log.txt", []byte(log), 0644)
	if err != nil {
		fmt.Println("No se pudo escribir en el log")
	}
}

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		fmt.Println("No se pudo escuchar en el puerto 5001")
		return
	}

	s := grpc.NewServer()
	pb.RegisterLaboratorio_5Server(s, &server{})
	fmt.Println("Fulcrum escuchando en puerto 5001")
	s.Serve(lis)
}
