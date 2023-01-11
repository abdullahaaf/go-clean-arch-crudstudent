This is an example implementation of Clean Architecture by Uncle Bob in Golang, and this repository referenced from this repo (https://github.com/bxcodec/go-clean-arch/tree/v1) with some customization by me.

This Project has 4 domain layer :
- Models Layer (student/model) => Used to store any Object's struct and its method, this layer will used in all layers
- Repository Layer (student/repository) => Used to handle any database operation. If our project need to call microservices, it will handled in this layer too.
- Usecase Layer (student/usecase) => Business process of our project will handled in this layer. This layer will decide what repository will used. This layer also have responsibility to provide data into delivery layer, but also can receive input data from delivery layer. The data from delivery layer must be in clean formatted data. This layer depends to Repository layer
- Delivery Layer (student/delivery) => This layer act like presenter layer, decide how data will be presented. The data could be Rest API, HTML, etc. On this example data presented as Rest API. This layer depend on Usecase Layer.

All layer connected through main.go file (except model layer).