<?php
/**
 * Run from shell:
 * php show_product.php
 * php show_product.php 1
 * php show_product.php 2
 */

require_once 'bootstrap.php';

$productRepository = $entityManager->getRepository('Product');
$products = $productRepository->findAll();

require_once 'render_products.php';
