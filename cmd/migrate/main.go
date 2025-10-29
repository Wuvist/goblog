package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "modernc.org/sqlite"
)

func main() {
	mysqlDefault := os.Getenv("MYSQL_CONNSTR")
	if mysqlDefault == "" {
		mysqlDefault = os.Getenv("connstr")
	}
	sqliteDefault := os.Getenv("SQLITE_DSN")
	if sqliteDefault == "" {
		sqliteDefault = "file:blogwind.db?_pragma=foreign_keys(OFF)"
	}

	mysqlDSN := flag.String("mysql", mysqlDefault, "MySQL DSN, e.g. user:pass@tcp(localhost:3306)/blogwind?parseTime=true")
	sqliteDSN := flag.String("sqlite", sqliteDefault, "SQLite DSN or file path (modernc driver syntax)")
	flag.Parse()

	if *mysqlDSN == "" {
		log.Fatal("missing MySQL DSN (set --mysql or MYSQL_CONNSTR)")
	}

	normalizedSQLite := normalizeSQLiteDSN(*sqliteDSN)

	ctx := context.Background()

	mysqlDB, err := sql.Open("mysql", *mysqlDSN)
	if err != nil {
		log.Fatalf("open mysql: %v", err)
	}
	defer mysqlDB.Close()

	if err := mysqlDB.PingContext(ctx); err != nil {
		log.Fatalf("ping mysql: %v", err)
	}

	sqliteDB, err := sql.Open("sqlite", normalizedSQLite)
	if err != nil {
		log.Fatalf("open sqlite: %v", err)
	}
	defer sqliteDB.Close()

	if err := sqliteDB.PingContext(ctx); err != nil {
		log.Fatalf("ping sqlite: %v", err)
	}

	if _, err := sqliteDB.ExecContext(ctx, "PRAGMA foreign_keys = OFF"); err != nil {
		log.Fatalf("disable foreign_keys: %v", err)
	}

	tables := []func(context.Context, *sql.DB, *sql.DB) error{
		copyBlogger,
		copyArticles,
		copyComment,
		copyLinks,
		copyUserdefinecategory,
	}

	for _, fn := range tables {
		if err := fn(ctx, mysqlDB, sqliteDB); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}

	log.Println("migration complete")
}

func normalizeSQLiteDSN(dsn string) string {
	if strings.HasPrefix(dsn, "file:") || strings.Contains(dsn, "://") {
		return dsn
	}
	if strings.Contains(dsn, "?") {
		return fmt.Sprintf("file:%s", dsn)
	}
	return fmt.Sprintf("file:%s?_pragma=foreign_keys(OFF)", dsn)
}

func copyBlogger(ctx context.Context, mysqlDB, sqliteDB *sql.DB) error {
	const selectSQL = "SELECT `index`, id, nick, DOB, blogname, blogskin, email, visitor, intro, blogs, IP, TS, Last_log, Last_post, Activate, Reveal, userpic, lang, widget FROM blogger"
	const insertSQL = `INSERT INTO blogger ("index", id, nick, DOB, blogname, blogskin, email, visitor, intro, blogs, IP, TS, Last_log, Last_post, Activate, Reveal, userpic, lang, widget)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	return copyTable(ctx, mysqlDB, sqliteDB, "blogger", selectSQL, insertSQL,
		func(rows *sql.Rows) ([]interface{}, error) {
			var (
				index    int64
				id       string
				nick     sql.NullString
				dob      time.Time
				blogname string
				blogskin int64
				email    string
				visitor  int64
				intro    sql.NullString
				blogs    int64
				ip       sql.NullString
				ts       int64
				lastLog  time.Time
				lastPost sql.NullTime
				activate int64
				reveal   int64
				userpic  int64
				lang     string
				widget   sql.NullString
			)
			if err := rows.Scan(&index, &id, &nick, &dob, &blogname, &blogskin, &email, &visitor, &intro, &blogs, &ip, &ts, &lastLog, &lastPost, &activate, &reveal, &userpic, &lang, &widget); err != nil {
				return nil, err
			}
			return []interface{}{
				index,
				id,
				nullString(nick),
				dob,
				blogname,
				blogskin,
				email,
				visitor,
				nullString(intro),
				blogs,
				nullString(ip),
				ts,
				lastLog,
				nullTime(lastPost),
				activate,
				reveal,
				userpic,
				lang,
				nullString(widget),
			}, nil
		})
}

func copyArticles(ctx context.Context, mysqlDB, sqliteDB *sql.DB) error {
	const selectSQL = "SELECT `index`, title, blogger, content, ref, add_date, cate, Comment, Recommend FROM articles"
	const insertSQL = `INSERT INTO articles ("index", title, blogger, content, ref, add_date, cate, Comment, Recommend)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	return copyTable(ctx, mysqlDB, sqliteDB, "articles", selectSQL, insertSQL,
		func(rows *sql.Rows) ([]interface{}, error) {
			var (
				index     int64
				title     sql.NullString
				blogger   int64
				content   sql.NullString
				ref       int64
				addDate   sql.NullTime
				cate      sql.NullInt64
				comment   int64
				recommend int64
			)
			if err := rows.Scan(&index, &title, &blogger, &content, &ref, &addDate, &cate, &comment, &recommend); err != nil {
				return nil, err
			}
			return []interface{}{
				index,
				nullString(title),
				blogger,
				nullString(content),
				ref,
				nullTime(addDate),
				nullInt(cate),
				comment,
				recommend,
			}, nil
		})
}

