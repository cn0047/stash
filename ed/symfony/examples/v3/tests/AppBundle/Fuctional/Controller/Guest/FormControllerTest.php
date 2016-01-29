<?php

namespace Tests\AppBundle\Functional\Controller;

use Symfony\Bundle\FrameworkBundle\Test\WebTestCase;

class FormControllerTest extends WebTestCase
{
    public function testIndex()
    {
        $client = static::createClient();
        $crawler = $client->request('GET', '/my_csrf');
        static::assertEquals(200, $client->getResponse()->getStatusCode());
        static::assertTrue($client->getResponse()->isSuccessful());

        $buttonCrawlerNode = $crawler->selectButton('Submit');
        $form = $buttonCrawlerNode->form([
            'my_csrf[message]' => '007',
        ], 'POST');

        $crawler = $client->submit($form);

        static::assertEquals(
            'Form submitted with message: 007! Token valid: method 1 = true, method 2 = true',
            $crawler->filterXPath('//*[@id="message"]')->html()
        );
    }
}
