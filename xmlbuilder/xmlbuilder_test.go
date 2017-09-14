package xmlbuilder

import (
	"bytes"
	"fmt"
	"github.com/ChrisTrenkamp/goxpath"
	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	format := "pretty"
	for _, arg := range os.Args[1:] {
		if arg == "-test.v=true" { // go test transforms -v option
			format = "pretty"
			break
		}
	}

	status := godog.RunWithOptions("calculator", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format: format,
		Paths:  []string{"features"},

		StopOnFailure: true,
		Randomize:     time.Now().UTC().UnixNano(), // randomize scenario execution order
	})

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

var xmlContent string

func iWantToBuildAXMLBasedOnFollowingJSON(jsonContent *gherkin.DocString) error {
	var err error

	xmlContent, err = ToXmlFromJson(jsonContent.Content)

	if err != nil {
		return fmt.Errorf("failure on convert the JSON to XML: %v", err)
	}

	fmt.Println(xmlContent)
	return nil
}

func theGeneratedXMLMustBeAValidXML() error {

	return theXMLTagOfTheGeneratedXMLMustExist("/*")
}

func theXMLTagOfTheGeneratedXMLMustExist(xpath string) error {
	res := goxpath.MustParse(xpath).MustExec(xmltree.MustParseXML(bytes.NewBufferString(xmlContent)),
		func(o *goxpath.Opts) { o.NS = make(map[string]string, 0) }).(tree.NodeSet)

	if len(res) == 0 {
		return fmt.Errorf("result length not valid in XPath expression '%v': %v %v", xpath, len(res), ", expecting at least 1")
	}

	return nil
}

func theXMLTagOfTheGeneratedXMLShouldBeEqualsTo(xpath, expectedValue string) error {
	res := goxpath.MustParse(xpath).MustExec(xmltree.MustParseXML(bytes.NewBufferString(xmlContent)),
		func(o *goxpath.Opts) { o.NS = make(map[string]string, 0) }).(tree.NodeSet)

	if len(res) == 0 {
		return fmt.Errorf("result length not valid in XPath expression '%v': %v %v", xpath, len(res), ", expecting at least 1")
	}

	value, err := goxpath.MarshalStr(res[0].(tree.Node))
	if err == nil {
		if value != expectedValue {
			return fmt.Errorf("invalid value: expected <%v> but actual is <%v>", expectedValue, value)
		}
	} else {
		return fmt.Errorf("no exists the xpath: %v", xpath)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I want to build a XML based on following JSON:$`, iWantToBuildAXMLBasedOnFollowingJSON)
	s.Step(`^the generated XML must be a valid XML$`, theGeneratedXMLMustBeAValidXML)
	s.Step(`^the XML tag "([^"]*)" of the generated XML must exist$`, theXMLTagOfTheGeneratedXMLMustExist)
	s.Step(`^the XML tag "([^"]*)" of the generated XML should be equals to "([^"]*)"$`, theXMLTagOfTheGeneratedXMLShouldBeEqualsTo)
}
