<?php

header('Content-Type: text/csv; charset=utf-8');
header('Content-Disposition: attachment; filename=file.csv');
$output = fopen('php://output', 'w');
fputcsv(
    $output,
    [
        'Id',
        'User id',
    ]
);
foreach ([[1, 'James Bond'], [2, 'Felix Leiter']] as $row) {
    fputcsv(
        $output,
        $row
    );
}
