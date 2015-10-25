<?php
/**
 * php update_product.php 1 new-ORM
 * php update_product.php 2 new-DBAL
 * php update_product.php 3 new-DAO
 */

require_once "bootstrap.php";

$id = $argv[1];
$newName = $argv[2];
$product = $entityManager->find('Product', $id);
if ($product === null) {
    echo "Product $id does not exist.\n";
    exit(1);
}
$product->setName($newName);
$entityManager->flush();
