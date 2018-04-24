# restfulgo
CREATING A SIMPLE API USING GOLANG

<h1>CLONE</h1>
<pre>git clone git@github.com:DiegoSantosWS/restfulgo.git</pre>

<h1>CREATING DATABASE</h1>

<pre>
    CREATE A DATABASE AND AFTER EXECUTE THE CODE BELOW.

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
</pre>

<h1>SECRET KEY</h1>

Api requires a serect key to toke validation, to create a key create a file with name secret.str and add your key in his and system will identify.

<h1>EXECUTING</h1>
<pre>go run *.go</pre>
<h4>ACCESS</h4>
After of all created open your browser and copy and paste the url below

<pre>http://localhost:4000/v1/products</pre>
<pre>http://localhost:4000/v1/products/1 or number of the id</pre>