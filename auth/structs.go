package auth

type Stage2AuthRequest struct {
	mac      string
	id       string
	auth     string
	sid      string
	version  string
	username string
	region   string
	language string
	country  string
	birthday string
	datetime string
	color    string
}
