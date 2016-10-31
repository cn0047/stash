<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;
use App\Services\Helper;

class RiakServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap the application services.
     *
     * @return void
     */
    public function boot()
    {
        $this->app->bind('App\Services\Helper', function ($app) {
            return new Helper();
        });

        $this->app->tag(['App\Services\Helper'], 'myFromRiakServiceProvider');
    }

    /**
     * Register the application services.
     *
     * @return void
     */
    public function register()
    {
        $h = $this->app->make('App\Services\Helper');
    }
}
