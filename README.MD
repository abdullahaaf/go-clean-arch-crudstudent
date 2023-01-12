This is an example implementation of Hexagonal Architecture that proposed by [Alistair Cockburn](https://alistair.cockburn.us/hexagonal-architecture/) in Golang.

Hexagonal Architecture in what i understand is to make our core business logic not depend on any third parties, so when in our project have something to change on external systems, our core not impacted.

Hexagonal Architecture has four main concepts : core, adapter, ports, and actor.

Core
Core contains all business logic and entities, it is indpedent and dont know anything outside of itself, in my project defined in core folder.

Actor
Actor is everything that interact with core. Such as database, http request, and etc. Actor and Core speaks with different language, so they need translator and here ports and adapter come.

Ports
Ports in my project are used to define set of interface that define how communication between core and actors should be done, or in other words its a set of communication rule between core and actor.

Adapter
Adapters in what iam understand is translator between core and actor, in my project adapters are defined in repositories and handlers folder. repositories act as translator between core and mysql database, and handlers act as translator between core and http request.

Other that, i have usecase folder, this folder is aimed to define specific action that my projects need to perform.