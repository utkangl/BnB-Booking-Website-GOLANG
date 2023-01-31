package models

// new struct to hold the datas that will be sent to templates by handler functions
// making the struct hold the possible data types that we may send
type TemplateData struct {
	Stringmap  map[string]string
	İntegerMap map[string]int
	FloatMap   map[string]float64
	Data       map[string]interface{}
	CSRFToken  string // Csrf security token (cross site request forgery)
	Flash      string // Messages to send to user, like "Successfully logged in"
	Warning    string // Warning to send to user
	Error      string // Error to send to user
}
