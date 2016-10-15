<?php

namespace App\Providers;

use Illuminate\Support\ServiceProvider;

class RiakServiceProvider extends ServiceProvider
{
    /**
     * Bootstrap the application services.
     *
     * @return void
     */
    public function boot()
    {
        var_dump(__METHOD__);
        var_export(func_get_args());
    }

    /**
     * Register the application services.
     *
     * @return void
     */
    public function register()
    {
        var_dump(__METHOD__);
        var_export(func_get_args());
    }
}
