// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GamePlayer is an object representing the database table.
type GamePlayer struct {
	UUID     string      `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	GameUUID null.String `boil:"game_uuid" json:"game_uuid,omitempty" toml:"game_uuid" yaml:"game_uuid,omitempty"`
	UserUUID null.String `boil:"user_uuid" json:"user_uuid,omitempty" toml:"user_uuid" yaml:"user_uuid,omitempty"`
	Name     string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	HasShot  null.Bool   `boil:"has_shot" json:"has_shot,omitempty" toml:"has_shot" yaml:"has_shot,omitempty"`
	IsAlive  null.Bool   `boil:"is_alive" json:"is_alive,omitempty" toml:"is_alive" yaml:"is_alive,omitempty"`

	R *gamePlayerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L gamePlayerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GamePlayerColumns = struct {
	UUID     string
	GameUUID string
	UserUUID string
	Name     string
	HasShot  string
	IsAlive  string
}{
	UUID:     "uuid",
	GameUUID: "game_uuid",
	UserUUID: "user_uuid",
	Name:     "name",
	HasShot:  "has_shot",
	IsAlive:  "is_alive",
}

var GamePlayerTableColumns = struct {
	UUID     string
	GameUUID string
	UserUUID string
	Name     string
	HasShot  string
	IsAlive  string
}{
	UUID:     "game_players.uuid",
	GameUUID: "game_players.game_uuid",
	UserUUID: "game_players.user_uuid",
	Name:     "game_players.name",
	HasShot:  "game_players.has_shot",
	IsAlive:  "game_players.is_alive",
}

// Generated where

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var GamePlayerWhere = struct {
	UUID     whereHelperstring
	GameUUID whereHelpernull_String
	UserUUID whereHelpernull_String
	Name     whereHelperstring
	HasShot  whereHelpernull_Bool
	IsAlive  whereHelpernull_Bool
}{
	UUID:     whereHelperstring{field: "\"game_players\".\"uuid\""},
	GameUUID: whereHelpernull_String{field: "\"game_players\".\"game_uuid\""},
	UserUUID: whereHelpernull_String{field: "\"game_players\".\"user_uuid\""},
	Name:     whereHelperstring{field: "\"game_players\".\"name\""},
	HasShot:  whereHelpernull_Bool{field: "\"game_players\".\"has_shot\""},
	IsAlive:  whereHelpernull_Bool{field: "\"game_players\".\"is_alive\""},
}

// GamePlayerRels is where relationship names are stored.
var GamePlayerRels = struct {
	Game string
	User string
}{
	Game: "Game",
	User: "User",
}

// gamePlayerR is where relationships are stored.
type gamePlayerR struct {
	Game *Game `boil:"Game" json:"Game" toml:"Game" yaml:"Game"`
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*gamePlayerR) NewStruct() *gamePlayerR {
	return &gamePlayerR{}
}

func (r *gamePlayerR) GetGame() *Game {
	if r == nil {
		return nil
	}
	return r.Game
}

func (r *gamePlayerR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// gamePlayerL is where Load methods for each relationship are stored.
type gamePlayerL struct{}

var (
	gamePlayerAllColumns            = []string{"uuid", "game_uuid", "user_uuid", "name", "has_shot", "is_alive"}
	gamePlayerColumnsWithoutDefault = []string{}
	gamePlayerColumnsWithDefault    = []string{"uuid", "game_uuid", "user_uuid", "name", "has_shot", "is_alive"}
	gamePlayerPrimaryKeyColumns     = []string{"uuid"}
	gamePlayerGeneratedColumns      = []string{}
)

type (
	// GamePlayerSlice is an alias for a slice of pointers to GamePlayer.
	// This should almost always be used instead of []GamePlayer.
	GamePlayerSlice []*GamePlayer
	// GamePlayerHook is the signature for custom GamePlayer hook methods
	GamePlayerHook func(context.Context, boil.ContextExecutor, *GamePlayer) error

	gamePlayerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	gamePlayerType                 = reflect.TypeOf(&GamePlayer{})
	gamePlayerMapping              = queries.MakeStructMapping(gamePlayerType)
	gamePlayerPrimaryKeyMapping, _ = queries.BindMapping(gamePlayerType, gamePlayerMapping, gamePlayerPrimaryKeyColumns)
	gamePlayerInsertCacheMut       sync.RWMutex
	gamePlayerInsertCache          = make(map[string]insertCache)
	gamePlayerUpdateCacheMut       sync.RWMutex
	gamePlayerUpdateCache          = make(map[string]updateCache)
	gamePlayerUpsertCacheMut       sync.RWMutex
	gamePlayerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var gamePlayerAfterSelectMu sync.Mutex
var gamePlayerAfterSelectHooks []GamePlayerHook

var gamePlayerBeforeInsertMu sync.Mutex
var gamePlayerBeforeInsertHooks []GamePlayerHook
var gamePlayerAfterInsertMu sync.Mutex
var gamePlayerAfterInsertHooks []GamePlayerHook

var gamePlayerBeforeUpdateMu sync.Mutex
var gamePlayerBeforeUpdateHooks []GamePlayerHook
var gamePlayerAfterUpdateMu sync.Mutex
var gamePlayerAfterUpdateHooks []GamePlayerHook

var gamePlayerBeforeDeleteMu sync.Mutex
var gamePlayerBeforeDeleteHooks []GamePlayerHook
var gamePlayerAfterDeleteMu sync.Mutex
var gamePlayerAfterDeleteHooks []GamePlayerHook

var gamePlayerBeforeUpsertMu sync.Mutex
var gamePlayerBeforeUpsertHooks []GamePlayerHook
var gamePlayerAfterUpsertMu sync.Mutex
var gamePlayerAfterUpsertHooks []GamePlayerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GamePlayer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GamePlayer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GamePlayer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GamePlayer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GamePlayer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GamePlayer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GamePlayer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GamePlayer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GamePlayer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range gamePlayerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGamePlayerHook registers your hook function for all future operations.
func AddGamePlayerHook(hookPoint boil.HookPoint, gamePlayerHook GamePlayerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		gamePlayerAfterSelectMu.Lock()
		gamePlayerAfterSelectHooks = append(gamePlayerAfterSelectHooks, gamePlayerHook)
		gamePlayerAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		gamePlayerBeforeInsertMu.Lock()
		gamePlayerBeforeInsertHooks = append(gamePlayerBeforeInsertHooks, gamePlayerHook)
		gamePlayerBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		gamePlayerAfterInsertMu.Lock()
		gamePlayerAfterInsertHooks = append(gamePlayerAfterInsertHooks, gamePlayerHook)
		gamePlayerAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		gamePlayerBeforeUpdateMu.Lock()
		gamePlayerBeforeUpdateHooks = append(gamePlayerBeforeUpdateHooks, gamePlayerHook)
		gamePlayerBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		gamePlayerAfterUpdateMu.Lock()
		gamePlayerAfterUpdateHooks = append(gamePlayerAfterUpdateHooks, gamePlayerHook)
		gamePlayerAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		gamePlayerBeforeDeleteMu.Lock()
		gamePlayerBeforeDeleteHooks = append(gamePlayerBeforeDeleteHooks, gamePlayerHook)
		gamePlayerBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		gamePlayerAfterDeleteMu.Lock()
		gamePlayerAfterDeleteHooks = append(gamePlayerAfterDeleteHooks, gamePlayerHook)
		gamePlayerAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		gamePlayerBeforeUpsertMu.Lock()
		gamePlayerBeforeUpsertHooks = append(gamePlayerBeforeUpsertHooks, gamePlayerHook)
		gamePlayerBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		gamePlayerAfterUpsertMu.Lock()
		gamePlayerAfterUpsertHooks = append(gamePlayerAfterUpsertHooks, gamePlayerHook)
		gamePlayerAfterUpsertMu.Unlock()
	}
}

// One returns a single gamePlayer record from the query.
func (q gamePlayerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GamePlayer, error) {
	o := &GamePlayer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for game_players")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GamePlayer records from the query.
func (q gamePlayerQuery) All(ctx context.Context, exec boil.ContextExecutor) (GamePlayerSlice, error) {
	var o []*GamePlayer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GamePlayer slice")
	}

	if len(gamePlayerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GamePlayer records in the query.
func (q gamePlayerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count game_players rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q gamePlayerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if game_players exists")
	}

	return count > 0, nil
}

// Game pointed to by the foreign key.
func (o *GamePlayer) Game(mods ...qm.QueryMod) gameQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"uuid\" = ?", o.GameUUID),
	}

	queryMods = append(queryMods, mods...)

	return Games(queryMods...)
}

// User pointed to by the foreign key.
func (o *GamePlayer) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"uuid\" = ?", o.UserUUID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadGame allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (gamePlayerL) LoadGame(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGamePlayer interface{}, mods queries.Applicator) error {
	var slice []*GamePlayer
	var object *GamePlayer

	if singular {
		var ok bool
		object, ok = maybeGamePlayer.(*GamePlayer)
		if !ok {
			object = new(GamePlayer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGamePlayer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGamePlayer))
			}
		}
	} else {
		s, ok := maybeGamePlayer.(*[]*GamePlayer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGamePlayer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGamePlayer))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &gamePlayerR{}
		}
		if !queries.IsNil(object.GameUUID) {
			args[object.GameUUID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gamePlayerR{}
			}

			if !queries.IsNil(obj.GameUUID) {
				args[obj.GameUUID] = struct{}{}
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`game`),
		qm.WhereIn(`game.uuid in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Game")
	}

	var resultSlice []*Game
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Game")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for game")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for game")
	}

	if len(gameAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Game = foreign
		if foreign.R == nil {
			foreign.R = &gameR{}
		}
		foreign.R.GamePlayers = append(foreign.R.GamePlayers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.GameUUID, foreign.UUID) {
				local.R.Game = foreign
				if foreign.R == nil {
					foreign.R = &gameR{}
				}
				foreign.R.GamePlayers = append(foreign.R.GamePlayers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (gamePlayerL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGamePlayer interface{}, mods queries.Applicator) error {
	var slice []*GamePlayer
	var object *GamePlayer

	if singular {
		var ok bool
		object, ok = maybeGamePlayer.(*GamePlayer)
		if !ok {
			object = new(GamePlayer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGamePlayer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGamePlayer))
			}
		}
	} else {
		s, ok := maybeGamePlayer.(*[]*GamePlayer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGamePlayer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGamePlayer))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &gamePlayerR{}
		}
		if !queries.IsNil(object.UserUUID) {
			args[object.UserUUID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &gamePlayerR{}
			}

			if !queries.IsNil(obj.UserUUID) {
				args[obj.UserUUID] = struct{}{}
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.uuid in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.GamePlayers = append(foreign.R.GamePlayers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserUUID, foreign.UUID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.GamePlayers = append(foreign.R.GamePlayers, local)
				break
			}
		}
	}

	return nil
}

