package tests

import (
	"context"
	"testing"

	"github.com/project-flogo/rules/common/model"
	"github.com/project-flogo/rules/ruleapi"
)

//1 arithmetic operation
func Test_5_Expr(t *testing.T) {

	actionCount := map[string]int{"count": 0}
	rs, _ := createRuleSession()
	r1 := ruleapi.NewRule("r1")
	r1.AddExprCondition("c1", "(($.t1.p1 + $.t2.p1) == 5) && (($.t1.p2 > $.t2.p2) && ($.t1.p3 == $.t2.p3))", nil)
	r1.SetAction(a5)
	r1.SetContext(actionCount)

	rs.AddRule(r1)

	rs.Start(nil)

	var ctx context.Context

	t1, _ := model.NewTupleWithKeyValues("t1", "t1")
	t1.SetInt(context.TODO(), "p1", 1)
	t1.SetDouble(context.TODO(), "p2", 1.3)
	t1.SetString(context.TODO(), "p3", "t3")

	ctx = context.WithValue(context.TODO(), TestKey{}, t)
	rs.Assert(ctx, t1)

	t2, _ := model.NewTupleWithKeyValues("t2", "t2")
	t2.SetInt(context.TODO(), "p1", 4)
	t2.SetDouble(context.TODO(), "p2", 1.1)
	t2.SetString(context.TODO(), "p3", "t3")

	ctx = context.WithValue(context.TODO(), TestKey{}, t)
	rs.Assert(ctx, t2)
	rs.Unregister()
	count := actionCount["count"]
	if count != 1 {
		t.Errorf("expected [%d], got [%d]\n", 1, count)
	}
}

func a5(ctx context.Context, rs model.RuleSession, ruleName string, tuples map[model.TupleType]model.Tuple, ruleCtx model.RuleContext) {
	t := ctx.Value(TestKey{}).(*testing.T)
	t.Logf("Test_5_Expr executed!")
	actionCount := ruleCtx.(map[string]int)
	count := actionCount["count"]
	actionCount["count"] = count + 1
}
