Feature: build a XML based on a given template
  Scenario: Happy scenario 01
    Given I want to build a XML based on following JSON:
    """
    {
      "event":{
        "header":{
          "eventname":"test",
          "eventversion": 1
        },
        "content":{
          "message":"Simple message"
        }
      }
    }
    """
    Then the generated XML must be a valid XML
    And the XML tag "/*[local-name()='event']" of the generated XML must exist
    And the XML tag "/*[local-name()='event']/*[local-name()='header']/*[local-name()='eventname']/text()" of the generated XML should be equals to "test"
    And the XML tag "/*[local-name()='event']/*[local-name()='content']/*[local-name()='message']/text()" of the generated XML should be equals to "Simple message"

