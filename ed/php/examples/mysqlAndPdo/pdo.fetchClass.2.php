<?php

/*
create table orders (order_id int, product int, quantity int);
create table products (product_id int, name varchar(50), price int);
insert into products values(1, 'iphone', 1500);
insert into orders values (123, 1, 2);
*/

// docker
$dbh = new PDO('mysql:host=mysql-master;port=3306;dbname=test', 'dbu', 'dbp');

class Order {
    private $order_id;
    private $product_id;
    private $quantity;

    public function __set($name, $value) {
        if ($name === 'product') {
            $this->setProduct($value);
        }
    }

    public function setProduct($productId)
    {
        /** \PDO */
        global $dbh;
        $s = $dbh->prepare('SELECT * FROM products WHERE product_id = :pId');
        $s->bindParam(':pId', $productId, PDO::PARAM_INT);
        $s->execute();
        $this->product = $s->fetchAll(PDO::FETCH_CLASS, Product::class);
        $this->product = $this->product[0];
    }
}

class Product {
    private $product_id;
    private $name;
    private $price;
}

$s = $dbh->prepare('SELECT * FROM orders WHERE order_id = :oId');
$s->bindValue(':oId', 123, PDO::PARAM_INT);
$s->execute();
$result = $s->fetchAll(PDO::FETCH_CLASS, Order::class);
var_export($result[0]);
