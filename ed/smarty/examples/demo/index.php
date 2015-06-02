<?php

require_once 'vendor/autoload.php';

$smarty = new Smarty;
$smarty->assign("data", ["New York", "Nebraska", "Kansas", "Iowa", "Oklahoma", "Texas", "Doe", "Smith", "Johnson", "Case", "John", "Mary", "James", "Henry"]);
$smarty->display('index.tpl');
