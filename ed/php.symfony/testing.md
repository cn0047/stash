Testing
-

````php
phpunit -c app/
phpunit -c app src/AppBundle/Tests/Util

$client = static::createClient();
$crawler = $client->request('GET', '/post/hello-world');
$client->request('POST', '/submit', array('name' => 'Fabien'));
$client->request(
    'POST',
    '/submit',
    array(),
    array(),
    array('CONTENT_TYPE' => 'application/json'),
    '{"name":"Fabien"}'
);

$client->request(
    $method,
    $uri,
    array $parameters = array(),
    array $files = array(),
    array $server = array(),
    $content = null,
    $changeHistory = true
);

$link = $crawler->selectLink('Go elsewhere...')->link();
$crawler = $client->click($link);

$form = $crawler->selectButton('validate')->form();
$crawler = $client->submit($form, array('name' => 'Fabien'));

$link = $crawler
    ->filter('a:contains("Greet")') // find all links with the text "Greet"
    ->eq(1) // select the second link in the list
    ->link() // and click it
    ;
$crawler = $client->click($link);

$form = $crawler->selectButton('submit')->form();
// set some values
$form['name'] = 'Lucas';
$form['form_name[subject]'] = 'Hey there!';
// submit the form
$crawler = $client->submit($form);

// Assert that the response matches a given CSS selector.
$this->assertGreaterThan(0, $crawler->filter('h1')->count());

$this->assertContains(
    'Hello World',
    $client->getResponse()->getContent()
);

// Assert that there is at least one h2 tag with the class "subtitle"
$this->assertGreaterThan(
    0,
    $crawler->filter('h2.subtitle')->count()
);

// Assert that there are exactly 4 h2 tags on the page
$this->assertCount(4, $crawler->filter('h2'));

// Assert that the "Content-Type" header is "application/json"
$this->assertTrue(
$client->getResponse()->headers->contains(
    'Content-Type',
    'application/json'
);

// Assert that the response content contains a string
$this->assertContains('foo', $client->getResponse()->getContent());
// ...or matches a regex
$this->assertRegExp('/foo(bar)?/', $client->getResponse()->getContent());

// Assert that the response status code is 2xx
$this->assertTrue($client->getResponse()->isSuccessful());
// Assert that the response status code is 404
$this->assertTrue($client->getResponse()->isNotFound());
// Assert a specific 200 status code
$this->assertEquals(
    200, // or Symfony\Component\HttpFoundation\Response::HTTP_OK
    $client->getResponse()->getStatusCode()
);

// Assert that the response is a redirect to /demo/contact
$this->assertTrue(
    $client->getResponse()->isRedirect('/demo/contact')
);
// ...or simply check that the response is a redirect to any URL
$this->assertTrue($client->getResponse()->isRedirect());

// Form submission with a file upload
use Symfony\Component\HttpFoundation\File\UploadedFile;
$photo = new UploadedFile(
    '/path/to/photo.jpg',
    'photo.jpg',
    'image/jpeg',
    123
);
$client->request(
    'POST',
    '/submit',
    array('name' => 'Fabien'),
    array('photo' => $photo)
);

// Perform a DELETE request and pass HTTP headers
$client->request(
    'DELETE',
    '/post/12',
    array(),
    array(),
    array('PHP_AUTH_USER' => 'username', 'PHP_AUTH_PW' => 'pa$$word')
);
````
a
