<?php

namespace Builder;

use VO\PositiveInteger;
use VO\NotEmptyString;
use Bee\Gang;
use Bee\Factory;
use Config\Config;

/**
 * Main builder.
 *
 * This is builder that build all things necessary for game.
 * Encapsulates all game initialization logic.
 */
class Main
{
    /** @var array $allowedBeeTypes Specify particular types of bee that allowed in this particular builder. */
    private static $allowedBeeTypes = [
        'Queen',
        'Worker',
        'Drone',
    ];

    /** @var Config $config Contains configuration of game for particular game level. */
    private $config;

    /** @var Gang $beeGang Contains bee aggregate.*/
    private $beeGang;

    /**
     * Constructor.
     *
     * @param NotEmptyString $level Contains level by which will be loaded configs.
     */
    public function __construct(NotEmptyString $level)
    {
        $configName = "Config\\$level";
        $this->config = new $configName();
    }

    /**
     * Builds all game objects.
     *
     * @throws \InvalidArgumentException In case when factory cannot create bee.
     */
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

    /**
     * Provides access to bee gang aggregate.
     *
     * @return Gang Bee gang.
     */
    public function getBeeGang()
    {
        return $this->beeGang;
    }
}
