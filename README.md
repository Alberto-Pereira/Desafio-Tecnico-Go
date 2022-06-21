## 💼 Desafio Técnico GO

## 📝 Summary

- [About](#about)
- [Composition](#composition)
- [Dependencies](#dep)
- [How to Use](#howtouse)
- [License](#license)
- [Author](#author)

## 💻 About: <a name="about"></a>

- The API allows to create an account, make login, transfer amount between accounts, see how many accounts are registered in database and a balance of this accounts.

## 🏛 Composition: <a name="composition"></a>

The API is divided by two main parts:
  
  - Account:
    - Create account.
    - Get account balance.
    - Get accounts.
    - Login.
    
  - Transfer: 
    - Create transfer.
    - Get transfers.

## 🌴 Dependencies: <a name="dep"></a>

- Swagger (API Documentation): <a href="https://github.com/swaggo">github.com/swaggo</a>
- Gin Gonic (HTTP): <a href="https://github.com/gin-gonic/gin">github.com/gin-gonic/gin</a>
- JWT (Security): <a href="https://github.com/golang-jwt/jwt">github.com/golang-jwt/jwt</a>
- BCrypt (Hash): <a href="https://golang.org/x/crypto">golang.org/x/crypto</a>
- PQ (PostgreSQL): <a href="https://github.com/lib/pq">github.com/lib/pq</a>
- Testify (Tests): <a href="https://github.com/stretchr/testify">github.com/stretchr/testify</a>

## 🐳 How to Use: <a name="howtouse"></a>
⚠ Have <a href="https://www.docker.com/products/docker-desktop/">Docker</a> installed.
- Clone the repository:
```bash
git clone https://github.com/Alberto-Pereira/desafio-tecnico-go
```
- Access the cloned directory:
```bash
cd desafio-tecnico-go
```
- Build the API:
```bash
docker-compose build
```
- Start the API:
```bash
docker-compose up
```

👉 After started you can use the API with <a href="http://localhost:8080/swagger/index.html">Swagger</a>.

## 📃 License: <a name="license"></a>

- <a href="http://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

## 👁‍🗨 Author: <a name="author"></a>

- <a href="https://github.com/Alberto-Pereira">Alberto Pereira</a>

 
