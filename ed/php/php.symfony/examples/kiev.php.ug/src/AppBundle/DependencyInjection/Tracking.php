<?php

namespace AppBundle\DependencyInjection;

class Tracking
{
    public function __construct()
    {
    }

    public function log($msg)
    {
    }

    public function ioTrack($metric, $value = 1)
    {
        $app_id = 9752;
        $key = '405e020214ff53747a488d16cb03d5ea';
        $track_url = "https://tapi.onthe.io/?k={$app_id}:{$metric}&s={$key}&v={$value}";
        file_get_contents($track_url);
    }
}
