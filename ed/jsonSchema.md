JSON-Schema
-

https://spacetelescope.github.io/understanding-json-schema/

````
{ "$ref": "#/definitions/address" }
````

````
const httpSchema = {
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
        { type: 'null' },
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
          enableDtmfInfo: { enum: [true] },
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
          enableDtmfInfo: { enum: [false] },
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
          enableStun: { enum: [true] },
        },
      },
      then: {
````
        required: ['stunServer'],
      },
      continue: true,
    },
  ],
};
