# go-rest-template
This is a new template repository for create web service

## Overview
- Using port & adapter / hexagonal architecture pattern\
  	Some illustration : 
	![images](https://blog.janetacarr.com/content/images/size/w2000/2023/04/Architecture-and-Design-diagrams---Page-4-2.png "a title")
	More explanation about this architecture you can learn from this : 
	- [Port and Adapter Architecture](https://codesoapbox.dev/ports-adapters-aka-hexagonal-architecture-explained/) 
	- [Another Port and Adapter Architecture](https://medium.com/wearewaes/ports-and-adapters-as-they-should-be-6aa5da8893b)

- Frameworks / tools : 
	- [Consul](https://www.consul.io/) for config management
	- [Echo](https://echo.labstack.com/) for http router
	- [Grpc](google.golang.org/grpc) for grpc connection
	- [Sqlx](https://github.com/jmoiron/sqlx) for database stuff (configuration, query, etc)
