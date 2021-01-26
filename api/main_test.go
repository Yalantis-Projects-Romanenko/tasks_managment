package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fdistorted/task_managment/models"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	user1 = "user133"
	user2 = "user244"
	host  = "localhost:5000"
)

func GenerateToken(username string) string {
	return "Bearer " + b64.StdEncoding.EncodeToString([]byte(username))
}

func buildLink(path string) string {
	return "http://" + host + path
}

func createRequest(method, path, token string, body interface{}) (req *http.Request) {
	var err error
	if body != nil {
		reqBody, err := json.Marshal(body)
		if err != nil {
			fail()
		}
		buffer := bytes.NewBuffer(reqBody)
		req, err = http.NewRequest(method, buildLink(path), buffer)
	} else {
		req, err = http.NewRequest(method, buildLink(path), nil)
	}
	if err != nil {
		fail()
	}
	req.Header.Set("Authorization", token) // add fake token
	return req
}

func TestAuthorization(t *testing.T) {
	fmt.Println(GenerateToken(user1))
	fmt.Println(GenerateToken(user2))
	//assert(Equal(mtx.SpiralTransform(matrix), slice))
	client := &http.Client{}

	req := createRequest(http.MethodGet, "/projects/", "wrong token", nil)
	resp, err := client.Do(req)
	if err != nil {
		fail()
	}
	assert(resp.StatusCode == 401)

	req = createRequest(http.MethodGet, "/projects/", GenerateToken(user1), nil)
	resp, err = client.Do(req)
	if err != nil {
		fail()
	}
	fmt.Println(resp.StatusCode)
	assert(resp.StatusCode == 200)
}

func TestCreateUpdateProject(t *testing.T) {
	//assert(Equal(mtx.SpiralTransform(matrix), slice))
	client := &http.Client{}

	projectName := "test"
	projectDescription := "test description"
	updatedProjcectName := "updated name"
	updtedProjectDescription := "updated test description"

	// create project
	project := models.Project{
		Description: projectDescription,
		Name:        projectName,
	}
	req := createRequest(http.MethodPost, "/projects/", GenerateToken(user1), &project)
	resp, err := client.Do(req)
	if err != nil {
		fail()
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200)

	var respProject *models.Project
	err = json.NewDecoder(resp.Body).Decode(&respProject)
	if err != nil {
		fail()
	}

	assert(respProject.Name == projectName)
	assert(respProject.Description == projectDescription)

	// update project
	project = models.Project{
		Name:        updatedProjcectName,
		Description: updtedProjectDescription,
	}
	req = createRequest(http.MethodPut, "/projects/"+respProject.Id+"/", GenerateToken(user1), &project)
	resp, err = client.Do(req)
	if err != nil {
		fail()
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200)

	// get updated project
	req = createRequest(http.MethodGet, "/projects/"+respProject.Id+"/", GenerateToken(user1), nil)
	resp, err = client.Do(req)
	if err != nil {
		fail()
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 200)

	respProject = nil
	err = json.NewDecoder(resp.Body).Decode(&respProject)
	if err != nil {
		fail()
	}

	assert(respProject.Name == updatedProjcectName)
	assert(respProject.Description == updtedProjectDescription)

	// get project with other user credentials
	req = createRequest(http.MethodGet, "/projects/"+respProject.Id+"/", GenerateToken(user2), nil)
	resp, err = client.Do(req)
	if err != nil {
		fail()
	}
	defer resp.Body.Close()
	assert(resp.StatusCode == 404)
}

func fail() {
	fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
	os.Exit(1)
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
		os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
