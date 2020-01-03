<?php

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;

$spamMaintainer = $app['controllers_factory'];

$spamMaintainer->get('/', function () use ($app) {
    try {
        // Parse log.
        $logParser = new Screecher\DependencyInjection\LogParser();
        $stats = $logParser->getAggregatedInfo();
        $apisNames = array_keys($stats);

        // By $apisNames  we'll receive apis joined with maintainers from db.

        // By using twig we can receive text of message,
        // we just need traverse array of apis and render in each mail an array of errors.

        // We have rendered mails and we have emails - it isn't problem to send mails)

        $result = [
            'success' => true,
        ];
        return $app->json($result, 200);
    } catch (Exception $e) {
        $result = [
            'success' => false,
            'error' => 'Internal server error.',
        ];
        return $app->json($result, 500);
    }
});

return $spamMaintainer;
