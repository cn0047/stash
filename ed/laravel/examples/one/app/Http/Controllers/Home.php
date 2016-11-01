<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Requests;
use Illuminate\Contracts\Validation\Validator;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Facades\App;

class Home extends BaseController
{
    use DispatchesJobs, ValidatesRequests;

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

    public function flash(Request $request)
    {
        $request->session()->flash('status', 'Task was successful!');
        if ($request->session()->has('status')) {
            echo $request->session()->get('status');
        }
    }

    public function validation(Request $request)
    {
        $this->validate($request, [
            'title' => 'required|min:2',
        ]);
    }

    protected function formatValidationErrors(Validator $validator)
    {
        var_export($validator->errors()->all());
    }

    public function localization()
    {
        echo App::getLocale();
        echo trans('messages.welcome');
    }
}
