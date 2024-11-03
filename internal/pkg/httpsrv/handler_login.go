package httpsrv

import (
	"html/template"
	"log"
	"net/http"
)

var users = map[string]string{"test":"test"}

func (s *Server) handlerLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("into handler", r.Method)
	if r.Method == "GET" {
		template.Must(template.New("").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Login Page</title>
			<head>
			<body>
				<div class="login-container">
					<h2>Login</h2>
					<form action="/goapp/login" method="POST">
						<label for="username">Username:</label>
						<input type="text" id="username" name="username" required><br>
						<label for="password">Password:</label>
						<input type="password" id="password" name="password" required><br>
						<button type="submit">Login</button>
					</form>
				</div>
			</body>
			</html>
		`)).Execute(w, nil)
	} else if r.Method == "POST" {
			
		// ParseForm parses the raw query from the URL and updates r.Form
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}

		// Get username and password from the parsed form
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		// Check if user exists
		storedPassword, exists := users[username]
		if exists {
			// It returns a new session if the sessions doesn't exist
			session, _ := s.store.Get(r, "session.id")
			if storedPassword == password {
				session.Values["authenticated"] = true
				// Saves all sessions used during the current request
				session.Save(r, w)
			} else {
				http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			}
			w.Write([]byte("Login successfully!"))
		}
	} else {
		http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
		return
	}
}