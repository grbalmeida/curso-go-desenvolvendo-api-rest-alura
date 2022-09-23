### Comando para instalar Gorilla Mux

```
go get -u github.com/gorilla/mux
```

### Comando para subir a imagem Docker

```
docker-compose up
```

### Comandos para acessar hostname do container postgres

```
docker-compose exec postgres sh
```

```
hostname -i
```

```
 docker inspect {CONTAINER_ID} | grep IPAddress
```