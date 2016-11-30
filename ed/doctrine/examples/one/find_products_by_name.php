<?php
/**
 * Run from shell:
 * php show_product.php
 * php show_product.php 1
 * php show_product.php 2
 */

require_once "bootstrap.php";

if (isset($argv[1])) {
    $name = $argv[1];
    $product = $entityManager->getRepository('Product')->findBy(['name' => $name]);
    require_once 'render_product.php';
} else {
    require_once 'show_all_products.php';
}
