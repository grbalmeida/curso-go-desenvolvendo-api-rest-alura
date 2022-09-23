# Instalações

### Comando para instalar Gorilla Mux

```
go get -u github.com/gorilla/mux
```

### Comando para instalar ORM Gorm

```
go get -u gorm.io/gorm
```

### Comando para instalar driver do Postgres para o Gorm

```
go get gorm.io/driver/postgres
```

# Docker

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

# Curl

### Comando para criar uma nova personalidade

```
curl -d "{\"nome\":\"Nome 1\", \"historia\":\"História 1\"}" -H "Content-Type: application/json" -X POST http://localhost:8000/api/personalidades
```

### Comando para atualizar uma personalidade

```
curl -d "{\"nome\": \"Pessoa 1 Atualizada\", \"historia\": \"Historia Atualizada\"}" -H "Content-Type: application/json" -X PUT http://localhost:8000/api/personalidades/16
```

### Comando para excluir uma personalidade

```
curl -X DELETE http://localhost:8000/api/personalidades/3
```