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
        if ($this->product === null) {
            $this->product = new Product();
        }
        switch ($name) {
            case 'productId':
                 $this->product->setProductId($value);
                break;
            case 'productName':
                 $this->product->setName($value);
                break;
            case 'productPrice':
                 $this->product->setPrice($value);
                break;
        }
    }
}

class Product {
    private $product_id;
    private $name;
    private $price;

    public function setProductId($product_id)
    {
        $this->product_id = $product_id;
    }

    public function setName($name)
    {
        $this->name = $name;
    }

    public function setPrice($price)
    {
        $this->price = $price;
    }
}

$s = $dbh->prepare('
    SELECT o.*, p.product_id productId, p.name productName, p.price productPrice
    FROM orders o
    JOIN products p ON o.product = p.product_id
    WHERE order_id = :oId
');
$s->bindValue(':oId', 123, PDO::PARAM_INT);
$s->execute();
$result = $s->fetchAll(PDO::FETCH_CLASS, Order::class);
var_export($result[0]);
