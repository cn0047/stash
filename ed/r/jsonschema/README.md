JSON-Schema
-

https://spacetelescope.github.io/understanding-json-schema/
https://github.com/epoberezkin/ajv-keywords

````js
{"type": "string"}
{"type": "array", "uniqueItems": true}
{"type": "string", "enum": ["red", "amber", "green"]}

{"id": "http://yourdomain.com/schemas/myschema.json"}
{"$schema": "http://json-schema.org/draft-04/schema#"}
{"$ref": "#/definitions/address"}

{"allOf": [{"type": "string"}, {"maxLength": 5}]}
{"anyOf": [{"type": "string"}, {"type": "number"}]} // valid against any (one or more)
{"oneOf": [{"type": "number", "multipleOf": 5}, {"type": "number", "multipleOf": 3}]} // valid against exactly one
{"not": {"type": "string"}}
````

````js
// dependencies
{
  "type": "object",
  "properties": {
    "name": {"type": "string"},
    "credit_card": {"type": "number"},
    "billing_address": {"type": "string"}
 },
  "required": ["name"],
  "dependencies": {
    "credit_card": ["billing_address"]
 }
}

const schema = {
  type: 'object',
  required: [
    // 'multicastAddressBlock',
    'isAutoMulticast',
 ],
  properties: {
    multicastAddressBlock: {
      oneOf: [
        {type: 'null'},
        {type: 'string', format: 'ipv4'},
      ],
    },
    isAutoMulticast: {
      type: 'boolean',
    },
 },
  switch: [
    {
      if: {
        properties: {
          isAutoMulticast: {enum: [false]},
        },
      },
      then: {
        required: ['multicastAddressBlock'],
        properties: {
          multicastAddressBlock: {
            not: {type: 'null'},
          },
        },
      },
      continue: true,
    },
 ],
};

const schema = {
  type: 'object',
  required: [
    'interface',
    'sipLogging',
    'sipPort',
    'enableDtmfInfo',
    'enableStun',
    'enableSrtp',
 ],
  properties: {
    interface: {
      title: 'Core Interface',
      type: 'string',
      enum: ['LAN A', 'LAN B', 'Aux A', 'Aux B'],
    },
    sipLogging: {
      title: 'Enable Logging',
      type: 'boolean',
    },
    sipPort: {
      title: 'SIP Port',
      type: 'number',
      range: [1, 65535],
      default: 5060,
    },
    enableDtmfInfo: {
      title: 'Enable DTMF Info',
      type: 'boolean',
      default: false,
    },
    dtmfPayloadType: {
      title: 'RFC2833 DTMF Type',
      type: 'number',
      range: [96, 127],
      default: 101,
    },
    enableStun: {
      title: 'Enable Stun',
      type: 'boolean',
    },
    stunServer: {
      oneOf: [
        {
          title: 'Stun Server',
          type: 'string',
          minLength: 1,
          maxLength: 255,
        },
        {type: 'null'},
      ],
    },
    enableSrtp: {
      title: 'Enable SRTP',
      type: 'boolean',
    },
    softphones: {
      type: 'array',
      minItems: 1,
      items: SoftphoneHttpSchema.structure,
    },
    codecs: {
      type: 'array',
      minItems: 1,
      items: CodecHttpSchema.structure,
    },
  },
  switch: [
    {
      if: {
        properties: {
          enableDtmfInfo: {enum: [true]},
        },
      },
      then: {
        prohibited: ['dtmfPayloadType'],
      },
      continue: true,
    },
    {
      if: {
        properties: {
          enableDtmfInfo: {enum: [false]},
        },
      },
      then: {
        required: ['dtmfPayloadType'],
      },
      continue: true,
    },
    {
      if: {
        properties: {
          enableStun: {enum: [true]},
        },
      },
      then: {
        required: ['stunServer'],
      },
      continue: true,
    },
  ],
};
````
