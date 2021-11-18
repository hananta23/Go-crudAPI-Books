# Go-crudAPI-Books
Golang RESTful API using PostgreSQL

first, install PostgreSQL package with a -contrib support through terminal with command

    $ sudo apt update
    $ sudo apt install postgresql postgresql-contrib
    
make sure postgre already installed properly using this command

    $ sudo -i -u postgres
    $ psql
    
    postgres=#
!![alt text](https://github.com/hananta23/Go-crudAPI-Books/blob/main/postgres.png)

clone this project, or just simply using docker

    $ docker-compose up
and run using 

    $ go run main.go 
