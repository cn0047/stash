<?php

namespace Persistence;

use Domain\Interfaces\Persistence\User as UserInterface;

class User implements UserInterface
{
    public function getById(int $id) : array
    {
        return ['desc' => "User from MySQL with id: $id ."];
    }
}
