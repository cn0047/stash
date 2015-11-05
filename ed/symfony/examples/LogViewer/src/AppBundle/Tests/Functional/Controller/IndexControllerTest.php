<?php

namespace AppBundle\Tests\Controller;

use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

class IndexControllerTest extends WebTestCase
{
    public function testIndex()
    {
        $client = static::createClient();
        $crawler = $client->request('GET', '/', [], [], ['PHP_AUTH_USER' => 'user','PHP_AUTH_PW' => 'user']);
        $this->assertEquals(200, $client->getResponse()->getStatusCode());
        $this->assertContains('Items per page', $crawler->filter('body')->text());
    }
}
