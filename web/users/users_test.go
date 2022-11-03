package users

import (
	"api/pkg/user"
	"bytes"
	"encoding/json"
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
	u := &user.User{
		Id: 1,
	}
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/user/%d", u.Id))
	if err != nil {
		t.Error(err)
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(bb, u)
	if err != nil {
		t.Error(err)
	}

	mock := `{"first_name": "yernur", "last_name": "abishev", "interests": "coding,golang,ubuntu"}`
	if u.Data != mock {
		t.Error(err)
	}
}

// func Test_Delete(t *testing.T) {
// 	u := &user.User{
// 		Id: 1,
// 	}
// 	req, err := http.NewRequest("DELETE", (fmt.Sprintf("http://localhost:8080/user/%d", u.Id)), nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if resp.StatusCode != http.StatusNoContent {
// 		t.Error("error on delete")
// 	}

// 	// bb, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	t.Error(err)
// 	// }
// 	// err = json.Unmarshal(bb, u)
// 	// if err != nil {
// 	// 	t.Error(err)
// 	// }

// 	// if resp.StatusCode == http.StatusOK {
// 	// 	t.Error("error on finding user")
// 	// }
// }

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
