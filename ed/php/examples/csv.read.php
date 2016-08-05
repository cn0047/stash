<?php

$csv = array_map('str_getcsv', file('/home/kovpak/csv.csv'));
var_export($csv);
