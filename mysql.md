# Mysql

- import 
```go
import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"        
)
```

- connect 
```go
db, err := sql.Open("mysql", "username:password@tc(localhost:3306)/db_name?charset=utf8")
```

- insert 

```go
// stmt, err := db.Prepare("INSERT INTO users (username, role, status) VALUES(?, ?, ?"))

stmt, err := db.Prepare("INSERT users SET username=?, role=?, status=?")

res, err := stmt.Exec("gopher", "admin", "activate")

// lấy id của bản ghi vừa insert
id, err := res.LastInsertId()
```

- update
```go
stmt, err = db.Prepare("UPDATE users SET username=? WHERE role=?")

res, err = stmt.Exec("gopher", "admin")

// lấy số bản ghi bị ảnh hưởng
affect, err := res.RowsAffected()
```

- delete 
```go
stmt, err = db.Prepare("DELETE FROM users WHERE id=?")

res, err = stmt.Exec("1") // uid = 1

// lấy số bản ghi bị ảnh hưởng
affect, err = res.RowsAffected()
```

- select all
```go
rows, err := db.Query("SELECT username, role, status FROM users")

type User struct {
    Username string
    Role string
    Status string
}

var users []User

for rows.Next() {   
    var user User

    // chú ý: fields scan phải tương ứng với câu query lấy ra
    err = rows.Scan(&user.Username, &user.Role, &user.Status)    
}
```

- select one
```go
row, err := db.QueryRow("SELECT username, role, status FROM users WHERE id=?", 1)
```

- transaction, commit, rollback 
```go
func doubleInsert(db *sql.DB) error {

    defer showErr()

    // start transaction
    tx, err := db.Begin()
    checkErr(err)

    // group logic 01
    {
        var thing_1, whatever string

        // chuẩn bị truy vấn
        stmt, err := tx.Prepare(`INSERT INTO table_1 (thing_1, whatever) VALUES($1, $2);`)
        handleTransactionError(tx, err)

        // thực thi
        err = stmt.Exec(thing_1, whatever)
        handleTransactionError(tx, err)        
    }

    // group logic 02
    {
        var thing_2, whatever string
        stmt, err := tx.Prepare(`INSERT INTO table_2 (thing_2, whatever) VALUES($1, $2);`)
        handleTransactionError(tx, err)

        err = stmt.Exec(thing_2, whatever)
        handleTransactionError(tx, err)        
    }

    // nếu ko có lỗi, thì commit
    return tx.Commit()
}

func handleTransactionError(tx sql.Tx, err error){
    if err != nil {
        tx.Rollback()
        checkErr(err)
    }    
}

func checkErr(err error){
    panic(err)
}

func showErr(err error){
    if r := recover(); r != nil {
        fmt.Println("Error Msg", err)
    }
}

```