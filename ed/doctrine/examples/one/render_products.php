<?php

foreach ($products as $product) {
    echo sprintf("%d: %s\n", $product->getId(), $product->getName());
}
