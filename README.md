# Go Wistlist REST API

Go project can be used to create and manage any sort of **wishilists**, such as movie, book or shopping wishlists.

## Prerequirements

* Installed [go](https://go.dev/dl/) with version 1.20.x or highter.
* Installed [Docker](https://www.docker.com/get-started/).
* Installed [GNU make](https://leangaurav.medium.com/how-to-setup-install-gnu-make-on-windows-324480f1da69)

## Installation

Install **go-wishlist-api** with `git clone`

```bash
  git clone https://github.com/stas-bukovskiy/go-wishlist-api.git
  cd go-wishlist-api/
```

## Running locally

Before running application, thera are naccecery to set env and configs variables:

1. In the file `.env` change `MINIO_ENDPOINT` to `localhost:9000`
1. In the file `configs/config.yaml` change `host` to `localhost` and `port` to `5436`
   Run **go-wishlist-api** with `make`:

```
    make local-run
    make go-local-run
```

## Running as docker container

Before running application as docker container check env and cinfigs variableas are set properly:

1. In the file `.env`, `MINIO_ENDPOINT` var must be set to `minio:9000`.
2. In the file `configs/config.yaml` change `host` to `wishlist-db` and `port` to `5432`.
   Run **go-wishlist-api** with `make`:

```
    make docker-run
```

## Usage

After api starting you can go to swagger
documantation ([http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)), where every
endpoing is describer and evern you can execute each of them. Also you can use Postman or any CLI instruments for
exectiong http requests.