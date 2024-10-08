package examples

import (
	"testing"

	"github.com/schleppy/grule-rule-engine/ast"
	"github.com/schleppy/grule-rule-engine/builder"
	"github.com/schleppy/grule-rule-engine/engine"
	"github.com/schleppy/grule-rule-engine/logger"
	"github.com/schleppy/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	Index         int
	Strings       []string
	Concatenation string
}

func (f *TestData) GetStrings() []string {
	return f.Strings
}

const rule = ` rule test {
when 
	Fact.Index < Fact.Strings.Len()
then
	Fact.Concatenation = Fact.Concatenation + Fact.GetStrings()[Fact.Index];
	Fact.Index = Fact.Index + 1;
}`

func TestSliceFunctionTest(t *testing.T) {
	fact := &TestData{
		Index:   0,
		Strings: []string{"1", "2", "3"},
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("Fact", fact)
	assert.NoError(t, err)
	logs := logger.NewDefaultLogger()
	knowledgeLibrary := ast.NewKnowledgeLibrary(logs)
	ruleBuilder := builder.NewRuleBuilder(logs, knowledgeLibrary)
	err = ruleBuilder.BuildRuleFromResource("test", "0.0.1", pkg.NewBytesResource([]byte(rule)))
	assert.NoError(t, err)
	knowledgeBase, err := knowledgeLibrary.NewKnowledgeBaseInstance("test", "0.0.1")
	assert.NoError(t, err)
	engine := engine.NewGruleEngine(logs)

	err = engine.Execute(dataContext, knowledgeBase)
	assert.NoError(t, err)
	assert.Equal(t, "123", fact.Concatenation)
}
