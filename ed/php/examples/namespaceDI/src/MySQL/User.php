<?php

namespace Persistence;

use Domain\Interfaces\Persistence\UserInterface;

class User implements UserInterface
{
    public function getById($id)
    {
        echo "Gets user by id $id from MySQL.";
    }
}
