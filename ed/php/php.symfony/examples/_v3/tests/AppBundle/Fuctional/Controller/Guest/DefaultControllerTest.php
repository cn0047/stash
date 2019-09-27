<?php

namespace Tests\AppBundle\Functional\Controller\Guest;

use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

class DefaultControllerTest extends WebTestCase
{
    public function testIndex()
    {
        $client = static::createClient();
        $crawler = $client->request('GET', '/');

        $this->assertEquals(200, $client->getResponse()->getStatusCode());
        $this->assertTrue($client->getResponse()->isSuccessful());

        $this->assertContains('Welcome!', $crawler->filter('head')->text());

        $this->assertTrue($crawler->filter('body header a')->count() > 0);

        $this->assertContains('<code class="wrapper">HOME</code>', $crawler->filter('body > header')->html());
        $this->assertSame($crawler->filter('body header a')->first()->link()->getUri(), 'http://localhost/');

        $this->assertTrue($crawler->filter('html:contains("ACCOUNT")')->count() > 0);
        $this->assertContains('ACCOUNT', $client->getResponse()->getContent());
    }
}
