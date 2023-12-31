# Solution for Golang Software Engineer Challenge - Pagar.me

This repository contains one solution for the Software Engineer Challenge proposed by Pagar.me. The goal was to build a service to processing payment transactions from other systems.

The challenge description can be found at [Golang Software Engineer Challenge - Pagar.me](https://github.com/pagarme/vagas/blob/master/desafios/software-engineer-golang/README.md).

## How to Use

1. Clone this repository:
```
git clone https://github.com/sesaquecruz/go-payment-processor
```

2. Enter the project directory:
```
cd go-payment-processor
```

3. Run the docker compose:
```
docker compose up -d --build
```

4. Access the API Documentation for requests and response information at:
```
http://localhost:8080/api/v1/swagger/index.html
```

5. Get the API Authorization Token from the Auth Service at:
```
http://localhost:6062/token
```

Remember to add the prefix Bearer to the token when using the API Documentation, for example:
```
Bearer token-value
```

## Predefined Test Data

### Preregistered acquirers:

- cielo	(max supported transaction value 100.00)
- rede	(max supported transaction value 500.00)
- stone	(max supported transaction value 1000.00)

If the transaction value is greater than the max supported value by the transaction acquirer, it will fail.

### Preregistered card tokens:

- 461c9432d4d7eca7ba32b783aa22ca5c89e4f396288de5128b73b461c42d4f40
- 7d2cd4f89ffe5374013d68c64ec104182366f786a377da1d3103db201149d3b5
- 4939de8e7acf6011a9b4aa4abdd6496cec40240e418a7892723ef16c4cbb44f2
- d840c6fb8401c4bbefdc4ceddc1a88f1636734bdde88c344b8969d0cd5cfdaed
- f8a8a91d9626b66a74ff7c11b5f1c2cc59a103f5b2a3119b4729a70b25304074
- cc89fefc83d423395b11998646cc7eb7c32c04ece114d1373c3a519fbb612724

The test card data can be found at [Test Cards](.docker/test-data/cards.sql)

## Tech Stack

- [Go](https://go.dev)
- [Fiber](https://gofiber.io/)
- [Postgres](https://www.postgresql.org)


## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) file for more information.
