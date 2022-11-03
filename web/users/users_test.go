package users

import (
	"api/pkg/user"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_CreateUserHandler(t *testing.T) {
	u := user.User{
		Data: user.Data{
			First_name: "eren",
			Last_name:  "yeager",
			Interests:  "kill,freedom,rumbling",
		},
	}
	bb, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/user/3", "application/json", bytes.NewBuffer(bb))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusCreated {
		log.Fatal(err)
	}
}

func Test_GetUserHandler(t *testing.T) {
	// version with independence from user.Data structure
	// mock := struct {
	// 	First_name string `json:"first_name"`
	// 	Last_name  string `json:"last_name"`
	// 	Interests  string `json:"interests"`
	// }{
	// 	"Beka",
	// 	"Teka",
	// 	"1234",
	// }

	u := &user.User{
		Id: 9,
	}
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/user/%d", u.Id))
	if err != nil {
		t.Error(err)
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(bb, &u.Data)
	if err != nil {
		t.Error(err)
	}

	test := user.Data{First_name: "wolf", Last_name: "cat", Interests: "starwars,startrek"}

	// alternative test
	// if mock.First_name != u.Data.First_name {
	// 	t.Fatal("123423")
	// }

	if u.Data != test {
		t.Error(errors.New("error with Test_GetUserHandler"))
	}
}

func Test_Delete(t *testing.T) {
	u := &user.User{
		Id: 10,
	}
	req, err := http.NewRequest("DELETE", (fmt.Sprintf("http://localhost:8080/user/%d", u.Id)), nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Error("error on delete")
	}

	// bb, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Error(err)
	// }
	// err = json.Unmarshal(bb, u)
	// if err != nil {
	// 	t.Error(err)
	// }

	// if resp.StatusCode == http.StatusOK {
	// 	t.Error("error on finding user")
	// }
}

// func Test_Update(t *testing.T) {
// 	u := &user.User{
// 		Id:   1,
// 		Data: "test_mars",
// 	}
// 	bb, err := json.Marshal(u)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	req, err := http.NewRequest("PUT", (fmt.Sprintf("http://localhost:8080/user/%d", u.Id)), bytes.NewBuffer(bb))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		t.Error("error on update")
// 	}
// }
