package db

/*
 This file was created by mkdb-client.
 The intention is not to modify thils file, but you may extend the struct DBActivityLog
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence activitylog_seq;

Main Table:

 CREATE TABLE activitylog (id integer primary key default nextval('activitylog_seq'),ip text not null  ,userid text not null  ,email text not null  ,triggerhost text not null  ,occured integer not null  ,logmessage text not null  );

Alter statements:
ALTER TABLE activitylog ADD COLUMN ip text not null default '';
ALTER TABLE activitylog ADD COLUMN userid text not null default '';
ALTER TABLE activitylog ADD COLUMN email text not null default '';
ALTER TABLE activitylog ADD COLUMN triggerhost text not null default '';
ALTER TABLE activitylog ADD COLUMN occured integer not null default 0;
ALTER TABLE activitylog ADD COLUMN logmessage text not null default '';


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE activitylog_archive (id integer unique not null,ip text not null,userid text not null,email text not null,triggerhost text not null,occured integer not null,logmessage text not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/weblogin"
	"golang.conradwood.net/go-easyops/sql"
	"os"
)

var (
	default_def_DBActivityLog *DBActivityLog
)

type DBActivityLog struct {
	DB                  *sql.DB
	SQLTablename        string
	SQLArchivetablename string
}

func DefaultDBActivityLog() *DBActivityLog {
	if default_def_DBActivityLog != nil {
		return default_def_DBActivityLog
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBActivityLog(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBActivityLog = res
	return res
}
func NewDBActivityLog(db *sql.DB) *DBActivityLog {
	foo := DBActivityLog{DB: db}
	foo.SQLTablename = "activitylog"
	foo.SQLArchivetablename = "activitylog_archive"
	return &foo
}

// archive. It is NOT transactionally save.
func (a *DBActivityLog) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "archive_DBActivityLog", "insert into "+a.SQLArchivetablename+"+ (id,ip, userid, email, triggerhost, occured, logmessage) values ($1,$2, $3, $4, $5, $6, $7) ", p.ID, p.IP, p.UserID, p.Email, p.TriggerHost, p.Occured, p.LogMessage)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// Save (and use database default ID generation)
func (a *DBActivityLog) Save(ctx context.Context, p *savepb.ActivityLog) (uint64, error) {
	qn := "DBActivityLog_Save"
	rows, e := a.DB.QueryContext(ctx, qn, "insert into "+a.SQLTablename+" (ip, userid, email, triggerhost, occured, logmessage) values ($1, $2, $3, $4, $5, $6) returning id", p.IP, p.UserID, p.Email, p.TriggerHost, p.Occured, p.LogMessage)
	if e != nil {
		return 0, a.Error(ctx, qn, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, qn, fmt.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, qn, fmt.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// Save using the ID specified
func (a *DBActivityLog) SaveWithID(ctx context.Context, p *savepb.ActivityLog) error {
	qn := "insert_DBActivityLog"
	_, e := a.DB.ExecContext(ctx, qn, "insert into "+a.SQLTablename+" (id,ip, userid, email, triggerhost, occured, logmessage) values ($1,$2, $3, $4, $5, $6, $7) ", p.ID, p.IP, p.UserID, p.Email, p.TriggerHost, p.Occured, p.LogMessage)
	return a.Error(ctx, qn, e)
}

func (a *DBActivityLog) Update(ctx context.Context, p *savepb.ActivityLog) error {
	qn := "DBActivityLog_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set ip=$1, userid=$2, email=$3, triggerhost=$4, occured=$5, logmessage=$6 where id = $7", p.IP, p.UserID, p.Email, p.TriggerHost, p.Occured, p.LogMessage, p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBActivityLog) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBActivityLog_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBActivityLog) ByID(ctx context.Context, p uint64) (*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, fmt.Errorf("No ActivityLog with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, fmt.Errorf("Multiple (%d) ActivityLog with id %v", len(l), p))
	}
	return l[0], nil
}

// get all rows
func (a *DBActivityLog) All(ctx context.Context) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_all"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" order by id")
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("All: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, fmt.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBActivityLog" rows with matching IP
func (a *DBActivityLog) ByIP(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByIP"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where ip = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBActivityLog) ByLikeIP(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLikeIP"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where ip ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByIP: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBActivityLog" rows with matching UserID
func (a *DBActivityLog) ByUserID(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByUserID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where userid = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBActivityLog) ByLikeUserID(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLikeUserID"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where userid ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByUserID: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBActivityLog" rows with matching Email
func (a *DBActivityLog) ByEmail(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByEmail"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where email = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEmail: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEmail: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBActivityLog) ByLikeEmail(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLikeEmail"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where email ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEmail: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByEmail: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBActivityLog" rows with matching TriggerHost
func (a *DBActivityLog) ByTriggerHost(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByTriggerHost"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where triggerhost = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBActivityLog) ByLikeTriggerHost(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLikeTriggerHost"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where triggerhost ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByTriggerHost: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBActivityLog" rows with matching Occured
func (a *DBActivityLog) ByOccured(ctx context.Context, p uint32) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByOccured"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where occured = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByOccured: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByOccured: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBActivityLog) ByLikeOccured(ctx context.Context, p uint32) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLikeOccured"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where occured ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByOccured: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByOccured: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBActivityLog" rows with matching LogMessage
func (a *DBActivityLog) ByLogMessage(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLogMessage"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where logmessage = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLogMessage: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLogMessage: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBActivityLog) ByLikeLogMessage(ctx context.Context, p string) ([]*savepb.ActivityLog, error) {
	qn := "DBActivityLog_ByLikeLogMessage"
	rows, e := a.DB.QueryContext(ctx, qn, "select id,ip, userid, email, triggerhost, occured, logmessage from "+a.SQLTablename+" where logmessage ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLogMessage: error querying (%s)", e))
	}
	defer rows.Close()
	l, e := a.FromRows(ctx, rows)
	if e != nil {
		return nil, a.Error(ctx, qn, fmt.Errorf("ByLogMessage: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBActivityLog) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.ActivityLog, error) {
	rows, err := a.DB.QueryContext(ctx, "custom_query_"+a.Tablename(), "select "+a.SelectCols()+" from "+a.Tablename()+" where "+query_where, args...)
	if err != nil {
		return nil, err
	}
	return a.FromRows(ctx, rows)
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBActivityLog) Tablename() string {
	return a.SQLTablename
}

func (a *DBActivityLog) SelectCols() string {
	return "id,ip, userid, email, triggerhost, occured, logmessage"
}
func (a *DBActivityLog) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".ip, " + a.SQLTablename + ".userid, " + a.SQLTablename + ".email, " + a.SQLTablename + ".triggerhost, " + a.SQLTablename + ".occured, " + a.SQLTablename + ".logmessage"
}

func (a *DBActivityLog) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.ActivityLog, error) {
	var res []*savepb.ActivityLog
	for rows.Next() {
		foo := savepb.ActivityLog{}
		err := rows.Scan(&foo.ID, &foo.IP, &foo.UserID, &foo.Email, &foo.TriggerHost, &foo.Occured, &foo.LogMessage)
		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
		}
		res = append(res, &foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBActivityLog) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),ip text not null  ,userid text not null  ,email text not null  ,triggerhost text not null  ,occured integer not null  ,logmessage text not null  );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),ip text not null  ,userid text not null  ,email text not null  ,triggerhost text not null  ,occured integer not null  ,logmessage text not null  );`,
	}
	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}
	return nil
}

/**********************************************************************
* Helper to meaningful errors
**********************************************************************/
func (a *DBActivityLog) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return fmt.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}
