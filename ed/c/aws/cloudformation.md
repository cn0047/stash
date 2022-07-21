CloudFormation
-

Allows to use simple text file to model and provision all the resources needed for application.

````
Ref PubSubTopic # reference
Fn::Select
Fn::Join
Fn::Base64
````

````yaml
AWSTemplateFormatVersion: '2010-09-09'
Mappings
Parameters: # types: String, Number, List ([1,2,3]), CommaDelinitedList ("1,2,3").
  MaxSize:
    Type: Number
    Default: '1'
    Description: Maximum size.
    AllowedValues: ['1', '2', '3'],
Conditions:
Resources:
````
