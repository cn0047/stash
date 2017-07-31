<?php
/**
 * Abstract factory
 *
 * Abstract Factory offers the interface for creating a family of related objects,
 * without explicitly specifying their classes.
 * For example: DbProviderFactory (mysql/mongo/etc) - DbCommand, DbConnection, DbParameter, etc.
 *
 * @category Creational
 */

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
