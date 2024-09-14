package examples

import (
	"testing"

	"github.com/schelppy/grule-rule-engine/ast"
	"github.com/schelppy/grule-rule-engine/builder"
	"github.com/schelppy/grule-rule-engine/engine"
	"github.com/schelppy/grule-rule-engine/logger"
	"github.com/schelppy/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
)

const (
	SliceOORRule = `
		rule SliceOORRule {
			when
				PriceSlice.Prices[4] > 10 // will cause panic
			then
				Log("Price number 4 is greater than 10");
				Retract("SliceOORRule");
		}`
)

type AUserSliceIssue struct {
	Prices []int
}

func TestMethodCall_SliceOOR(t *testing.T) {
	ps := &AUserSliceIssue{
		Prices: []int{1, 2, 3},
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("PriceSlice", ps)
	assert.NoError(t, err)

	// Prepare knowledgebase library and load it with our rule.
	logs := logger.NewDefaultLogger()
	lib := ast.NewKnowledgeLibrary(logs)
	rb := builder.NewRuleBuilder(logs, lib)
	err = rb.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(SliceOORRule)))
	assert.NoError(t, err)

	eng1 := engine.NewGruleEngine(logs)
	eng1.MaxCycle = 5
	kb, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	assert.NoError(t, err)
	err = eng1.Execute(dataContext, kb)
	assert.NoError(t, err)

	eng1 = engine.NewGruleEngine(logs)
	eng1.MaxCycle = 5
	eng1.ReturnErrOnFailedRuleEvaluation = true
	kb, err = lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	assert.NoError(t, err)
	err = eng1.Execute(dataContext, kb)
	assert.Error(t, err)
}
