<?php

namespace App\Http\Middleware;

use Closure;

class GlobalBeforeMiddleware
{
    public function handle($request, Closure $next)
    {
        // echo __METHOD__;
        return $next($request);
    }
}
