<?php

namespace Config;

/**
 * Configuration settings for game level 1.
 */
class LevelOne extends Config
{
    protected static $config = [
        'lifespanQueen' => 100,
        'deduceStepQueen' => 8,
        'countQueen' => 1,

        'lifespanWorker' => 75,
        'deduceStepWorker' => 10,
        'countWorker' => 5,

        'lifespanDrone' => 50,
        'deduceStepDrone' => 12,
        'countDrone' => 8,
    ];
}
