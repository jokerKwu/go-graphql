# GraphQL Example Using GQL Gen

The purpose of this example is to provide details as to how one would go about creating a Restful API with the Echo Web Framework.

## Getting Started

## Software requirements

- [Go 1.10.3 or higher](https://golang.org/dl)

## Quick Installation

1. start the server

    ```text
    $ go run src/server.go
    ```

2. navigate to the site in the browser

    ```text
    $ open http://localhost:8080
    ```

3. test some queries

    ```text
    mutation CreateUser {
      createUser(input:{name:"joker", age:20,phone:"01012340000"}) {
          id
          name
          age
          phone
      }
    }

    query users {
      users{
        id
        name
        phone
        age
      }
    }
    ```

## GQL Gen References

- Official website: https://gqlgen.com
- Guides: https://gqlgen.com
- Docs: https://gqlgen.com
- Source: https://github.com/99designs/gqlgen

## Support

Bug reports and feature requests can be filed with the rest for the GraphQL Example Using GQL Gen project here:

- [File Bug Reports and Features](https://github.com/conradwt/gqlgen_example/issues)

## License

GraphQL Example Using GQL Gen is released under the [MIT license](https://mit-license.org).

## Copyright

copyright:: (c) Copyright 2018 Conrad Taylor. All Rights Reserved.