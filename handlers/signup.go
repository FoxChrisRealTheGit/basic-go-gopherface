package handlers

import(
	"net/http"
	"tutorials/backendwebdev/gopherface/validationkit"
	"tutorials/backendwebdev/gopherface/gopherfacedb/common"
)

type SignUpForm struct{
	FieldNames []string
	Fields map[string]string
	Errors map[string]string
}

//DisplaySignupForm displays the Sign Up form
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm){
	RenderTemplate(w, "./templates/signupform.html", s)
}

func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignUpForm){
	RenderTemplate(w, "./templates.signconfirmation.html", s)
}

func PopulateFormFields(r *http.Request, s *SignUpForm){
	for _, fieldName := range s.FieldNames{
		s.Fields[fieldName] = r.FormValue(fieldName)
	}
}

// ValidateSignUpForm validates the Sign Up form's fields
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, e *common.Env){
	PopulateFormFields(r,s)
	//Check if username was filled out
	if r.FormValue("username") == ""{
		s.Errors["usernameError"] = "The username field is required."
	}

	//Check if first name was filled out
	if r.FormValue("firstName") == "" {
		s.Errors["lastNameError"] = "The e-mail address field is required."
	}

	// Check if e-mail address was filled out
	if r.FormValue("email") == ""{
		s.Errors["emailError"] = "The e-mail address field is required."
	}

// Check if password was filled out
if r.FormValue("password") == ""{
	s.Errors["passwordError"] = "The password field is required."
}

	// Check if password was filled out
	if r.FormValue("confirmPassword") == ""{
		s.Errors["confirmPasswordError"] = "The confirm password field is required."
	}

	// Check username syntax
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false{
		usernameErrorMessage := "The username entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok{
			s.Errors["usernameError"] += " " + usernameErrorMessage
		} else{
			s.Errors["usernameError"] = usernameErrorMessage
		}
	}

	// Check e-mail address syntax
	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false{
		emailErrorMessage := "The e-mail address entered has an improper syntax."
		if _, ok := s.Errors["usernameError"]; ok{
			s.Errors["emailError"] = emailErrorMessage
		}
	}

	// Check if password and confirm password field values match
	if r.FormValue("password") != r.FormValue("confirmPassword"){
		s.Errors["confirmPasswordError"] = "The password and confirm password fields do not match."
	}

	if len(s.Errors) > 0{
		DisplaySignUpForm(w, r, s)
	} else{
		ProcessSignUpForm(w, r, s, e)
	}
}

// ProcessSignupForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, e *common.Env){
	//this indicates that there was a successful form submission	
	//this needs to hook up to the bolt db
	u := models.NewUser(r.FormValue("username"), r.FormValue("firstname"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("password"))
	err := e.DB.CreateUser(u)
	if err != nil{
		log.Print(err)
	}
	user, err := e.DB.GetUser("username")
	if err != nil{
		log.Print(err)
	}else{
		fmt.Printf("Fetch User Result: %+v\n", user)
	}
	//Display form confirmation message
	DisplayConfirmation(w, r, s)
}

func SignUpHandler(e *common.Env) http.Handler{
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	s:= SignUpForm{}
	s.FieldNames = []string{"username", "firstName", "lastName", "email"}
	s.Fields = make(map[string]string)
	s.Errors = make(map[string]string)

	switch r.Method{
	case "GET":
		DisplaySignUpForm(w, r, &s)
	case "POST":
		ValidateSignUpForm(w, r, &s, e)
	default:
		DisplaySignUpForm(w, r, &s)
	}
})
}