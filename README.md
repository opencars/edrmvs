# edrmvs

> Tool for extracting infromation from government resources

## Development

Build the binary

```sh
make
```

Start postgres

```sh
docker-compose up -Vd postgres
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1/edrmvs\?sslmode=disable up
```

Run the web server

```sh
./bin/server
```

## Test

Start postgres

```sh
docker-compose up -Vd postgres
```

Run sql migrations

```sh
migrate -source file://migrations -database postgres://postgres:password@127.0.0.1/edrmvs\?sslmode=disable up
```

Run tests

```sh
go test -v ./...
```

## Usage

For example, you get information about this amazing Tesla Model X

```sh
http http://localhost:8080/api/v1/registrations/СХН484154
```

```json
{
    "brand": "TESLA",
    "code": "CXH484154",
    "color": "ЧОРНИЙ",
    "date": "2019-06-05",
    "first_reg_date": "2016-10-13",
    "fuel": "ЕЛЕКТРО",
    "kind": "ЛЕГКОВИЙ УНІВЕРСАЛ-B",
    "model": "MODEL X",
    "num_seating": 7,
    "number": "AA9359PC",
    "own_weight": 2485,
    "rank_category": "B",
    "total_weight": 3021,
    "vin": "5YJXCCE40GF010543",
    "year": 2016
}
```

## License

Project released under the terms of the MIT [license](./LICENSE).