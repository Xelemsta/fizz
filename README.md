# fizz

- "-" is not allowed in str1 et str2 (car sert de délimiteur)
- errors returned properly (bad request, not found, etc...)
- faire test top request (bonus)
- faire test redis package
- internalRedis et redis

choix/ assumption
- incr dans handler au lieu de middleware pour être après l'authentification et la validation des params
- redis: key/value parfait pour le besoin + data type pratique pour récupérer le top hit

amélioration
- process fizzbuzz avec goroutines pour accélerer le rendu
- spawn swagger ui
- alert from alertmanager/elastalert

other
curl "localhost:3001/v1/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz"

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

* docker-compose

## About application

### Remaining tasks undone

* test

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
* go-swagger
* redis
* prometheus/grafana

### Choices/Assumptions

First of all, all my choices are mostly going towards stabilized and widely used technologies because it mostly offers a large community that can help on that matters if necessary and an exhaustive documentation.

About the API, i decided to choose go-swagger to build it because itbrings mostly all the common features we would need to build that application and is "contract-first" including skeleton generation + interactive documentation. From that we will be able to focus mostly on business code (implemeting handlers) and routes definitions.

### Extensions/Improvements

* todo