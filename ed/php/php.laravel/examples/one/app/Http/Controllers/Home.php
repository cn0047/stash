<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Requests;
use Illuminate\Contracts\Validation\Validator;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;
use Illuminate\Support\Facades\App;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\DB;
use App\Events\OrderShipped;

class Home extends BaseController
{
    use DispatchesJobs, ValidatesRequests;

    public function __construct()
    {
        $this->middleware('auth.basic', ['only' => 'auth']);
    }

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

    public function db()
    {
        // simple part
        $dbh = new \PDO('mysql:host=mysql-master;port=3306;dbname=test', 'dbu', 'dbp');
        $s = $dbh->prepare('SELECT NOW()');
        $s->execute();
        $result = $s->fetchAll(\PDO::FETCH_ASSOC);
        var_export($result);

        echo '<hr>';

        var_dump(DB::connection()->getPdo());
    }

    public function userFromDb()
    {
        $user = DB::select('select * from users where email = ?', ['cnfxlr@gmail.com']);
        var_export($user);
    }

    public function auth()
    {
        Auth::onceBasic();
        echo 200;
    }

    public function event()
    {
        event(new OrderShipped(100));
    }
}
