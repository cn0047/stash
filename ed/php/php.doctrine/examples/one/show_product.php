<?php
/**
 * Run from shell:
 * php show_product.php
 * php show_product.php 1
 * php show_product.php 2
 */

require_once "bootstrap.php";

if (isset($argv[1])) {
    $id = $argv[1];
    $product = $entityManager->find('Product', $id);
    if ($product === null) {
        echo "No product found.\n";
        exit(1);
    }
    echo sprintf("-%s\n", $product->getName());
} else {
    require_once 'show_all_products.php';
}
