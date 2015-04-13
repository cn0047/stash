<?php

echo 'Are you sure you want to do this [y/n]';
$confirmation = trim(fgets(STDIN));
if ($confirmation !== 'y') {
    echo "400 \n";
    exit;
}
echo "200 \n";