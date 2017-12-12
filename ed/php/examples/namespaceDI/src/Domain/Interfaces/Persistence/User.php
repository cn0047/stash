<?php

namespace Domain\Interfaces\Persistence;

interface User
{
    public function getById(int $id) : array;
}
