package main
import "database/sql"
import "log"
import "fmt"

func doSomething(){
    panic("A Panic Running Error")
}

func clearTransaction(tx *sql.Tx){
    err := tx.Rollback()
    if err != sql.ErrTxDone && err != nil{
        log.Fatalln(err)
    }
}

func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?parseTime=true")
    if err != nil {
        log.Fatalln(err)
    }

    defer db.Close()

    tx, err := db.Begin()   //tx对象，可以实现当出现异常时，数据库进行回滚
    if err != nil {
        log.Fatalln(err)
    }
    defer clearTransaction(tx)

    rs, err := tx.Exec("UPDATE user SET gold=50 WHERE real_name='vanyarpy'")
    if err != nil {
        log.Fatalln(err)
    }
    rowAffected, err := rs.RowsAffected()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(rowAffected)

    rs, err = tx.Exec("UPDATE user SET gold=150 WHERE real_name='noldorpy'")
    if err != nil {
        log.Fatalln(err)
    }
    rowAffected, err = rs.RowsAffected()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(rowAffected)
    doSomething()

    if err := tx.Commit(); err != nil {
        tx.Rollback() //此时处理错误，会忽略doSomthing的异常
        log.Fatalln(err)
    }
}