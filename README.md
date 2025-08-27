# Nome do projeto: ReLaunch

## Integrantes:
- Matheus Oliveira Mangualde - 22301194
- Henrique de Freitas Issa - 22300732
- João Pedro Bastos Neves - 22301330
- Eduardo Mapa Avelar Damasceno - 22301674
- Eike Levy Albano Neves - 22402772
- Vinícius Theodoro Giovani - 22300821

**Turma 3B2**

# como rodar
- baixar o golang
- setar no terminal 'go env GOPRIVATE=*' para conseguir acessar os repositorios privados do github
- rodar 'go mod tidy' no terminal para instalar as dependencias
- setar as variaveis de ambiente:
  - PORT: (porta que o bff vai rodar)
  - IS_INSECURE = true
  - USER_MICROSEVICE_CONN: (url de conexão para o microserviço de user)
- rodar 'go build main.go' no terminal
- rodar 'go run main.go' no terminal
