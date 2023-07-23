// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/round"
)

// CheckQuery is the builder for querying Check entities.
type CheckQuery struct {
	config
	ctx             *QueryContext
	order           []check.OrderOption
	inters          []Interceptor
	predicates      []predicate.Check
	withRound       *RoundQuery
	withHostservice *HostServiceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CheckQuery builder.
func (cq *CheckQuery) Where(ps ...predicate.Check) *CheckQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CheckQuery) Limit(limit int) *CheckQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CheckQuery) Offset(offset int) *CheckQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CheckQuery) Unique(unique bool) *CheckQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CheckQuery) Order(o ...check.OrderOption) *CheckQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryRound chains the current query on the "round" edge.
func (cq *CheckQuery) QueryRound() *RoundQuery {
	query := (&RoundClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(check.Table, check.FieldID, selector),
			sqlgraph.To(round.Table, round.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, check.RoundTable, check.RoundColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHostservice chains the current query on the "hostservice" edge.
func (cq *CheckQuery) QueryHostservice() *HostServiceQuery {
	query := (&HostServiceClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(check.Table, check.FieldID, selector),
			sqlgraph.To(hostservice.Table, hostservice.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, check.HostserviceTable, check.HostserviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Check entity from the query.
// Returns a *NotFoundError when no Check was found.
func (cq *CheckQuery) First(ctx context.Context) (*Check, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{check.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CheckQuery) FirstX(ctx context.Context) *Check {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Check ID from the query.
// Returns a *NotFoundError when no Check ID was found.
func (cq *CheckQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{check.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CheckQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Check entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Check entity is found.
// Returns a *NotFoundError when no Check entities are found.
func (cq *CheckQuery) Only(ctx context.Context) (*Check, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{check.Label}
	default:
		return nil, &NotSingularError{check.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CheckQuery) OnlyX(ctx context.Context) *Check {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Check ID in the query.
// Returns a *NotSingularError when more than one Check ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CheckQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{check.Label}
	default:
		err = &NotSingularError{check.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CheckQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Checks.
func (cq *CheckQuery) All(ctx context.Context) ([]*Check, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Check, *CheckQuery]()
	return withInterceptors[[]*Check](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CheckQuery) AllX(ctx context.Context) []*Check {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Check IDs.
func (cq *CheckQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(check.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CheckQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CheckQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CheckQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CheckQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CheckQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CheckQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CheckQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CheckQuery) Clone() *CheckQuery {
	if cq == nil {
		return nil
	}
	return &CheckQuery{
		config:          cq.config,
		ctx:             cq.ctx.Clone(),
		order:           append([]check.OrderOption{}, cq.order...),
		inters:          append([]Interceptor{}, cq.inters...),
		predicates:      append([]predicate.Check{}, cq.predicates...),
		withRound:       cq.withRound.Clone(),
		withHostservice: cq.withHostservice.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithRound tells the query-builder to eager-load the nodes that are connected to
// the "round" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CheckQuery) WithRound(opts ...func(*RoundQuery)) *CheckQuery {
	query := (&RoundClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withRound = query
	return cq
}

// WithHostservice tells the query-builder to eager-load the nodes that are connected to
// the "hostservice" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CheckQuery) WithHostservice(opts ...func(*HostServiceQuery)) *CheckQuery {
	query := (&HostServiceClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withHostservice = query
	return cq
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
//	client.Check.Query().
//		GroupBy(check.FieldCreateTime).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (cq *CheckQuery) GroupBy(field string, fields ...string) *CheckGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CheckGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = check.Label
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
//	client.Check.Query().
//		Select(check.FieldCreateTime).
//		Scan(ctx, &v)
func (cq *CheckQuery) Select(fields ...string) *CheckSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CheckSelect{CheckQuery: cq}
	sbuild.label = check.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CheckSelect configured with the given aggregations.
func (cq *CheckQuery) Aggregate(fns ...AggregateFunc) *CheckSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CheckQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !check.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CheckQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Check, error) {
	var (
		nodes       = []*Check{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withRound != nil,
			cq.withHostservice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Check).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Check{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withRound; query != nil {
		if err := cq.loadRound(ctx, query, nodes, nil,
			func(n *Check, e *Round) { n.Edges.Round = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withHostservice; query != nil {
		if err := cq.loadHostservice(ctx, query, nodes, nil,
			func(n *Check, e *HostService) { n.Edges.Hostservice = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CheckQuery) loadRound(ctx context.Context, query *RoundQuery, nodes []*Check, init func(*Check), assign func(*Check, *Round)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Check)
	for i := range nodes {
		fk := nodes[i].RoundID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(round.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "round_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CheckQuery) loadHostservice(ctx context.Context, query *HostServiceQuery, nodes []*Check, init func(*Check), assign func(*Check, *HostService)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Check)
	for i := range nodes {
		fk := nodes[i].HostServiceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(hostservice.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "host_service_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *CheckQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CheckQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(check.Table, check.Columns, sqlgraph.NewFieldSpec(check.FieldID, field.TypeString))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, check.FieldID)
		for i := range fields {
			if fields[i] != check.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withRound != nil {
			_spec.Node.AddColumnOnce(check.FieldRoundID)
		}
		if cq.withHostservice != nil {
			_spec.Node.AddColumnOnce(check.FieldHostServiceID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CheckQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(check.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = check.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CheckGroupBy is the group-by builder for Check entities.
type CheckGroupBy struct {
	selector
	build *CheckQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CheckGroupBy) Aggregate(fns ...AggregateFunc) *CheckGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CheckGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CheckQuery, *CheckGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CheckGroupBy) sqlScan(ctx context.Context, root *CheckQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CheckSelect is the builder for selecting fields of Check entities.
type CheckSelect struct {
	*CheckQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CheckSelect) Aggregate(fns ...AggregateFunc) *CheckSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CheckSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CheckQuery, *CheckSelect](ctx, cs.CheckQuery, cs, cs.inters, v)
}

func (cs *CheckSelect) sqlScan(ctx context.Context, root *CheckQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