// SetGame of the gamePlayer to the related item.
// Sets o.R.Game to related.
// Adds o to related.R.GamePlayers.
func (o *GamePlayer) SetGame(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Game) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"game_players\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"game_uuid"}),
		strmangle.WhereClause("\"", "\"", 2, gamePlayerPrimaryKeyColumns),
	)
	values := []interface{}{related.UUID, o.UUID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.GameUUID, related.UUID)
	if o.R == nil {
		o.R = &gamePlayerR{
			Game: related,
		}
	} else {
		o.R.Game = related
	}

	if related.R == nil {
		related.R = &gameR{
			GamePlayers: GamePlayerSlice{o},
		}
	} else {
		related.R.GamePlayers = append(related.R.GamePlayers, o)
	}

	return nil
}

// RemoveGame relationship.
// Sets o.R.Game to nil.
// Removes o from all passed in related items' relationships struct.
func (o *GamePlayer) RemoveGame(ctx context.Context, exec boil.ContextExecutor, related *Game) error {
	var err error

	queries.SetScanner(&o.GameUUID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("game_uuid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Game = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.GamePlayers {
		if queries.Equal(o.GameUUID, ri.GameUUID) {
			continue
		}

		ln := len(related.R.GamePlayers)
		if ln > 1 && i < ln-1 {
			related.R.GamePlayers[i] = related.R.GamePlayers[ln-1]
		}
		related.R.GamePlayers = related.R.GamePlayers[:ln-1]
		break
	}
	return nil
}

// SetUser of the gamePlayer to the related item.
// Sets o.R.User to related.
// Adds o to related.R.GamePlayers.
func (o *GamePlayer) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"game_players\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_uuid"}),
		strmangle.WhereClause("\"", "\"", 2, gamePlayerPrimaryKeyColumns),
	)
	values := []interface{}{related.UUID, o.UUID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserUUID, related.UUID)
	if o.R == nil {
		o.R = &gamePlayerR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			GamePlayers: GamePlayerSlice{o},
		}
	} else {
		related.R.GamePlayers = append(related.R.GamePlayers, o)
	}

	return nil
}

