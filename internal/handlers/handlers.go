package handlers

//type UserHandler struct {
//	controller.Users
//}
//
//func NewUserHandler(s controller.UserService) *UserHandler {
//	return &UserHandler{Users: controller.NewUsers(s)}
//}
//
//func (u UserHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
//	var newUser entity.User
//	err := json.NewDecoder(r.Body).Decode(&newUser)
//	if err != nil {
//		http.Error(w, "Invalid request payload", http.StatusBadRequest)
//		return
//	}
//	if newUser.Login == "" || newUser.Password == "" {
//		http.Error(w, "Login and password are required", http.StatusBadRequest)
//		return
//	}
//
//	err = u.Registration(context.TODO(), newUser)
//	if err != nil {
//		switch {
//		case strings.Contains(err.Error(), "user already exists"):
//			http.Error(w, "Login already exists", http.StatusConflict)
//		default:
//			http.Error(w, "Internal server error", http.StatusInternalServerError)
//		}
//		return
//	}
//
//	w.Write([]byte("User successfully registered and authenticated"))
//}
