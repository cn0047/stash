<?php

namespace App\Http\Middleware;

use Closure;

class ConsoleLog
{
    /**
     * Run the request filter.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @return mixed
     */
    public function handle($request, Closure $next)
    {
        // Here must be some stuff with $request.
        echo '<script>console.log('.json_encode($request).')</script>';
        return $next($request);
    }


    public function terminate($request, $response)
    {
        // Store the session data...
    }
}