func copyComment(ctx context.Context, mysqlDB, sqliteDB *sql.DB) error {
	const selectSQL = "SELECT `index`, blogger, article, title, content, author, homepage, add_date, ip, utype, uid FROM comment"
	const insertSQL = `INSERT INTO comment ("index", blogger, article, title, content, author, homepage, add_date, ip, utype, uid)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	return copyTable(ctx, mysqlDB, sqliteDB, "comment", selectSQL, insertSQL,
		func(rows *sql.Rows) ([]interface{}, error) {
			var (
				index    int64
				blogger  sql.NullInt64
				article  int64
				title    string
				content  string
				author   sql.NullString
				homepage sql.NullString
				addDate  time.Time
				ip       sql.NullString
				utype    int64
				uid      sql.NullString
			)
			if err := rows.Scan(&index, &blogger, &article, &title, &content, &author, &homepage, &addDate, &ip, &utype, &uid); err != nil {
				return nil, err
			}
			return []interface{}{
				index,
				nullInt(blogger),
				article,
				title,
				content,
				nullString(author),
				nullString(homepage),
				addDate,
				nullString(ip),
				utype,
				nullString(uid),
			}, nil
		})
}

func copyLinks(ctx context.Context, mysqlDB, sqliteDB *sql.DB) error {
	const selectSQL = "SELECT id, blogger, link, URL, reveal FROM links"
	const insertSQL = `INSERT INTO links (id, blogger, link, URL, reveal)
VALUES (?, ?, ?, ?, ?)`

	return copyTable(ctx, mysqlDB, sqliteDB, "links", selectSQL, insertSQL,
		func(rows *sql.Rows) ([]interface{}, error) {
			var (
				id      int64
				blogger int64
				link    string
				url     string
				reveal  int64
			)
			if err := rows.Scan(&id, &blogger, &link, &url, &reveal); err != nil {
				return nil, err
			}
			return []interface{}{
				id,
				blogger,
				link,
				url,
				reveal,
			}, nil
		})
}

func copyUserdefinecategory(ctx context.Context, mysqlDB, sqliteDB *sql.DB) error {
	const selectSQL = "SELECT `index`, blogger, cate, files FROM userdefinecategory"
	const insertSQL = `INSERT INTO userdefinecategory ("index", blogger, cate, files)
VALUES (?, ?, ?, ?)`

	return copyTable(ctx, mysqlDB, sqliteDB, "userdefinecategory", selectSQL, insertSQL,
		func(rows *sql.Rows) ([]interface{}, error) {
			var (
				index   int64
				blogger int64
				cate    string
				files   int64
			)
			if err := rows.Scan(&index, &blogger, &cate, &files); err != nil {
				return nil, err
			}
			return []interface{}{
				index,
				blogger,
				cate,
				files,
			}, nil
		})
}

type rowMapper func(*sql.Rows) ([]interface{}, error)

func copyTable(ctx context.Context, mysqlDB, sqliteDB *sql.DB, tableName, selectSQL, insertSQL string, mapper rowMapper) error {
	log.Printf("copying %s...", tableName)

	tx, err := sqliteDB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("%s: begin tx: %w", tableName, err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s", tableName)); err != nil {
		return fmt.Errorf("%s: truncate: %w", tableName, err)
	}

	stmt, err := tx.PrepareContext(ctx, insertSQL)
	if err != nil {
		return fmt.Errorf("%s: prepare insert: %w", tableName, err)
	}
	defer stmt.Close()

	rows, err := mysqlDB.QueryContext(ctx, selectSQL)
	if err != nil {
		return fmt.Errorf("%s: select: %w", tableName, err)
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		args, err := mapper(rows)
		if err != nil {
			return fmt.Errorf("%s: map row: %w", tableName, err)
		}
		if _, err := stmt.ExecContext(ctx, args...); err != nil {
			return fmt.Errorf("%s: insert: %w", tableName, err)
		}
		count++
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("%s: rows: %w", tableName, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("%s: commit: %w", tableName, err)
	}

	log.Printf("copied %s rows: %d", tableName, count)
	return nil
}

func nullString(ns sql.NullString) interface{} {
	if ns.Valid {
		return ns.String
	}
	return nil
}

func nullTime(nt sql.NullTime) interface{} {
	if nt.Valid {
		return nt.Time
	}
	return nil
}

func nullInt(ni sql.NullInt64) interface{} {
	if ni.Valid {
		return ni.Int64
	}
	return nil
}
