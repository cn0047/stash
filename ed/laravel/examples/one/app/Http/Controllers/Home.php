<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Requests;

class Home extends Controller
{
    public function index($code)
    {
        echo $code;
    }

    public function home()
    {
        return view('front.home', ['time' => time()]);
    }

    public function conf()
    {
        echo '<pre>';
        var_export([
            config('app.timezone'),
            \App::environment(),
        ]);
        // Print log to ./storage/logs/laravel.log.
        \Log::info('Showed configs.');
    }

    public function di()
    {
    }
}
