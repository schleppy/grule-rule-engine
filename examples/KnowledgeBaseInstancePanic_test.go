//  Copyright schleppy/grule-rule-engine Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package examples

import (
	"testing"

	"github.com/schleppy/grule-rule-engine/ast"
	"github.com/schleppy/grule-rule-engine/builder"
	"github.com/schleppy/grule-rule-engine/logger"
	"github.com/schleppy/grule-rule-engine/pkg"
	"github.com/stretchr/testify/assert"
)

func TestNoPanicForNoDescription(t *testing.T) {
	GRL := `rule TestNoDesc { when true then Ok(); }`
	logs := logger.NewDefaultLogger()
	lib := ast.NewKnowledgeLibrary(logs)
	ruleBuilder := builder.NewRuleBuilder(logs, lib)
	err := ruleBuilder.BuildRuleFromResource("CallingLog", "0.1.1", pkg.NewBytesResource([]byte(GRL)))
	assert.NoError(t, err)
	_, err = lib.NewKnowledgeBaseInstance("CallingLog", "0.1.1")
	assert.NoError(t, err)
}
