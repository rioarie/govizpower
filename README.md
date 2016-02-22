# govizpower
High Performace Backend data service of Indonesia power consumtion. This project is for [*Koding Hackanthon*](https://www.koding.com/Hackathon) purpose. This project is using [Gin](https://github.com/gin-gonic/gin) for Web Framework and [Go-sql-driver](https://github.com/go-sql-driver/mysql) for Mysql Driver.

## Requirements
- [Gin](https://github.com/gin-gonic/gin) Web Framework
- [Gin Gzip](https://github.com/gin-gonic/contrib) Middleware
- [Go-sql-driver](https://github.com/go-sql-driver/mysql) Mysql Driver
- Go Version 1.3.1+
 
## Get Dependencies  
```
$ go get github.com/gin-gonic/gin  
$ go get github.com/gin-gonic/contrib/gzip  
$ go get github.com/go-sql-driver/mysql  
```  
## Project Structure
```
.
└── govizpower
    ├── controllers       
    ├── datamappers       // Implementing DTO
    ├── define            // Define for app
    ├── models
    ├── router        
    └── main.go
```
