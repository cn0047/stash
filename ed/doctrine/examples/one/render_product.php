<?php

if ($product === null) {
    echo "No product found.\n";
} else {
    $product = $product[0];
    echo sprintf("%d: %s\n", $product->getId(), $product->getName());
}
