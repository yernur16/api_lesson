package user

import (
	"testing"
)

func Test_CRUD(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	log.Fatalf("an error occured '%s' expected when opening db connection", err)
	// }

	// testDB := sqlx.NewDb(db, "sqlmock")
	// d.DB = testDB

	// u := &User{
	// 	Data: Data{
	// 		First_name: "yernur_testcrud",
	// 		Last_name:  "abishev_testcrud",
	// 		Interests:  "coding,pushups,etc",
	// 	},
	// }
	// bytes, err := json.Marshal(u)
	// if err != nil {
	// 	t.Error(err)
	// }
	// // mock.ExpectExec(`INSERT INTO users`).WithArgs([]byte(bytes)).WillReturnResult(sqlmock.NewResult(1, 1))

	// err = u.CreateUser()
	// if err != nil {
	// 	t.Errorf("error with CreateUser method: %s", err)
	// }
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("error with expectations: %s", err)

	// }
}
