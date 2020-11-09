# pokemon-api

Pokemon-api is an API which provides Shakespearean description for pokemons in the `ruby` version of the game.

```
├───api
│   └───swagger
│       ├───implementation
│       ├───models
│       ├───modelsdefinitions
│       └───restapi
│           └───operations
│               └───pokemondescription
├───bin
├───cache
├───cmd
├───config
├───core
│   ├───pokemon
│   │   └───pokeapi
│   ├───service
│   └───shakespeare
└───tooling
```

# `api` folder
Api folder contains swagger generated server files and the implementations of the endpoints in the implementation folder as implementation package.

`swagger.yml` is also in the `api` folder. 

# `bin` folder
bin folder contains useful scripts and binary executables, such as `mockery` which is the tool used to generate mock files.

# `cache` folder
Cache folder contains the cache package which implements the cache functionality used within the app. At the moment `Cache` interface has only one implementation and that is `RedisCache`. Results of the calls are stored in redis by making use of this package.

# `cmd` folder
Cmd folder contains the `main.go` file which is the entry point for the build of the app. This file is generated by `swagger` command and should not be edited.

# `config` folder
Config folder contains configuration functionality of the project. At the moment 
`.env.local` file is read for environment variables. However anything provided to the container during execution will overwrite the contents of this file, which can be observed by the redis address passed in the `docker-compose.yml`. 

# `core` folder
Core folder contains the main functionality of the app. Integrations with the `poke-api` and the `shakespeare-translation` api is contained within this folder.
`core/pokemon/pokeapi` contains the integration with the `poke-api`.

`core/shakespeare` contains the integration with `shakespeare-translation`.

`core/service` contains the service implementation of the app. This is practically the main functionality of the app as of now and is responsible for returning the desired description of the pokemon in Shakespearean English.


# Working with the project
Dev environment is designed to be crossplatform and should be able to work on a linux, windows or a mac (fingers crossed).

## Prepare the dev environment

Execute the following, which will create a container named `tooling` that is workable to run any operations on the code base regardless of the platform development is being made.

```bash
bin/prepare.sh
```

## Runing the tests

```bash
bin/make.sh unit-tests
```
This command will run the makefile target `unit-tests` within the `tooling` container.