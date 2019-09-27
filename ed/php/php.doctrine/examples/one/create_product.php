<?php
/**
 * Run from shell:
 * php create_product.php ORM
 * php create_product.php DBAL
 * php create_product.php DAO
 */

require_once "bootstrap.php";

$newProductName = $argv[1];
$product = new Product();
$product->setName($newProductName);
$entityManager->persist($product);
$entityManager->flush();
echo "Created Product with ID " . $product->getId() . "\n";
