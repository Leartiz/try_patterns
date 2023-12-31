# v1

## Diagrams 📊

### Class

```mermaid
classDiagram

    class Type {
        <<enumeration>>
        MEMORY
        SQL_POSTGRES
        SQL_LITE
    }

    class PersonStorage {
        <<interface>>
        GetPersons() []Person, error
        GetPersonById(int) Person, error
        InsertPerson(string, string, int) int, error    
        UpdatePerson(int, string, string, int) error
        DeletePersonById(int) error
        ...()
    }

    class CompanyStorage {
        <<interface>>
        GetCompanies() []Company, error
        InsertCompany(string) int, error
        ExistsCompany(int) bool, error
        ...()
    }

    class Storage {
        <<interface>>
        GetPersonById(int) Person, error
        InsertPerson(string, string, int) int, error
        UpdatePerson(int, string, string, int) error
        DeletePersonById(int) error
        ...()

        GetCompanies() []Company, error
        InsertCompany(string) int, error
        ExistsCompany(int) bool, error
        ...()
    }

    Storage --|> PersonStorage
    Storage --|> CompanyStorage

    %% ----------------------------------------------

    class MemoryStorage {
        GetPersonById(int) Person, error
        InsertPerson(string, string, int) int, error
        ...()

        - rwMx sync.RWMutex
        - persons   map[int]Person
        - companies map[int]Company

        - nextPersonId  int
        - nextCompanyId int
    }
    MemoryStorage --|> Storage

    class Person {
        + GetId() int
        + GetFirstName() string
        + GetLastName() string
        + GetCompanyId() int

        + SetFirstName(string) bool
        + SetLastName(string) bool

        - id        int
        - firstName string
        - lastName  string
        - companyId int
    }

    class Company {
        + GetId() int
        + GetName() string

        + SetName(string) bool

        - id   int
        - name string
    }

    MemoryStorage o-- Person
    MemoryStorage o-- Company

    %% ----------------------------------------------

    %% redirect getters/setters calls to person!
    class PersonRowGateway {
        + Make(Type, string, string, int) PersonRowGateway, error$
        + Find(Type, id int) PersonRowGateway, error$

        + GetId() int
        + GetFirstName() string
        + GetLastName() string
        + GetCompanyId() int

        + SetFirstName(string) bool
        + SetLastName(string) bool

        + Insert() error
        + Update() error
        + Delete() error

        - person Person
        - storageInstance Storage
    }

    PersonRowGateway --> Type
    
    %% or inheritance?
    PersonRowGateway *-- Person
    PersonRowGateway o-- Storage
```

## Details

- main [here](main.go)
- Storage [here](./storage/storage.go)
- MemoryStorage [here](./storage/impl/memory/storage.go)
- Person [here](./domain/person.go)
- PersonRowGateway [here](./storage/gateway/row/person.go)
