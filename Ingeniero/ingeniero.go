package main

import (
	"context"
	"fmt"
	"log"

	pb "Lab5/proto"

	"google.golang.org/grpc"
)

func AgregarBase(sv pb.Laboratorio_5Client, sector, base string, params ...int) {
	resp, err := sv.Agregar_Base_Broker(context.Background(), &pb.Ing_Broker_AgregarBase{Sector: sector, Base: base})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Broker: %v", err)
	}
	fmt.Println("Servidor fulcrum: ", resp.Direccion)

	connFulcrum, err := grpc.Dial(resp.Direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Fulcrum: %v", err)
	}
	defer connFulcrum.Close()

	fulcrumClient := pb.NewLaboratorio_5Client(connFulcrum)
	fulcrumResp, err := fulcrumClient.Agregar_Base_Fulcrum(context.Background(), &pb.Ing_Fulcrum_AgregarBase{Sector: sector, Base: base, Valor: int32(params[0])})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Fulcrum: %v", err)
	}
	log.Printf("Reloj vectorial: %s", fulcrumResp.Reloj)

}

func RenombrarBase(sv pb.Laboratorio_5Client, sector, base, nuevo_nombre string) {
	resp, err := sv.Renombrar_Base_Broker(context.Background(), &pb.Ing_Broker_RenombrarBase{Sector: sector, Base: base, NuevoNombre: nuevo_nombre})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Broker: %v", err)
	}
	fmt.Println("Servidor fulcrum: ", resp.Direccion)

	connFulcrum, err := grpc.Dial(resp.Direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Fulcrum: %v", err)
	}
	defer connFulcrum.Close()

	fulcrumClient := pb.NewLaboratorio_5Client(connFulcrum)
	fulcrumResp, err := fulcrumClient.Renombrar_Base_Fulcrum(context.Background(), &pb.Ing_Fulcrum_RenombrarBase{Sector: sector, Base: base, NuevoNombre: nuevo_nombre})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Fulcrum: %v", err)
	}
	log.Printf("Reloj vectorial: %s", fulcrumResp.Reloj)

}

func ActualizarValor(sv pb.Laboratorio_5Client, sector, base string, nuevo_valor int32) {
	resp, err := sv.Actualizar_Valor_Broker(context.Background(), &pb.Ing_Broker_ActualizarValor{Sector: sector, Base: base, NuevoValor: nuevo_valor})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Broker: %v", err)
	}
	fmt.Println("Servidor fulcrum: ", resp.Direccion)

	connFulcrum, err := grpc.Dial(resp.Direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Fulcrum: %v", err)
	}
	defer connFulcrum.Close()

	fulcrumClient := pb.NewLaboratorio_5Client(connFulcrum)
	fulcrumResp, err := fulcrumClient.Actualizar_Valor_Fulcrum(context.Background(), &pb.Ing_Fulcrum_ActualizarValor{Sector: sector, Base: base, NuevoValor: nuevo_valor})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Fulcrum: %v", err)
	}
	log.Printf("Reloj vectorial: %s", fulcrumResp.Reloj)
}

func BorrarBase(sv pb.Laboratorio_5Client, sector, base string) {
	resp, err := sv.Borrar_Base_Broker(context.Background(), &pb.Ing_Broker_BorrarBase{Sector: sector, Base: base})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Broker: %v", err)
	}
	fmt.Println("Servidor fulcrum: ", resp.Direccion)

	connFulcrum, err := grpc.Dial(resp.Direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Fulcrum: %v", err)
	}
	defer connFulcrum.Close()

	fulcrumClient := pb.NewLaboratorio_5Client(connFulcrum)
	fulcrumResp, err := fulcrumClient.Borrar_Base_Fulcrum(context.Background(), &pb.Ing_Fulcrum_BorrarBase{Sector: sector, Base: base})
	if err != nil {
		log.Fatalf("No se pudo obtener respuesta del Fulcrum: %v", err)
	}
	log.Printf("Reloj vectorial: %s", fulcrumResp.Reloj)
}

func main() {
	connBroker, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar con el Broker: %v", err)
	}
	defer connBroker.Close()

	sv := pb.NewLaboratorio_5Client(connBroker)

	var sector, base, nuevo_nombre string
	var nuevo_valor int32

	fmt.Println("1. Agregar base")
	fmt.Println("2. Renombrar base")
	fmt.Println("3. Actualizar valor")
	fmt.Println("4. Borrar base")
	fmt.Println("Ingresa el número de la operación que deseas realizar: ")
	var operacion int
	fmt.Scanln(&operacion)
	switch operacion {
	case 1:
		fmt.Println("Ingresa el sector: ")
		fmt.Scanln(&sector)
		fmt.Println("Ingresa la base: ")
		fmt.Scanln(&base)
		AgregarBase(sv, sector, base)
	case 2:
		fmt.Println("Ingresa el sector: ")
		fmt.Scanln(&sector)
		fmt.Println("Ingresa la base: ")
		fmt.Scanln(&base)
		fmt.Println("Ingresa el nuevo nombre: ")
		fmt.Scanln(&nuevo_nombre)
		RenombrarBase(sv, sector, base, nuevo_nombre)
	case 3:
		fmt.Println("Ingresa el sector: ")
		fmt.Scanln(&sector)
		fmt.Println("Ingresa la base: ")
		fmt.Scanln(&base)
		fmt.Println("Ingresa el nuevo valor: ")
		fmt.Scanln(&nuevo_valor)
		ActualizarValor(sv, sector, base, nuevo_valor)
	case 4:
		fmt.Println("Ingresa el sector: ")
		fmt.Scanln(&sector)
		fmt.Println("Ingresa la base: ")
		fmt.Scanln(&base)
		BorrarBase(sv, sector, base)
	default:
		fmt.Println("Operación no válida")

	}
}
