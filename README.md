# Integreantes:
- Bastián Navarrete         Rol: 202073510-7
- Felipe Rojas              Rol: 201873112-9

## Consideraciones:
- Máquina 033 Broker
- Máquina 034 Fulcrum1 e Ingeniero1
- Máquina 035 Fulcrum2 e Ingeniero2
- Máquina 034 Fulcrum3 y Kais

## Estado de la tarea
La tarea se encuentra sin terminar, no contempla las acciones de consistencia de ningún tipo.
Parta ser sinceros, entregamos lo que teníamos nomás, nos dio pajita. Te queremos Iñaki, eres una bestia.

Debe ejecutar los comandos de abajo considerando estos puntos, de otro modo la tarea no funciona

# Instrucciones de ejecución
Corre el Broker
```
sudo docker compose up
```

Corre el Fulcrum
```
cd Fulcrum
sudo docker compose up
```

Corre el Ingeniero
```
cd Ingeniero
sudo docker compose up
```

Corre el Kais
```
cd Kais
docker compose up
```

Para compilar lab5.proto:
```
cd proto
```
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative lab5.proto
```
