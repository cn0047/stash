<?php

require_once 'vendor/autoload.php';

$smarty = new Smarty;
$smarty->assign("skill_abbreviation", "HULK");
$smarty->assign("skill_abbreviations", ["HULK", "MAC"]);
$smarty->assign("skill_HULK", 5);
$smarty->assign("skill_MAC", 2);
$smarty->display('index.2.tpl');
