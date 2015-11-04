<?php

namespace AppBundle\DependencyInjection;

use AppBundle\Entity\LogPointer;

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

    final private function getPatterForParsing()
    {
        $patterns = [
            'default' => '/^(.*) - (.*) \[(.*)\] "(.*)" (\d+) (\d+) ".*" "(.*)"$/',
        ];
        if (!isset($patterns[$this->format])) {
            throw new \DomainException("Unknown format: {$this->format}");
        }
        return $patterns[$this->format];
    }

    final private function readLogFile(\SplFileInfo $file)
    {
        $filePath = $file->getRealPath();
        if (!is_readable($filePath)) {
            throw new \RuntimeException("Can't read file: $filePath");
        }
        $logPointer = $this->em
            ->getRepository('AppBundle:LogPointer')
            ->findOneByFile($filePath)
        ;
        if (is_null($logPointer)) {
            $logPointer = new LogPointer();
            $logPointer->setFile($filePath);
            $this->em->persist($logPointer);
        }
        $this->em->getConnection()->beginTransaction();
        try {
            $pattern = $this->getPatterForParsing();
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
                    throw new \RuntimeException("Can't parse line: $line");
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
