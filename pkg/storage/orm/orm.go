package orm

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type dbObject struct {
	self interface{}
	ids  []interface{}
}

type childToParent struct {
	self interface{}
	parents []interface{}
}

var propertyDependency = childToParent{
	self:   &property.Property{},
	parents: []interface{}{&service.Service{}},
}


var propertyDependency = childToParent{
	self:   &property.Property{},
	parents: []interface{}{&check.Check{}},
}


func verifyStruct(db gorm.DB, parent dbObject, self dbObject){


	switch v := parent.self.(type) {
	case *team.Team:
	case *service_group.ServiceGroup:
	default:
		fmt.Println("unknown")
	}

}

type AddFromToUpdate struct {}

func (AddFromToUpdate) ModifyStatement(stmt *gorm.Statement) {
	stmt.Statement.BuildClauses = []string{"UPDATE", "SET", "FROM", "WHERE"}
}

func (AddFromToUpdate) Name() string {
	return ""
}

func (AddFromToUpdate) Build(clause.Builder) {
}

func (AddFromToUpdate) MergeClause(*clause.Clause) {
}

type WhereClause struct {
	Query interface{}
	Arg   []interface{}
}

func GenericGet(ctx context.Context, db *gorm.DB, whereChain []WhereClause, ret interface{}, isSingle bool, fromAdditionalTables []string ) error{
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(ret)
	tableName := strings.ToLower(stmt.Schema.Table)
	db = db.WithContext(ctx)
	uniqueAdditionalTables := SliceUniqMap(fromAdditionalTables)
	if !(len(uniqueAdditionalTables) == 0 || (len(uniqueAdditionalTables) == 1 && strings.ToLower(uniqueAdditionalTables[0]) == tableName)){
		db = db.Select(fmt.Sprintf("%s.*", tableName))
	}
	var fromTables []clause.Table
	for i := range uniqueAdditionalTables {
		fromTables = append(fromTables, clause.Table{Name: strings.ToLower(uniqueAdditionalTables[i])})
	}
	containsCurrentTable := false
	for i := range fromTables {
		if fromTables[i].Name == tableName{
			containsCurrentTable=true
		}
	}
	if !containsCurrentTable{
		fromTables = append(fromTables, clause.Table{Name: tableName})
	}
	db = db.Clauses(clause.From{
		Tables: fromTables,
	})
	for i := range whereChain{
		db = db.Where(whereChain[i].Query, whereChain[i].Arg...)
	}
	if isSingle{
		db = db.First(ret)
	} else{
		db = db.Find(ret)
	}
	return nil
}

func SliceUniqMap(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
