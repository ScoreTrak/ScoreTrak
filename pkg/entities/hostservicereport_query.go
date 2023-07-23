// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/hostservicereport"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/predicate"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/teamreport"
)

// HostServiceReportQuery is the builder for querying HostServiceReport entities.
type HostServiceReportQuery struct {
	config
	ctx             *QueryContext
	order           []hostservicereport.OrderOption
	inters          []Interceptor
	predicates      []predicate.HostServiceReport
	withHostservice *HostServiceQuery
	withService     *ServiceQuery
	withTeam        *TeamQuery
	withTeamreport  *TeamReportQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HostServiceReportQuery builder.
func (hsrq *HostServiceReportQuery) Where(ps ...predicate.HostServiceReport) *HostServiceReportQuery {
	hsrq.predicates = append(hsrq.predicates, ps...)
	return hsrq
}

// Limit the number of records to be returned by this query.
func (hsrq *HostServiceReportQuery) Limit(limit int) *HostServiceReportQuery {
	hsrq.ctx.Limit = &limit
	return hsrq
}

// Offset to start from.
func (hsrq *HostServiceReportQuery) Offset(offset int) *HostServiceReportQuery {
	hsrq.ctx.Offset = &offset
	return hsrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hsrq *HostServiceReportQuery) Unique(unique bool) *HostServiceReportQuery {
	hsrq.ctx.Unique = &unique
	return hsrq
}

// Order specifies how the records should be ordered.
func (hsrq *HostServiceReportQuery) Order(o ...hostservicereport.OrderOption) *HostServiceReportQuery {
	hsrq.order = append(hsrq.order, o...)
	return hsrq
}

// QueryHostservice chains the current query on the "hostservice" edge.
func (hsrq *HostServiceReportQuery) QueryHostservice() *HostServiceQuery {
	query := (&HostServiceClient{config: hsrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hsrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hsrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hostservicereport.Table, hostservicereport.FieldID, selector),
			sqlgraph.To(hostservice.Table, hostservice.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, hostservicereport.HostserviceTable, hostservicereport.HostserviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(hsrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryService chains the current query on the "service" edge.
func (hsrq *HostServiceReportQuery) QueryService() *ServiceQuery {
	query := (&ServiceClient{config: hsrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hsrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hsrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hostservicereport.Table, hostservicereport.FieldID, selector),
			sqlgraph.To(service.Table, service.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hostservicereport.ServiceTable, hostservicereport.ServiceColumn),
		)
		fromU = sqlgraph.SetNeighbors(hsrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeam chains the current query on the "team" edge.
func (hsrq *HostServiceReportQuery) QueryTeam() *TeamQuery {
	query := (&TeamClient{config: hsrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hsrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hsrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hostservicereport.Table, hostservicereport.FieldID, selector),
			sqlgraph.To(team.Table, team.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hostservicereport.TeamTable, hostservicereport.TeamColumn),
		)
		fromU = sqlgraph.SetNeighbors(hsrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeamreport chains the current query on the "teamreport" edge.
func (hsrq *HostServiceReportQuery) QueryTeamreport() *TeamReportQuery {
	query := (&TeamReportClient{config: hsrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := hsrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := hsrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(hostservicereport.Table, hostservicereport.FieldID, selector),
			sqlgraph.To(teamreport.Table, teamreport.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, hostservicereport.TeamreportTable, hostservicereport.TeamreportColumn),
		)
		fromU = sqlgraph.SetNeighbors(hsrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first HostServiceReport entity from the query.
// Returns a *NotFoundError when no HostServiceReport was found.
func (hsrq *HostServiceReportQuery) First(ctx context.Context) (*HostServiceReport, error) {
	nodes, err := hsrq.Limit(1).All(setContextOp(ctx, hsrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{hostservicereport.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) FirstX(ctx context.Context) *HostServiceReport {
	node, err := hsrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first HostServiceReport ID from the query.
// Returns a *NotFoundError when no HostServiceReport ID was found.
func (hsrq *HostServiceReportQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = hsrq.Limit(1).IDs(setContextOp(ctx, hsrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{hostservicereport.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) FirstIDX(ctx context.Context) string {
	id, err := hsrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single HostServiceReport entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one HostServiceReport entity is found.
// Returns a *NotFoundError when no HostServiceReport entities are found.
func (hsrq *HostServiceReportQuery) Only(ctx context.Context) (*HostServiceReport, error) {
	nodes, err := hsrq.Limit(2).All(setContextOp(ctx, hsrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{hostservicereport.Label}
	default:
		return nil, &NotSingularError{hostservicereport.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) OnlyX(ctx context.Context) *HostServiceReport {
	node, err := hsrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only HostServiceReport ID in the query.
// Returns a *NotSingularError when more than one HostServiceReport ID is found.
// Returns a *NotFoundError when no entities are found.
func (hsrq *HostServiceReportQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = hsrq.Limit(2).IDs(setContextOp(ctx, hsrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{hostservicereport.Label}
	default:
		err = &NotSingularError{hostservicereport.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) OnlyIDX(ctx context.Context) string {
	id, err := hsrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of HostServiceReports.
func (hsrq *HostServiceReportQuery) All(ctx context.Context) ([]*HostServiceReport, error) {
	ctx = setContextOp(ctx, hsrq.ctx, "All")
	if err := hsrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*HostServiceReport, *HostServiceReportQuery]()
	return withInterceptors[[]*HostServiceReport](ctx, hsrq, qr, hsrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) AllX(ctx context.Context) []*HostServiceReport {
	nodes, err := hsrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of HostServiceReport IDs.
func (hsrq *HostServiceReportQuery) IDs(ctx context.Context) (ids []string, err error) {
	if hsrq.ctx.Unique == nil && hsrq.path != nil {
		hsrq.Unique(true)
	}
	ctx = setContextOp(ctx, hsrq.ctx, "IDs")
	if err = hsrq.Select(hostservicereport.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) IDsX(ctx context.Context) []string {
	ids, err := hsrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hsrq *HostServiceReportQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, hsrq.ctx, "Count")
	if err := hsrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, hsrq, querierCount[*HostServiceReportQuery](), hsrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) CountX(ctx context.Context) int {
	count, err := hsrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hsrq *HostServiceReportQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, hsrq.ctx, "Exist")
	switch _, err := hsrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (hsrq *HostServiceReportQuery) ExistX(ctx context.Context) bool {
	exist, err := hsrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HostServiceReportQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hsrq *HostServiceReportQuery) Clone() *HostServiceReportQuery {
	if hsrq == nil {
		return nil
	}
	return &HostServiceReportQuery{
		config:          hsrq.config,
		ctx:             hsrq.ctx.Clone(),
		order:           append([]hostservicereport.OrderOption{}, hsrq.order...),
		inters:          append([]Interceptor{}, hsrq.inters...),
		predicates:      append([]predicate.HostServiceReport{}, hsrq.predicates...),
		withHostservice: hsrq.withHostservice.Clone(),
		withService:     hsrq.withService.Clone(),
		withTeam:        hsrq.withTeam.Clone(),
		withTeamreport:  hsrq.withTeamreport.Clone(),
		// clone intermediate query.
		sql:  hsrq.sql.Clone(),
		path: hsrq.path,
	}
}

// WithHostservice tells the query-builder to eager-load the nodes that are connected to
// the "hostservice" edge. The optional arguments are used to configure the query builder of the edge.
func (hsrq *HostServiceReportQuery) WithHostservice(opts ...func(*HostServiceQuery)) *HostServiceReportQuery {
	query := (&HostServiceClient{config: hsrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hsrq.withHostservice = query
	return hsrq
}

// WithService tells the query-builder to eager-load the nodes that are connected to
// the "service" edge. The optional arguments are used to configure the query builder of the edge.
func (hsrq *HostServiceReportQuery) WithService(opts ...func(*ServiceQuery)) *HostServiceReportQuery {
	query := (&ServiceClient{config: hsrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hsrq.withService = query
	return hsrq
}

// WithTeam tells the query-builder to eager-load the nodes that are connected to
// the "team" edge. The optional arguments are used to configure the query builder of the edge.
func (hsrq *HostServiceReportQuery) WithTeam(opts ...func(*TeamQuery)) *HostServiceReportQuery {
	query := (&TeamClient{config: hsrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hsrq.withTeam = query
	return hsrq
}

// WithTeamreport tells the query-builder to eager-load the nodes that are connected to
// the "teamreport" edge. The optional arguments are used to configure the query builder of the edge.
func (hsrq *HostServiceReportQuery) WithTeamreport(opts ...func(*TeamReportQuery)) *HostServiceReportQuery {
	query := (&TeamReportClient{config: hsrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	hsrq.withTeamreport = query
	return hsrq
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
//	client.HostServiceReport.Query().
//		GroupBy(hostservicereport.FieldCreateTime).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (hsrq *HostServiceReportQuery) GroupBy(field string, fields ...string) *HostServiceReportGroupBy {
	hsrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &HostServiceReportGroupBy{build: hsrq}
	grbuild.flds = &hsrq.ctx.Fields
	grbuild.label = hostservicereport.Label
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
//	client.HostServiceReport.Query().
//		Select(hostservicereport.FieldCreateTime).
//		Scan(ctx, &v)
func (hsrq *HostServiceReportQuery) Select(fields ...string) *HostServiceReportSelect {
	hsrq.ctx.Fields = append(hsrq.ctx.Fields, fields...)
	sbuild := &HostServiceReportSelect{HostServiceReportQuery: hsrq}
	sbuild.label = hostservicereport.Label
	sbuild.flds, sbuild.scan = &hsrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a HostServiceReportSelect configured with the given aggregations.
func (hsrq *HostServiceReportQuery) Aggregate(fns ...AggregateFunc) *HostServiceReportSelect {
	return hsrq.Select().Aggregate(fns...)
}

func (hsrq *HostServiceReportQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range hsrq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, hsrq); err != nil {
				return err
			}
		}
	}
	for _, f := range hsrq.ctx.Fields {
		if !hostservicereport.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if hsrq.path != nil {
		prev, err := hsrq.path(ctx)
		if err != nil {
			return err
		}
		hsrq.sql = prev
	}
	return nil
}

func (hsrq *HostServiceReportQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*HostServiceReport, error) {
	var (
		nodes       = []*HostServiceReport{}
		_spec       = hsrq.querySpec()
		loadedTypes = [4]bool{
			hsrq.withHostservice != nil,
			hsrq.withService != nil,
			hsrq.withTeam != nil,
			hsrq.withTeamreport != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*HostServiceReport).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &HostServiceReport{config: hsrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hsrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := hsrq.withHostservice; query != nil {
		if err := hsrq.loadHostservice(ctx, query, nodes, nil,
			func(n *HostServiceReport, e *HostService) { n.Edges.Hostservice = e }); err != nil {
			return nil, err
		}
	}
	if query := hsrq.withService; query != nil {
		if err := hsrq.loadService(ctx, query, nodes, nil,
			func(n *HostServiceReport, e *Service) { n.Edges.Service = e }); err != nil {
			return nil, err
		}
	}
	if query := hsrq.withTeam; query != nil {
		if err := hsrq.loadTeam(ctx, query, nodes, nil,
			func(n *HostServiceReport, e *Team) { n.Edges.Team = e }); err != nil {
			return nil, err
		}
	}
	if query := hsrq.withTeamreport; query != nil {
		if err := hsrq.loadTeamreport(ctx, query, nodes, nil,
			func(n *HostServiceReport, e *TeamReport) { n.Edges.Teamreport = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (hsrq *HostServiceReportQuery) loadHostservice(ctx context.Context, query *HostServiceQuery, nodes []*HostServiceReport, init func(*HostServiceReport), assign func(*HostServiceReport, *HostService)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*HostServiceReport)
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
func (hsrq *HostServiceReportQuery) loadService(ctx context.Context, query *ServiceQuery, nodes []*HostServiceReport, init func(*HostServiceReport), assign func(*HostServiceReport, *Service)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*HostServiceReport)
	for i := range nodes {
		fk := nodes[i].ServiceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(service.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "service_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (hsrq *HostServiceReportQuery) loadTeam(ctx context.Context, query *TeamQuery, nodes []*HostServiceReport, init func(*HostServiceReport), assign func(*HostServiceReport, *Team)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*HostServiceReport)
	for i := range nodes {
		fk := nodes[i].TeamID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(team.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (hsrq *HostServiceReportQuery) loadTeamreport(ctx context.Context, query *TeamReportQuery, nodes []*HostServiceReport, init func(*HostServiceReport), assign func(*HostServiceReport, *TeamReport)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*HostServiceReport)
	for i := range nodes {
		fk := nodes[i].TeamReportID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(teamreport.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "team_report_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (hsrq *HostServiceReportQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hsrq.querySpec()
	_spec.Node.Columns = hsrq.ctx.Fields
	if len(hsrq.ctx.Fields) > 0 {
		_spec.Unique = hsrq.ctx.Unique != nil && *hsrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, hsrq.driver, _spec)
}

func (hsrq *HostServiceReportQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(hostservicereport.Table, hostservicereport.Columns, sqlgraph.NewFieldSpec(hostservicereport.FieldID, field.TypeString))
	_spec.From = hsrq.sql
	if unique := hsrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if hsrq.path != nil {
		_spec.Unique = true
	}
	if fields := hsrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hostservicereport.FieldID)
		for i := range fields {
			if fields[i] != hostservicereport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if hsrq.withHostservice != nil {
			_spec.Node.AddColumnOnce(hostservicereport.FieldHostServiceID)
		}
		if hsrq.withService != nil {
			_spec.Node.AddColumnOnce(hostservicereport.FieldServiceID)
		}
		if hsrq.withTeam != nil {
			_spec.Node.AddColumnOnce(hostservicereport.FieldTeamID)
		}
		if hsrq.withTeamreport != nil {
			_spec.Node.AddColumnOnce(hostservicereport.FieldTeamReportID)
		}
	}
	if ps := hsrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hsrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hsrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hsrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hsrq *HostServiceReportQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hsrq.driver.Dialect())
	t1 := builder.Table(hostservicereport.Table)
	columns := hsrq.ctx.Fields
	if len(columns) == 0 {
		columns = hostservicereport.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hsrq.sql != nil {
		selector = hsrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hsrq.ctx.Unique != nil && *hsrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range hsrq.predicates {
		p(selector)
	}
	for _, p := range hsrq.order {
		p(selector)
	}
	if offset := hsrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hsrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HostServiceReportGroupBy is the group-by builder for HostServiceReport entities.
type HostServiceReportGroupBy struct {
	selector
	build *HostServiceReportQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hsrgb *HostServiceReportGroupBy) Aggregate(fns ...AggregateFunc) *HostServiceReportGroupBy {
	hsrgb.fns = append(hsrgb.fns, fns...)
	return hsrgb
}

// Scan applies the selector query and scans the result into the given value.
func (hsrgb *HostServiceReportGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hsrgb.build.ctx, "GroupBy")
	if err := hsrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HostServiceReportQuery, *HostServiceReportGroupBy](ctx, hsrgb.build, hsrgb, hsrgb.build.inters, v)
}

func (hsrgb *HostServiceReportGroupBy) sqlScan(ctx context.Context, root *HostServiceReportQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(hsrgb.fns))
	for _, fn := range hsrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*hsrgb.flds)+len(hsrgb.fns))
		for _, f := range *hsrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*hsrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hsrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// HostServiceReportSelect is the builder for selecting fields of HostServiceReport entities.
type HostServiceReportSelect struct {
	*HostServiceReportQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (hsrs *HostServiceReportSelect) Aggregate(fns ...AggregateFunc) *HostServiceReportSelect {
	hsrs.fns = append(hsrs.fns, fns...)
	return hsrs
}

// Scan applies the selector query and scans the result into the given value.
func (hsrs *HostServiceReportSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, hsrs.ctx, "Select")
	if err := hsrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*HostServiceReportQuery, *HostServiceReportSelect](ctx, hsrs.HostServiceReportQuery, hsrs, hsrs.inters, v)
}

func (hsrs *HostServiceReportSelect) sqlScan(ctx context.Context, root *HostServiceReportQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(hsrs.fns))
	for _, fn := range hsrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*hsrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hsrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
