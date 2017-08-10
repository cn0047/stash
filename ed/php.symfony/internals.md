Internals
-

````php
$profile = $container->get('profiler')->loadProfileFromResponse($response);
$profile = $container->get('profiler')->loadProfile($token);
// get the latest 10 tokens
$tokens = $container->get('profiler')->find('', '', 10, '', '');
// get the latest 10 tokens for all URL containing /admin/
$tokens = $container->get('profiler')->find('', '/admin/', 10, '', '');
// get the latest 10 tokens for local requests
$tokens = $container->get('profiler')->find('127.0.0.1', '', 10, '', '');
// get the latest 10 tokens for requests that happened between 2 and 4 days ago
$tokens = $container->get('profiler')->find('', '', 10, '4 days ago', '2 days ago');

// on the production machine
$profile = $container->get('profiler')->loadProfile($token);
$data = $profiler->export($profile);
// on the development machine
$profiler->import($data);

// Configuration
# load the profiler
framework:
    profiler: { only_exceptions: false }
# enable the web profiler
web_profiler:
    toolbar: true
    intercept_redirects: true

_profiler:
    resource: "@WebProfilerBundle/Resources/config/routing/profiler.xml"
    prefix: /_profiler
````
