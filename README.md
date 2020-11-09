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
│               ├───pokemondescription
│               └───pokemons
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

# ENDPOINTS

```
GET /v1/pokemon?offset={integer}

{
    "count" : "1050",
    [
        {
            "name":"pikachu"
        },
        ...
    ]
}
```


```
GET /v1/pokemon/{pokemon-name}

{
    "description": "Description of the pokemon",
    "name": "Name of the pokemon"
}
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

Execute the following, which will create a container named `tooling` that is workable to run any operations on the code base regardless of the platform on which the development is being made.

```bash
bin/prepare.sh
```
If permission is denied, run

```bash
chmod u+x bin/prepare.sh
```

## Runing the tests

```bash
bin/make.sh unit-tests
```
This command will run the makefile target `unit-tests` within the `tooling` container.

If developing on windows, use the following

```bash
bin/make.sh --win unit-tests
```

## Generating swagger files
To generate new swagger files if any change to `swagger.yml` is made. Run the following:

```bash
bin/make.sh swagger
```
If developing on windows, use the following

```bash
bin/make.sh --win swagger
```

## Generating mocks with `mockify`
```bash
bin/make.sh generate-mocks
```

If on a windows machine:
```bash
bin/make.sh --win generate-mocks
```


## Running the project
To run the project simply use the `docker-compose`.

```bash
docker-compose up
```

If for some reason, docker compose is not accessable, redis and the app can be run seperately like the following:

```bash
docker build -t pokemon-api .
docker run --name redis -d -p 6379:6379 redis
docker run --link redis -p 8080:8080 --env ENV=local pokemon-api:latest
```

Redis has to be run first as the app is dependent on it for caching purposes


# Future additions
- Tests are missing for the pokemons API, unforunately no time left for it
- Redis integration tests
- e2e tests would be nice
- Load tests should be added
- TLS configuration would be needed to be done
- Make target to generate swagger documentation
- Sidecar app to serve swagger docs
- Logging in the api calls can be made into a generic design buy wrapping the functions
- At the moment the first requests which does not hit the cache takes about 500-600 milliseconds to complete, which is not great for a production endpoint
even though it would happen once within the time to live duration of the cache. To make this better, a second app could be created either as a separate thread
(or go routine in this app's case) within the app or as a sidecar to periodically pull all the existing pokemons for the ruby version and generate the descriptions and save into 
the cache whenever possible. This would prevent the users from experiencing the extra wait during the first requests for a pokemon after the time to live of cache values end.
Even though current version is not great, after the first request the average time for the requests completion is about 2 ms within the app, and 5 ms end to end for the user.
- We need a rate limiter in this app or the load balancer/router app in the platform, however that is a bigger piece of work to do properly, so it is skipped for this project
