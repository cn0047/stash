<?php

$data = [
    ['id' => '154', 'signUpCountry' => 'United Kingdom'],
    ['id' => '291', 'signUpCountry' => 'Ukraine'],
    ['id' => '885', 'signUpCountry' => 'United Kingdom'],
    ['id' => '1086', 'signUpCountry' => 'Ukraine'],
    ['id' => '1976', 'signUpCountry' => 'United Kingdom'],
    ['id' => '1994', 'signUpCountry' => 'Ukraine'],
    ['id' => '3080', 'signUpCountry' => 'Ukraine'],
    ['id' => '3119', 'signUpCountry' => 'India'],
    ['id' => '3782', 'signUpCountry' => 'Malaysia'],
    ['id' => '4169', 'signUpCountry' => 'China'],
    ['id' => '4443', 'signUpCountry' => 'Pakistan'],
    ['id' => '4475', 'signUpCountry' => 'United Kingdom'],
    ['id' => '4646', 'signUpCountry' => 'Ukraine'],
    ['id' => '4769', 'signUpCountry' => 'Bangladesh'],
    ['id' => '4821', 'signUpCountry' => 'India'],
];
usort($data, function ($a, $b) {
    // ASC
    return strcmp($a['signUpCountry'], $b['signUpCountry']);
    // DESC
    // return strcmp($b['signUpCountry'], $a['signUpCountry']);
});
var_export($data);