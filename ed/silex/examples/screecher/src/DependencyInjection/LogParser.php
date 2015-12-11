<?php

namespace Screecher\DependencyInjection;

class LogParser
{
    private $fileName = 'api_usage.log';

    public function getPath()
    {
        return __DIR__.'/../../'.$this->fileName;
    }

    public function getAggregatedInfo()
    {
        $file = $this->getPath();
        if (!file_exists($file)) {
            throw new \DomainException('Log file not found.');
        }
        $stats = [];
        if (($handle = fopen($file, 'r')) !== false) {
            while (($data = fgetcsv($handle, 1000, ",")) !== false) {
                list($date, $time, $apiName, $status, $rawError) = $data;
                if ($status === 'error') {
                    foreach (explode('.', $rawError) as $error) {
                        $error = trim($error);
                        if (!empty($error)) {
                            if (isset($stats[$apiName][$error])) {
                                $stats[$apiName][$error]++;
                            } else {
                                $stats[$apiName][$error] = 1;
                            }
                        }
                    }
                }
            }
            fclose($handle);
        }
        return $stats;
    }
}
