package search

import (
  "fmt"
  "github.com/algolia/algoliasearch-client-go/algoliasearch"

  "content/article.v1/models"
  config "deploys/config/search/algolia"
)

func AddIntoIndex(article models.Article) error {
  client := algoliasearch.NewClient(config.AppID, config.APIKey)
  index := client.InitIndex("articles")

  object := algoliasearch.Object{
    "objectID":   article.GoID,
    "title":      article.Title,
    "headline":   article.Headline,
    "promoImage": article.PromoImage,
    "tags":       article.TagSlugs,
    "channel":    article.ChannelSlug,
    "section":    article.SectionSlug,
    "siteID":     article.SiteID,
  }

  _, err := index.UpdateObject(object)
  if err != nil {
    return fmt.Errorf("algolia, failed to perform UpdateObject, error: %v", err)
  }

  return nil
}
