# Packt Publishing - Hands on Microservices with Go

## Section 3 - Video 4 - Introduction to gRPC and Protocol Buffers - Part 2

## Creating The Postgress Container with the data

### Go to home dir and create data directory

``` bash

cd ~
mkdir data

```

### Create Docker Data Volume

``` bash

sudo docker volume create s3v6pgdata

```

### Restore the data from the backup to the Data Volume

**VERY IMPORTANT:** YOU NEED TO CHANGE $DATA_PATH TO THE FULL PATH OF YOUR DATA DIRECTORY

``` bash
export DATA_PATH="/Users/hpe_mrusso/Documents/Projects/russomi-labs/Hands-on-Microservices-with-Go/section-3/video-4/data"

sudo docker run -it --name restore-data-container -v s3v6pgdata:/volume -v ${DATA_PATH}:/backup ubuntu \
    sh -c 'rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/s3v6.tar.bz2'

```

### Start a postgress container with this volume

``` bash

sudo docker run --name postgres -v s3v6pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=packt -e POSTGRES_USER=packt -e POSTGRES_DB=wta -d postgres:10

```

### Clean Up

``` bash
rm -Rf ~/home/data
sudo docker rm restore-data-container

```

## Verify that the Postgres Container is running and has the data

### Run bash on the container in interactive mode

``` bash

sudo docker exec -it postgres /bin/bash

```

### Get into the Postgres CLI

``` bash

psql -U postgres

```

### Navigate through the data

``` bash
#Connect to wta database
\c wta
#Show tables
\dt
#Show schema of players
\d+ players
#Show schema of rankings
\d+ rankings
#Count of players
select count(*) from players;
#Count of Rankings
select count(*) from rankings;

# quit
\q
```

## Credits

This video uses the following Dataset: (<https://www.kaggle.com/joaoevangelista/wta-matches-and-rankings>) that was created by [Jeff Sackmann](https://github.com/JeffSackmann) and was downloaded from [Kaggle](https://www.kaggle.com/joaoevangelista/wta-matches-and-rankings). It has a [Attribution-NonCommercial-ShareAlike 4.0 International (CC BY-NC-SA 4.0)](https://creativecommons.org/licenses/by-nc-sa/4.0/) Creative Commons License.

## Learn More

[Postgresql](https://www.postgresql.org/)

[Protocol Buffers Developer's Guide](https://developers.google.com/protocol-buffers/docs/overview)

[Protocol buffers Version 3 - Documentation](https://developers.google.com/protocol-buffers/docs/proto3)

### Learn even more

[Protocol buffers Encoding](https://developers.google.com/protocol-buffers/docs/encoding)

[The GRPC Blog](https://grpc.io/blog/)

[GRPC Beyond the Basics](https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/)

[Handling Errors in GRPC](https://hackernoon.com/handling-errors-in-golang-grpc-and-go-kit-services-d0fa0a112449)

[Quick start](https://grpc.io/docs/languages/go/quickstart/)
