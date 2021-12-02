package map2xml

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBasicXML(t *testing.T) {
	var value = gin.H{
		"test": gin.H{
			"num":    2,
			"string": "hi",
			"nil":    nil,
			"slices": []gin.H{
				{"xml_child_name": "slice", "one": 1},
				{"two": 2},
				{"three": 3},
			},
			"multi": gin.H{
				"level": "value",
			},
		},
	}
	var expected = `<test>
  <multi>
    <level>value</level>
  </multi>
  <nil></nil>
  <num>2</num>
  <slices>
    <slice>
      <one>1</one>
    </slice>
    <slice>
      <two>2</two>
    </slice>
    <slice>
      <three>3</three>
    </slice>
  </slices>
  <string>hi</string>
</test>`

	xml := New(value)
	xml.WithIndent("", "  ")
	ans, err := xml.MarshalToString()
	if err != nil {
		t.Error(err)
	}
	if ans != expected {
		t.Logf("expected:\n---\n'%s'\n---\nvalue:\n+++\n%s\n+++\n", expected, ans)
		t.Fail()
	}
}
