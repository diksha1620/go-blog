package helper

// // func Respond(w http.ResponseWriter, data map[string]interface{}) {
// // 	w.Header().Add("Content-Type", "application/json")
// // 	json.NewEncoder(w).Encode(data)
// // }

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}
