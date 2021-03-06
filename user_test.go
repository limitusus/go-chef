package chef

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"testing"
)

var (
	testUserJSON = "test/user.json"
	testVerboseUserJSON = "test/verbose_user.json"
)

func TestUserFromJSONDecoder(t *testing.T) {
	if file, err := os.Open(testUserJSON); err != nil {
		t.Error("Unexpected error '", err, "' during os.Open on", testUserJSON)
	} else {
		dec := json.NewDecoder(file)
		var g User
		if err := dec.Decode(&g); err == io.EOF {
			log.Fatal(g)
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func TestVerboseUserFromJSONDecoder(t *testing.T) {
	if file, err := os.Open(testVerboseUserJSON); err != nil {
		t.Error("Unexpected error '", err, "' during os.Open on", testVerboseUserJSON)
	} else {
		dec := json.NewDecoder(file)
		var g User
		if err := dec.Decode(&g); err == io.EOF {
			log.Fatal(g)
		} else if err != nil {
			log.Fatal(err)
		}
	}
}

func TestUserslist(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "GET":
			fmt.Fprintf(w, `{ "user_name1": "https://url/for/user_name1", "user_name2": "https://url/for/user_name2"}`)
		}
	})

	// Test list
	users, err := client.Users.List()
	if err != nil {
		t.Errorf("Users.List returned error: %v", err)
	}
	listWant := map[string]string{"user_name1": "https://url/for/user_name1", "user_name2": "https://url/for/user_name2"}
	if !reflect.DeepEqual(users, listWant) {
		t.Errorf("Users.List returned %+v, want %+v", users, listWant)
	}
}

func TestVerboseUserslist(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "GET":
			fmt.Fprintf(w, `{
                                "janechef": { "email": "jane.chef@user.com", "first_name": "jane", "last_name": "chef_user" },
                                "yaelsmith": { "email": "yael.chef@user.com", "first_name": "yael", "last_name": "smith" }
                        }`)

		}
	})

	// Test list
	users, err := client.Users.VerboseList()
	if err != nil {
		t.Errorf("Verbose Users.List returned error: %v", err)
	}
	jane :=  UserVerboseResult{ Email: "jane.chef@user.com", FirstName: "jane", LastName: "chef_user" }
        yael :=  UserVerboseResult{ Email: "yael.chef@user.com", FirstName: "yael", LastName: "smith" }
	listWant := map[string]UserVerboseResult{ "janechef": jane, "yaelsmith": yael }
	if !reflect.DeepEqual(users, listWant) {
		t.Errorf("Verbose Users.List returned %+v, want %+v", users, listWant)
	}
}

func TestUserCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "POST":
			fmt.Fprintf(w, `{
                                "uri": "https://url/for/user_name1",
                                "private_key": "-----BEGIN RSA PRIVATE KEY-----"
                         }`)
		}
	})

	// Create User
	user := User{UserName: "user_name1", Email: "user_name1@mail.com", Password: "dummypass"}
	userresult, err := client.Users.Create(user)
	if err != nil {
		t.Errorf("Users.Create returned error: %v", err)
	}
	Want := UserResult{Uri: "https://url/for/user_name1", PrivateKey: "-----BEGIN RSA PRIVATE KEY-----"}
	if !reflect.DeepEqual(userresult, Want) {
		t.Errorf("Users.Create returned %+v, want %+v", userresult, Want)
	}
}

func TestUserGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/user1", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "GET":
			fmt.Fprintf(w, `{
                                "username": "user1",
                                "display_name": "User Name",
                                "email": "user1@mail.com",
                                "external_authentication_uid": "user1",
                                "first_name": "User",
                                "full_name": "User S Name",
                                "last_name": "Name",
                                "middle_name": "S",
                                "public_key": "-----BEGIN RSA PUBLIC KEY-----",
                                "recovery_authentication_enabled": true
                         }`)
		}
	})

	// Get User
	user, err := client.Users.Get("user1")
	if err != nil {
		t.Errorf("User.Get returned error: %v", err)
	}
	Want := User{UserName: "user1", DisplayName: "User Name", Email: "user1@mail.com", ExternalAuthenticationUid: "user1", FirstName: "User", FullName: "User S Name", LastName: "Name", MiddleName: "S", PublicKey: "-----BEGIN RSA PUBLIC KEY-----", RecoveryAuthenticationEnabled: true}
	if !reflect.DeepEqual(user, Want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, Want)
	}
}

func TestUserDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/user1", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == "DELETE":
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{
                        }`)
		}
	})

	// DELETE User
	err := client.Users.Delete("user1")
	if err != nil {
		t.Errorf("User.Get returned error: %v", err)
	}
}
