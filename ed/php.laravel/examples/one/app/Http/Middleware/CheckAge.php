<?php

namespace App\Http\Middleware;

use Closure;

class CheckAge
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
        return $next($request);
    }


    public function terminate($request, $response)
    {
        // Store the session data...
    }
}
