<?php

class Host
{
    private $ip = '';
    private $name = '';

    public function __construct($ip, $name)
    {
        $this->ip = $ip;
        $this->name = $name;
    }

    public function __get($name)
    {
        if (isset($this->$name)) {
            return $this->$name;
        }
        throw new DomainException("Unknown property $name");
    }
}

$h = new Host('ip', 'name');
var_export([
    $h->name,
    $h->ip,
]);
$h->code;

/*
array (
  0 => 'name',
  1 => 'ip',
)
Fatal error: Uncaught exception 'DomainException' with message 'Unknown property code'
*/
