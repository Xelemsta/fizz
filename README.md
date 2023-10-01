# Fizz Buzz
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...".

## Description

Your goal is to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:
- Ready for production
- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request.

## Prerequisites

* make
* docker-compose
* docker
* git
* go (to launch tests)

### Start application

  ```sh
  make run
  ```

### Test application

  ```sh
  make test
  ```

### Technical environment

* docker
* go swagger
* redis
* prometheus
* alertmanager
* grafana

### Technical choices/assumptions

* Go Swagger to generate API (contract first + cover all basic needs)
* Redis to store top request stats (using sorted set data type that perfectly fits what we need to achieve)
* Prometheus/Alertmanager to monitor app and perform alerts if necessary
* Grafana to visualize app metrics

### Extensions/Improvements

* Spawn swagger UI
* Use goroutines to perform fizzbuzz
* Use persistent storage for prometheus metrics (mimir, opensearch...)
* use a vault server to store secrets (passwords, configuration)
* Configure alertmanager