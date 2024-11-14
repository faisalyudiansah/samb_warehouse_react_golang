# samb_warehouse_react_golang
Application to record Incoming Goods at Warehouse A and Outgoing Goods from Warehouse A with Golang (PostgreSQL) and React JS (shadcn ui and aceternity ui)

## Backend
- `cd be/`
- set up your .env file (just copy from .env.examples)
- `go mod tidy`
- for seeding data, it will automatically migrate and seed data every time the server starts. But if you use LINUX or you can use Makefile, you can type on CLI `make db-generate`
- `go run main.go` or if you use LINUX or you can use Makefile, you can type on CLI with `make run` or `make start-nodemon` http://localhost:8081

## Frontend
- `cd fe/`
- `npm install`
- set up .env file (just copy from .env.examples)
- `npm run dev` http://localhost:3000

## Postman docs = https://documenter.getpostman.com/view/29632965/2sAY4rDPgt
## github = https://github.com/faisalyudiansah/samb_warehouse_react_golang
