[![Build Status](https://travis-ci.org/DiegoSantosWS/restfulgo.svg?branch=master)](https://travis-ci.org/DiegoSantosWS/restfulgo)

# RESTFUL-GO 
CREATING A SIMPLE API USING GOLANG

## CLONE

```bash
$ git clone git@github.com:DiegoSantosWS/restfulgo.git
```
## CREATING DATABASE

CREATE A DATABASE AND AFTER EXECUTE THE CODE BELOW.
```bash
$ mysql -u root -p ****
mysql> CREATE DATABASE nameyourdb;
mysql> SHOW DATABASES;
```

```sql
CREATE TABLE `products` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`name` varchar(50) DEFAULT NULL,
`description` longtext,
`stock` decimal(10,3) DEFAULT NULL,
`width` decimal(10,3) DEFAULT NULL,
`height` decimal(10,3) DEFAULT NULL,
`amount` decimal(10,3) DEFAULT NULL,
`weight` decimal(10,3) DEFAULT NULL,
`price` decimal(10,3) DEFAULT NULL,
`discount` decimal(10,3) DEFAULT NULL,
`promotion` decimal(10,3) DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

# SECRET KEY

Api requires a serect key to toke validation, to create a key create a file with name secret.str and add your key in his and system will identify.

## EXECUTING

```bash
$ go run main.go
```

OR

```bash
$ go build
$ ./restfulgo
```
## TEST OF ACCESS
After of all created open your browser and copy and paste the url below

Open browser or postman and enter with url: `http://localhost:4000/v1/products`

After `http://localhost:4000/v1/products/1` or number of the id generated in your database