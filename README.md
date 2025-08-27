# Nome do Projeto: ReLaunch

## Integrantes:
- Matheus Oliveira Mangualde - 22301194
- Henrique de Freitas Issa - 22300732
- João Pedro Bastos Neves - 22301330
- Eduardo Mapa Avelar Damasceno - 22301674
- Eike Levy Albano Neves - 22402772
- Vinícius Theodoro Giovani - 22300821

**Turma 3B2**

# Como rodar
- Baixar o golang
- Setar no terminal 'go env -w GOPRIVATE==*' para conseguir acessar os repositorios privados do github
- Rodar 'go mod tidy' no terminal para instalar as dependencias
- Setar as variaveis de ambiente:
  - PORT: (porta em que o bff vai rodar)
  - IS_INSECURE = true
  - USER_MICROSEVICE_CONN: (url de conexão para o microserviço de user)
- Rodar 'go build main.go' no terminal
- Rodar 'go run main.go' no terminal
