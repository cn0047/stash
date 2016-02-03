<?php

namespace Config;

use VO\NotEmptyString;

abstract class Config implements ConfigInterface
{
    protected static $config = [];

    public function get(NotEmptyString $configName)
    {
        $key = $configName->get();
        if (!array_key_exists($key, static::$config)) {
            throw new \InvalidArgumentException('Config not found.');
        }
        return static::$config[$key];
    }
}
