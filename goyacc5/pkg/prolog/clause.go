package prolog

import (
	"fmt"
	"strings"
)

// either a Fact or a Rule
type clause interface {
	String() string
	Head() term
	Body() []term // Facts have no body (empty)
	Evaluate(context map[string]term) clause
	Dump()
}

// factClause: represents a fact, which is just a single compound term.
type factClause struct {
	Fact term // The fact is a single compound term
}

func (fc *factClause) String() string {
	return fmt.Sprintf("%s.", fc.Fact.String())
}

func (fc *factClause) Head() term {
	return fc.Fact
}

func (fc *factClause) Body() []term {
	return nil // Facts have no body
}

func (fc *factClause) Evaluate(context map[string]term) clause {
	// A fact evaluates by evaluating its single term
	return &factClause{Fact: fc.Fact.Evaluate(context)}
}

func (fc *factClause) Dump() {
	fmt.Printf("fact) %s/%d\n", fc.Head().GetFunctor(), len(fc.Fact.GetArgs()))
	fmt.Println(" " + fc.String())
}

// ruleClause: represents a rule, which has a head and a body (conditions).
type ruleClause struct {
	HeadTerm  term   // The head (conclusion) of the rule
	BodyTerms []term // The body (conditions) of the rule
}

func (rc *ruleClause) String() string {
	bodyStr := []string{}
	for _, term := range rc.BodyTerms {
		bodyStr = append(bodyStr, term.String())
	}
	return fmt.Sprintf("%s :- %s.", rc.HeadTerm.String(), strings.Join(bodyStr, ", "))
}

func (rc *ruleClause) Head() term {
	return rc.HeadTerm
}

func (rc *ruleClause) Body() []term {
	return rc.BodyTerms
}

func (rc *ruleClause) Evaluate(context map[string]term) clause {
	// Evaluate the head and each term in the body
	evaluatedHead := rc.HeadTerm.Evaluate(context)
	evaluatedBody := []term{}
	for _, bodyTerm := range rc.BodyTerms {
		evaluatedBody = append(evaluatedBody, bodyTerm.Evaluate(context))
	}
	return &ruleClause{
		HeadTerm:  evaluatedHead,
		BodyTerms: evaluatedBody,
	}
}

func (rc *ruleClause) Dump() {
	fmt.Printf("rule) %s/%d\n", rc.Head().GetFunctor(), len(rc.BodyTerms))
	fmt.Println(" " + rc.String())
}
