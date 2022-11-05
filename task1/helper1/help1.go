package helper1

type Response struct{
	SlackUsername string `json:"slackUsername"`
	Backend bool `json:"backend"`
	Age int `json:"age"`
	Bio string `json:"bio"`
}