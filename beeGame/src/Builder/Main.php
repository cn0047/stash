<?php

namespace Builder;

use VO\PositiveInteger;
use VO\NotEmptyString;
use Bee\Gang;
use Bee\Factory;
use Config\ConfigInterface;

class Main
{
    // Specify particular types of bee that allowed in this particular builder.
    private static $allowedBeeTypes = [
        'Queen',
        'Worker',
        'Drone',
    ];

    /** @var ConfigInterface */
    private $config;

    /** @var Gang $beeGang */
    private $beeGang;

    public function __construct(NotEmptyString $level)
    {
        $configName = "Config\\$level";
        $this->config = new $configName();
    }

    public function buildLevel()
    {
        $factory = new Factory();
        $this->beeGang = new Gang();
        // Receive configs and create bees.
        foreach (self::$allowedBeeTypes as $beeName) {
            $lifespan = $this->config->get(new NotEmptyString("lifespan$beeName"));
            $deduceStep = $this->config->get(new NotEmptyString("deduceStep$beeName"));
            $count = $this->config->get(new NotEmptyString("count$beeName"));
            // Fill bee gang.
            for ($i = 0; $i < $count; $i++) {
                $this->beeGang->add($factory->create(
                    new NotEmptyString($beeName),
                    new PositiveInteger($lifespan),
                    new PositiveInteger($deduceStep)
                ));
            }
        }
        // This help improve random behaviour..
        $this->beeGang->shuffle();
    }

    public function getBeeGang()
    {
        return $this->beeGang;
    }
}
