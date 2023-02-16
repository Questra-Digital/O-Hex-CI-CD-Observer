package main

import "time"

type commit struct {
	Author_email string `json:"author_email"`
	Author_name  string `json:"author_name"`
	Created_at   string `json:"created_at"`
	ID           string `json:"id"`
	Message      string `json:"message"`
	Short_id     string `json:"short_id"`
	Title        string `json:"title"`
}

type time_details struct {
	Coverage        string    `json:"coverage"`
	Allow_feature   bool      `json:"allow_feature"`
	Created_at      time.Time `json:"created_at"`
	Started_at      time.Time `json:"started_at"`
	Finshed_at      time.Time `json:"finished_at"`
	Erased_at       time.Time `json:"erased_at"`
	Duration        int       `json:"duration"`
	Queued_duration int       `json:"queued_duration"`
}

type artifacts_main struct {
	Filename string `json:"filename"`
	Size     int    `json:"size"`
}

type artifact_solo struct {
	File_type     string `json:"file_type"`
	Size_artifact int    `json:"size"`
	Filename      string `json:"filename"`
	File_foramt   string `json:"file_foramt"`
}

type ext struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type pipeline struct {
	ID         int    `json:"id"`
	Project_id int    `json:"project_id"`
	Ref        string `json:"ref"`
	Sha        string `json:"sha"`
	Status     string `json:"status"`
}

type user struct {
	Userid       int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	State        string `json:"state"`
	Avatar_url   string `json:"avatar_url"`
	Web_url      string `json:"web_url"`
	Created_at   string `json:"created_at"`
	Bio          string `json:"bio"`
	Location     string `json:"location"`
	Public_email string `json:"public_email"`
	Skype        string `json:"skype"`
	Linkedin     string `json:"linkedin"`
	Twitter      string `json:"twitter"`
	Website_url  string `json:"website_url"`
	Organization string `json:"organization"`
}
