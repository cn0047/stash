<?php

namespace Config;

use VO\NotEmptyString;

/**
 * Game level config.
 *
 * This config provides convenient way to work with settings for bees in particular game level.
 */
abstract class Config
{
    protected static $config = [];

    /**
     * Gets setting value by name.
     *
     * @param NotEmptyString $configName Name of setting.
     *
     * @throws \InvalidArgumentException In case when config not found.
     *
     * @return mixed Value of setting.
     */
    public function get(NotEmptyString $configName)
    {
        $key = $configName->get();
        if (!array_key_exists($key, static::$config)) {
            throw new \InvalidArgumentException('Config not found.');
        }
        return static::$config[$key];
    }
}
