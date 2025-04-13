# 🚀 go-auto-scale-app

Uma API desenvolvida em **Go** usando o framework [Gin](https://github.com/gin-gonic/gin), criada para simular **carga de CPU** e testar comportamentos de **Auto Scaling** em ambientes como o **AWS Auto Scaling Group**.

## 🚀 Propósito

Essa API foi criada para **testes práticos de escalabilidade**, especialmente em ambientes que reagem a picos de uso de CPU, como:

- AWS Auto Scaling Group
- Kubernetes Horizontal Pod Autoscaler (HPA)
- Testes de observabilidade e monitoramento

Ideal para quem quer entender na prática como a infraestrutura responde sob demanda!

## 🔧 Endpoints

### `GET /fibonacci/:qtd`

Calcula a sequência de Fibonacci até a quantidade `:qtd` informada.

⚠️ O cálculo é intencionalmente **ineficiente** (recursivo) para **estressar o processador** e acionar regras de escalabilidade com base em uso de CPU.

**Exemplo:**
```json
{"result":102334155}
```

### `GET /info`
Retorna o hostname e o nome do sistema operacional da instância/container.

**Exemplo:**
```json
{
  "message": {
    "hostname": "ip-10-11-0-94.ec2.internal",
    "os": "linux"
  }
}
```

## 📦 Como rodar

### Com Go instalado

```bash
go get ./cmd/api/
go run cmd/api/main.go
```

### Gerar executavel para Linux

```bash
env GOOS=linux GOARCH=amd64 go build -o auto-scale-app cmd/api/main.go
```

### Com Docker

```bash
docker build -t go-auto-scale-app .
docker run -p 8080:8080 go-auto-scale-app
```

## 📦 Dica para testes

Você pode usar ferramentas como Apache Benchmark (ab), hey ou wrk para gerar requisições e simular carga.

### Exemplo com hey

```bash
hey -n 1000 -c 10 http://localhost:8080/fibonacci/35
```