<?php

namespace Config;

use VO\NotEmptyString;

interface ConfigInterface
{
    public function get(NotEmptyString $configName);
}