// RemoveUser relationship.
// Sets o.R.User to nil.
// Removes o from all passed in related items' relationships struct.
func (o *GamePlayer) RemoveUser(ctx context.Context, exec boil.ContextExecutor, related *User) error {
	var err error

	queries.SetScanner(&o.UserUUID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("user_uuid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.User = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.GamePlayers {
		if queries.Equal(o.UserUUID, ri.UserUUID) {
			continue
		}

		ln := len(related.R.GamePlayers)
		if ln > 1 && i < ln-1 {
			related.R.GamePlayers[i] = related.R.GamePlayers[ln-1]
		}
		related.R.GamePlayers = related.R.GamePlayers[:ln-1]
		break
	}
	return nil
}

// GamePlayers retrieves all the records using an executor.
func GamePlayers(mods ...qm.QueryMod) gamePlayerQuery {
	mods = append(mods, qm.From("\"game_players\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"game_players\".*"})
	}

	return gamePlayerQuery{q}
}

// FindGamePlayer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGamePlayer(ctx context.Context, exec boil.ContextExecutor, uUID string, selectCols ...string) (*GamePlayer, error) {
	gamePlayerObj := &GamePlayer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"game_players\" where \"uuid\"=$1", sel,
	)

	q := queries.Raw(query, uUID)

	err := q.Bind(ctx, exec, gamePlayerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from game_players")
	}

	if err = gamePlayerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return gamePlayerObj, err
	}

	return gamePlayerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GamePlayer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no game_players provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gamePlayerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	gamePlayerInsertCacheMut.RLock()
	cache, cached := gamePlayerInsertCache[key]
	gamePlayerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			gamePlayerAllColumns,
			gamePlayerColumnsWithDefault,
			gamePlayerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(gamePlayerType, gamePlayerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(gamePlayerType, gamePlayerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"game_players\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"game_players\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into game_players")
	}

	if !cached {
		gamePlayerInsertCacheMut.Lock()
		gamePlayerInsertCache[key] = cache
		gamePlayerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GamePlayer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GamePlayer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	gamePlayerUpdateCacheMut.RLock()
	cache, cached := gamePlayerUpdateCache[key]
	gamePlayerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			gamePlayerAllColumns,
			gamePlayerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update game_players, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"game_players\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, gamePlayerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(gamePlayerType, gamePlayerMapping, append(wl, gamePlayerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update game_players row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for game_players")
	}

	if !cached {
		gamePlayerUpdateCacheMut.Lock()
		gamePlayerUpdateCache[key] = cache
		gamePlayerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q gamePlayerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for game_players")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for game_players")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GamePlayerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gamePlayerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"game_players\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, gamePlayerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in gamePlayer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all gamePlayer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GamePlayer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no game_players provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(gamePlayerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	gamePlayerUpsertCacheMut.RLock()
	cache, cached := gamePlayerUpsertCache[key]
	gamePlayerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			gamePlayerAllColumns,
			gamePlayerColumnsWithDefault,
			gamePlayerColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			gamePlayerAllColumns,
			gamePlayerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert game_players, could not build update column list")
		}

		ret := strmangle.SetComplement(gamePlayerAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(gamePlayerPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert game_players, could not build conflict column list")
			}

			conflict = make([]string, len(gamePlayerPrimaryKeyColumns))
			copy(conflict, gamePlayerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"game_players\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(gamePlayerType, gamePlayerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(gamePlayerType, gamePlayerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert game_players")
	}

	if !cached {
		gamePlayerUpsertCacheMut.Lock()
		gamePlayerUpsertCache[key] = cache
		gamePlayerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GamePlayer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GamePlayer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GamePlayer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), gamePlayerPrimaryKeyMapping)
	sql := "DELETE FROM \"game_players\" WHERE \"uuid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from game_players")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for game_players")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q gamePlayerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no gamePlayerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from game_players")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_players")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GamePlayerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(gamePlayerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gamePlayerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"game_players\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gamePlayerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from gamePlayer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for game_players")
	}

	if len(gamePlayerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *GamePlayer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGamePlayer(ctx, exec, o.UUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GamePlayerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GamePlayerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), gamePlayerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"game_players\".* FROM \"game_players\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, gamePlayerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GamePlayerSlice")
	}

	*o = slice

	return nil
}

// GamePlayerExists checks if the GamePlayer row exists.
func GamePlayerExists(ctx context.Context, exec boil.ContextExecutor, uUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"game_players\" where \"uuid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uUID)
	}
	row := exec.QueryRowContext(ctx, sql, uUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if game_players exists")
	}

	return exists, nil
}

// Exists checks if the GamePlayer row exists.
func (o *GamePlayer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GamePlayerExists(ctx, exec, o.UUID)
}
