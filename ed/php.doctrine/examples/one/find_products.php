<?php
/**
 * Run from shell:
 * php show_product.php
 * php show_product.php 1
 * php show_product.php 2
 */

require_once "bootstrap.php";

if (isset($argv[1])) {
    $ids = explode(',', $argv[1]);
    $dql = "SELECT p FROM Product p WHERE p.id IN (?1)";
    $q = $entityManager->createQuery($dql)->setParameter(1, $ids);
    $products = $q->execute();
    require_once 'render_products.php';
} else {
    require_once 'show_all_products.php';
}
