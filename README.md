## 💼 Desafio Técnico GO

## 📝 Summary

- [About](#about)
- [Hierarchy](#pattern)
- [Dependencies](#dep)
- [How to Use](#howtodownload)
- [License](#license)
- [Author](#author)

## 💻 About: <a name="about"></a>

- The API allows to create an account, make login, transfer amount between accounts, see how many accounts are registered in database and a balance of this accounts.

## 🏛 Hierarchy: <a name="pattern"></a>

- A aplicação é composta das seguintes partes:
    - model: contém o arquivo <b>item.go</b>, que representa a entidade principal da aplicação.
    - service: contém o arquivo <b>item-service.go</b>, que contém as regras de serviço da entidade item.
    - util: contém os arquivos <b>gerador-de-emails.go</b> e <b>gerador-de-itens.go</b>, que auxiliam na criação de testes/casos para a aplicação.

## 🌴 Dependencies: <a name="dep"></a>

- Testify: <a href="https://github.com/stretchr/testify">github.com/stretchr/testify</a>
- Randstr: <a href="https://github.com/thanhpk/randstr">github.com/thanhpk/randstr</a>

## 🐳 How to Use: <a name="howtodownload"></a>
⚠ Ter o <a href="https://www.docker.com/products/docker-desktop/">Docker</a> instalado.
- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/Desafio-Tecnico.git
```
- Acessar o diretório clonado:
```bash
cd Desafio-Tecnico
```
- Fazer a build da aplicação:
```bash
docker build -t docker-desafio-tecnico .
```
- Iniciar a aplicação:
```bash
docker run docker-desafio-tecnico
```

## 📃 License: <a name="license"></a>

- <a href="http://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

## 👁‍🗨 Author: <a name="author"></a>

- <a href="https://github.com/Alberto-Pereira">Alberto Pereira</a>

 
