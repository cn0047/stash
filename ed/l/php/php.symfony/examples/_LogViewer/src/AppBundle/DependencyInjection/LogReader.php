<?php

namespace AppBundle\DependencyInjection;

use AppBundle\Entity\LogPointer;

/**
 * Class that provides methods to read logs.
 */
class LogReader
{
    private $em;
    private $logSaver;
    private $format = '';
    private $path = '';

    public function __construct($em, $logSaver, $format, $path)
    {
        $this->em = $em;
        $this->logSaver = $logSaver;
        $this->format = $format;
        $this->path = $path;
    }

    /**
     * Recursively find log files in directory and read them.
     * @return array Array with info about found log files.
     */
    public function readLogs()
    {
        $items = new \RecursiveIteratorIterator(
            new \RecursiveDirectoryIterator($this->path),
            \RecursiveIteratorIterator::SELF_FIRST
        );
        $result = [];
        foreach ($items as $filePath => $file) {
            $basename = $file->getBasename();
            if ($basename === '.' or $basename === '..') {
                continue;
            }
            if ($file->isDir()) {
                continue;
            }
            if (!preg_match('/^access\.log$/', $basename)) {
                continue;
            }
            $result[] = $this->readLogFile($file);
        }
        return $result;
    }

    /**
     * Provides ability to change pattern for parsing data from log file.
     * @throw \DomainException In case when is service config setted unknown pattern.
     * @return string RegExp pattern.
     */
    final private function getPatterForParsing()
    {
        $patterns = [
            'default' => '/^(.*) - (.*) \[(.*)\] "(.*)" (\d+) (\d+) ".*" "(.*)"$/',
        ];
        if (!isset($patterns[$this->format])) {
            throw new \DomainException("Unknown format: {$this->format}.");
        }
        return $patterns[$this->format];
    }

    /**
     * Read particular log file and send data to log saver (to store data at db).
     * @param \SplFileInfo $file Log file.
     * @throw \RuntimeException If log file not readable, or something going wrong.
     * @return string Info about how much data was read.
     */
    final private function readLogFile(\SplFileInfo $file)
    {
        $filePath = $file->getRealPath();
        if (!is_readable($filePath)) {
            throw new \RuntimeException("Can't read file: $filePath.");
        }
        // This pointer contains info
        // about how much information was read
        // from log file during last time file read.
        $logPointer = $this->em
            ->getRepository('AppBundle:LogPointer')
            ->findOneByFile($filePath)
        ;
        if (is_null($logPointer)) {
            $logPointer = new LogPointer();
            $logPointer->setFile($filePath);
            $this->em->persist($logPointer);
        }
        // In case when something go wrong,
        // we don't need to save logs to db
        // and should avoid saving pointer too.
        $this->em->getConnection()->beginTransaction();
        try {
            $pattern = $this->getPatterForParsing();
            // Update pointer.
            $pointer = $logPointer->getPointer();
            $size = filesize($filePath);
            $logPointer->setPointer($size);
            $this->em->flush();
            $f = fopen($filePath, 'r');
            if ($pointer <= $size) {
                fseek($f, $pointer);
            }
            while (($line = fgets($f)) !== false) {
                if (preg_match($pattern, $line, $matches)) {
                    list(
                        $allLine,
                        $host,
                        $user,
                        $dateTime,
                        $firstRequestLine,
                        $status,
                        $size,
                        $userAgent
                    ) = $matches;
                    // Push parsed line to the saver.
                    $this->logSaver->save([
                        'owner' => basename($file->getPath()),
                        'host' => $host,
                        'user' => $user,
                        'dateTime' => new \DateTime($dateTime),
                        'firstRequestLine' => $firstRequestLine,
                        'status' => $status,
                        'size' => $size,
                        'userAgent' => $userAgent,
                    ]);
                } else {
                    throw new \RuntimeException("Can't parse line: $line.");
                }
            }
            $this->logSaver->flush();
            fclose($f);
            $this->em->getConnection()->commit();
        } catch (\Exception $e) {
            $this->em->getConnection()->rollback();
            throw new \RuntimeException($e->getMessage());
        }
        return "File: $filePath, last read: $pointer, new read: $size.";
    }
}
