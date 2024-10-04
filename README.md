 ## Getting Started

These instructions will help you get the project up and running on your local machine.

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- [Dbmate](https://github.com/amacneil/dbmate)
- [Golang](https://go.dev/dl/)

### Usage

Explain how to use the project and provide examples.
1. **Clone Repository**

   Clone repositori:

   ```shell
   git clone https://github.com/aripzz/UserPost
   cd UserPost
    ```

2. **Create .env File**

   Copy the example environment file:

   ```shell
    cp .env.example .env
    ```

3. **Start the Docker Compose**

   ```shell
    docker compose up --build
    ```
   Run Containers in Detached Mode (containers will run in the background):
   ```shell
    docker compose up --build -d
    ```

4. **Start the Migration**
   ```shell
    dbmate up
    ```

5. **Services**
- API (Golang)          = http://localhost:3000
- Documention Swagger   = http://localhost:3000/swagger 
- Postgree Database     = localhost:5432
- Redis Cache           = localhost:6379

# alasan menggunakan clean architecture design pattern
clean architecture mudah dibaca. Struktur yang jelas membantu untuk memahami alur aplikasi dengan cepat, sehingga memudahkan proses maintenance dan development.
