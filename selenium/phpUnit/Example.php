<?php
class Example extends PHPUnit_Extensions_SeleniumTestCase
{
  protected function setUp()
  {
    $this->setBrowser("*chrome");
    $this->setBrowserUrl("http://www.007.com/");
  }

  public function testMyTestCase()
  {
    $this->open("http://www.007.com/");
    for ($second = 0; ; $second++) {
        if ($second >= 60) $this->fail("timeout");
        try {
            if ("News" == $this->getText("link=News")) break;
        } catch (Exception $e) {}
        sleep(1);
    }

    $this->click("link=News");
    $this->waitForPageToLoad("30000");
    for ($second = 0; ; $second++) {
        if ($second >= 60) $this->fail("timeout");
        try {
            if ("" == $this->getText("css=div.entry_link > a[title=\"Permalink to BOND IN MOTION COMES TO LONDON\"]")) break;
        } catch (Exception $e) {}
        sleep(1);
    }

    $this->click("css=div.entry_link > a[title=\"Permalink to BOND IN MOTION COMES TO LONDON\"]");
    $this->waitForPageToLoad("30000");
    for ($second = 0; ; $second++) {
        if ($second >= 60) $this->fail("timeout");
        try {
            if ("Bond in Motion, the largest official collection of James Bond vehicles, will be on display at the London Film Museum from 21 March. The exhibition contains more than 100 pieces including models, production art, props and iconic vehicles including ‘Little Nellie’ from YOU ONLY LIVE TWICE, Goldfinger’s Rolls-Royce Phantom III and the 1964 Aston Martin DB5 from GOLDENEYE. Also on display, for the first time, will be the 1/3 scale model of the Agusta Westland AW101 helicopter used in the filming of SKYFALL. For more information and tickets go to: www.londonfilmmuseum.com and www.ticketmaster.co.uk." == $this->getText("//div[@id='sidebar_content_top']/div/div[3]/p[2]")) break;
        } catch (Exception $e) {}
        sleep(1);
    }
    $this->assertTitle('The Official James Bond 007 Website | BOND IN MOTION COMES TO LONDON');

  }
}
?>