// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/apitoken"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
)

// ApiTokenQuery is the builder for querying ApiToken entities.
type ApiTokenQuery struct {
	config
	ctx        *QueryContext
	order      []apitoken.OrderOption
	inters     []Interceptor
	predicates []predicate.ApiToken
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApiTokenQuery builder.
func (atq *ApiTokenQuery) Where(ps ...predicate.ApiToken) *ApiTokenQuery {
	atq.predicates = append(atq.predicates, ps...)
	return atq
}

// Limit the number of records to be returned by this query.
func (atq *ApiTokenQuery) Limit(limit int) *ApiTokenQuery {
	atq.ctx.Limit = &limit
	return atq
}

// Offset to start from.
func (atq *ApiTokenQuery) Offset(offset int) *ApiTokenQuery {
	atq.ctx.Offset = &offset
	return atq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (atq *ApiTokenQuery) Unique(unique bool) *ApiTokenQuery {
	atq.ctx.Unique = &unique
	return atq
}

// Order specifies how the records should be ordered.
func (atq *ApiTokenQuery) Order(o ...apitoken.OrderOption) *ApiTokenQuery {
	atq.order = append(atq.order, o...)
	return atq
}

// First returns the first ApiToken entity from the query.
// Returns a *NotFoundError when no ApiToken was found.
func (atq *ApiTokenQuery) First(ctx context.Context) (*ApiToken, error) {
	nodes, err := atq.Limit(1).All(setContextOp(ctx, atq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{apitoken.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (atq *ApiTokenQuery) FirstX(ctx context.Context) *ApiToken {
	node, err := atq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ApiToken ID from the query.
// Returns a *NotFoundError when no ApiToken ID was found.
func (atq *ApiTokenQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = atq.Limit(1).IDs(setContextOp(ctx, atq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{apitoken.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (atq *ApiTokenQuery) FirstIDX(ctx context.Context) string {
	id, err := atq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ApiToken entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ApiToken entity is found.
// Returns a *NotFoundError when no ApiToken entities are found.
func (atq *ApiTokenQuery) Only(ctx context.Context) (*ApiToken, error) {
	nodes, err := atq.Limit(2).All(setContextOp(ctx, atq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{apitoken.Label}
	default:
		return nil, &NotSingularError{apitoken.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (atq *ApiTokenQuery) OnlyX(ctx context.Context) *ApiToken {
	node, err := atq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ApiToken ID in the query.
// Returns a *NotSingularError when more than one ApiToken ID is found.
// Returns a *NotFoundError when no entities are found.
func (atq *ApiTokenQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = atq.Limit(2).IDs(setContextOp(ctx, atq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{apitoken.Label}
	default:
		err = &NotSingularError{apitoken.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (atq *ApiTokenQuery) OnlyIDX(ctx context.Context) string {
	id, err := atq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ApiTokens.
func (atq *ApiTokenQuery) All(ctx context.Context) ([]*ApiToken, error) {
	ctx = setContextOp(ctx, atq.ctx, "All")
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ApiToken, *ApiTokenQuery]()
	return withInterceptors[[]*ApiToken](ctx, atq, qr, atq.inters)
}

// AllX is like All, but panics if an error occurs.
func (atq *ApiTokenQuery) AllX(ctx context.Context) []*ApiToken {
	nodes, err := atq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ApiToken IDs.
func (atq *ApiTokenQuery) IDs(ctx context.Context) (ids []string, err error) {
	if atq.ctx.Unique == nil && atq.path != nil {
		atq.Unique(true)
	}
	ctx = setContextOp(ctx, atq.ctx, "IDs")
	if err = atq.Select(apitoken.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (atq *ApiTokenQuery) IDsX(ctx context.Context) []string {
	ids, err := atq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (atq *ApiTokenQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, atq.ctx, "Count")
	if err := atq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, atq, querierCount[*ApiTokenQuery](), atq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (atq *ApiTokenQuery) CountX(ctx context.Context) int {
	count, err := atq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (atq *ApiTokenQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, atq.ctx, "Exist")
	switch _, err := atq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (atq *ApiTokenQuery) ExistX(ctx context.Context) bool {
	exist, err := atq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApiTokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (atq *ApiTokenQuery) Clone() *ApiTokenQuery {
	if atq == nil {
		return nil
	}
	return &ApiTokenQuery{
		config:     atq.config,
		ctx:        atq.ctx.Clone(),
		order:      append([]apitoken.OrderOption{}, atq.order...),
		inters:     append([]Interceptor{}, atq.inters...),
		predicates: append([]predicate.ApiToken{}, atq.predicates...),
		// clone intermediate query.
		sql:  atq.sql.Clone(),
		path: atq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ApiToken.Query().
//		GroupBy(apitoken.FieldCreateTime).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (atq *ApiTokenQuery) GroupBy(field string, fields ...string) *ApiTokenGroupBy {
	atq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApiTokenGroupBy{build: atq}
	grbuild.flds = &atq.ctx.Fields
	grbuild.label = apitoken.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ApiToken.Query().
//		Select(apitoken.FieldCreateTime).
//		Scan(ctx, &v)
func (atq *ApiTokenQuery) Select(fields ...string) *ApiTokenSelect {
	atq.ctx.Fields = append(atq.ctx.Fields, fields...)
	sbuild := &ApiTokenSelect{ApiTokenQuery: atq}
	sbuild.label = apitoken.Label
	sbuild.flds, sbuild.scan = &atq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApiTokenSelect configured with the given aggregations.
func (atq *ApiTokenQuery) Aggregate(fns ...AggregateFunc) *ApiTokenSelect {
	return atq.Select().Aggregate(fns...)
}

func (atq *ApiTokenQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range atq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, atq); err != nil {
				return err
			}
		}
	}
	for _, f := range atq.ctx.Fields {
		if !apitoken.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if atq.path != nil {
		prev, err := atq.path(ctx)
		if err != nil {
			return err
		}
		atq.sql = prev
	}
	return nil
}

func (atq *ApiTokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ApiToken, error) {
	var (
		nodes = []*ApiToken{}
		_spec = atq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ApiToken).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ApiToken{config: atq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, atq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (atq *ApiTokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := atq.querySpec()
	_spec.Node.Columns = atq.ctx.Fields
	if len(atq.ctx.Fields) > 0 {
		_spec.Unique = atq.ctx.Unique != nil && *atq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, atq.driver, _spec)
}

func (atq *ApiTokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(apitoken.Table, apitoken.Columns, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeString))
	_spec.From = atq.sql
	if unique := atq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if atq.path != nil {
		_spec.Unique = true
	}
	if fields := atq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apitoken.FieldID)
		for i := range fields {
			if fields[i] != apitoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := atq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := atq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := atq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := atq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (atq *ApiTokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(atq.driver.Dialect())
	t1 := builder.Table(apitoken.Table)
	columns := atq.ctx.Fields
	if len(columns) == 0 {
		columns = apitoken.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if atq.sql != nil {
		selector = atq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if atq.ctx.Unique != nil && *atq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range atq.predicates {
		p(selector)
	}
	for _, p := range atq.order {
		p(selector)
	}
	if offset := atq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := atq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApiTokenGroupBy is the group-by builder for ApiToken entities.
type ApiTokenGroupBy struct {
	selector
	build *ApiTokenQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (atgb *ApiTokenGroupBy) Aggregate(fns ...AggregateFunc) *ApiTokenGroupBy {
	atgb.fns = append(atgb.fns, fns...)
	return atgb
}

// Scan applies the selector query and scans the result into the given value.
func (atgb *ApiTokenGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, atgb.build.ctx, "GroupBy")
	if err := atgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiTokenQuery, *ApiTokenGroupBy](ctx, atgb.build, atgb, atgb.build.inters, v)
}

func (atgb *ApiTokenGroupBy) sqlScan(ctx context.Context, root *ApiTokenQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(atgb.fns))
	for _, fn := range atgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*atgb.flds)+len(atgb.fns))
		for _, f := range *atgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*atgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := atgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApiTokenSelect is the builder for selecting fields of ApiToken entities.
type ApiTokenSelect struct {
	*ApiTokenQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ats *ApiTokenSelect) Aggregate(fns ...AggregateFunc) *ApiTokenSelect {
	ats.fns = append(ats.fns, fns...)
	return ats
}

// Scan applies the selector query and scans the result into the given value.
func (ats *ApiTokenSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ats.ctx, "Select")
	if err := ats.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApiTokenQuery, *ApiTokenSelect](ctx, ats.ApiTokenQuery, ats, ats.inters, v)
}

func (ats *ApiTokenSelect) sqlScan(ctx context.Context, root *ApiTokenQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ats.fns))
	for _, fn := range ats.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ats.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ats.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
