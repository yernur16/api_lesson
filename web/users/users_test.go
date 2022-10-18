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
		Username: "test_yernur",
		Password: "test_chief",
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

	if u.Username != "yernur" && u.Username != "chef" && u.Username != "rahat" && u.Username != "oleg" && u.Username != "Asya" {
		t.Error(err)
	}

}

func Test_Delete(t *testing.T) {
	u := &user.User{
		Id: 22,
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
