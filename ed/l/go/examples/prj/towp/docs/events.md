# wp service events

## Order's status updated

Topic: `uat.wp.events, dev-nonprod.wp.events, qai-nonprod.wp.events, prod.wp.events`

Event-type: `wp.Created`

Message schema:

````json
{
  "$schema": "https://json-schema.org/draft/2019-09/schema",
  "type": "object",
  "properties": {
    "attributes": {
      "type": "object",
      "properties": {
        "env_type": {
          "type": "string"
        },
        "retailer_id": {
          "type": "string"
        },
        "mfc_id": {
          "type": "string"
        },
        "event_type": {
          "type": "string",
          "enum": ["wp.Created"]
        },
        "source": {
          "type": "string",
          "enum": ["wp"]
        }
      },
      "required": ["env_type","retailer_id","mfc_id","event_type","source"]
    },
    "data": {
      "type": "object",
      "properties": {
        "created_time": {
          "type": "string"
        },
        "wave_plan": {
          "type": "object",
          "properties": {
            "id": {
              "type": "string"
            },
            "retailer_id": {
              "type": "string"
            },
            "mfc_id": {
              "type": "string"
            },
            "timezone": {
              "type": "string"
            },
            "created_time": {
              "type": "string"
            },
            "created_by": {
              "type": "string"
            },
            "waves": {
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "id": {
                    "type": "string"
                  },
                  "cutoff_time": {
                    "type": "string"
                  },
                  "from_time": {
                    "type": "string"
                  },
                  "to_time": {
                    "type": "string"
                  },
                  "schedules": {
                    "type": "array",
                    "items": {
                      "type": "object",
                      "properties": {
                        "id": {
                          "type": "string"
                        },
                        "schedule_type": {
                          "type": "string"
                        },
                        "schedule_time": {
                          "type": "string"
                        }
                      },
                      "required": ["id","schedule_type","schedule_time"]
                    }
                  }
                },
                "required": ["id","cutoff_time","from_time","to_time","schedules"]
              }
            }
          },
          "required": ["id","retailer_id","mfc_id","timezone","created_time","created_by","waves"]
        }
      },
      "required": ["created_time","wave_plan"]
    }
  },
  "required": ["attributes","data"]
}
````

Source code for message you can find
[here](https://github.com/to-com/wp/blob/fbef4dca6f9a80d97d6986ae7bec0f8334324b75/internal/dto/dto.go#L65:L68)
and [here](https://github.com/to-com/wp/blob/fbef4dca6f9a80d97d6986ae7bec0f8334324b75/internal/business/business.go#L67:L78).

Example:

````json
{
  "attributes": {
    "env_type": "dev",
    "retailer_id": "MAF",
    "mfc_id": "DO2",
    "event_type": "wp.Created",
    "source": "wp"
  },
  "data": {
    "created_time": "2022-11-15T17:00:00Z",
    "wave_plan": {
      "id": "14d4d27a-1cd2-4190-aead-c3c965c78727",
      "retailer_id": "MAF",
      "mfc_id": "DO2",
      "timezone": "Asia/Dubai",
      "created_time": "2022-11-15T17:00:00Z",
      "created_by": "ng0ao6VyeJevxlQ6aduorIxKzBm0",
      "waves": [
        {
          "id": "51000d18-8454-45d1-b348-b011148b4a2d",
          "cutoff_time": "17:00",
          "from_time": "00:00",
          "to_time": "12:59",
          "schedules": [
            {
              "id": "5ea8b432-77da-4ca2-aa3e-42512f1f6bd0",
              "schedule_type": "delta_picklist",
              "schedule_time": "16:00"
            }
          ]
        }
      ]
    }
  }
}
````

Event-type: `wp.TriggersFired`

Message schema:

````json
{
  "$schema": "https://json-schema.org/draft/2019-09/schema",
  "type": "object",
  "properties": {
    "attributes": {
      "type": "object",
      "properties": {
        "env_type": {
          "type": "string"
        },
        "retailer_id": {
          "type": "string"
        },
        "mfc_id": {
          "type": "string"
        },
        "event_type": {
          "type": "string",
          "enum": ["wp.TriggersFired"]
        },
        "source": {
          "type": "string",
          "enum": ["wp"]
        }
      },
      "required": ["env_type","retailer_id","mfc_id","event_type","source"]
    },
    "data": {
      "type": "object",
      "properties": {
        "schedule_type": {
          "type": "string"
        },
        "cutoffs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      },
      "required": ["schedule_type","cutoffs"]
    }
  },
  "required": ["attributes","data"]
}
````

Source code for message you can find
[here](https://github.com/to-com/wp/blob/7d3f9befb10a6f23c10571dc5bf659d7d42c0281/internal/dto/dto.go#L108:L111)
and [here](https://github.com/to-com/wp/blob/7d3f9befb10a6f23c10571dc5bf659d7d42c0281/internal/business/business.go#L305:L316).

Example:

````json
{
  "attributes": {
    "env_type": "dev",
    "retailer_id": "MAF",
    "mfc_id": "DO2",
    "event_type": "wp.TriggersFired",
    "source": "wp"
  },
  "data": {
    "schedule_type":"prelim_picklist",
    "cutoffs": [
      "2022-11-15T17:00:00Z"
    ]
  }
}
````
