Abstract factory
-
1

````php
abstract class Document
{
   abstract public function createPage();
}

class PortraitDocument extends Document
{
   public function createPage()
   {
      return new PortraitPage;
   }
}

class LandscapeDocument extends Document
{
   public function createPage()
   {
      return new LandscapePage;
   }
}
````