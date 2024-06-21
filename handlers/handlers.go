package handlers

import (
	"encoding/json"
	"net/http"
    "strconv"
	"github.com/kar5960/chat_service/database"
)

type Message struct {
	ID        int    `json:"id"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := database.DB.Prepare("INSERT INTO messages (sender, receiver, content) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(msg.Sender, msg.Receiver, msg.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := database.DB.Query("SELECT id, sender, receiver, content, timestamp FROM messages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Sender, &msg.Receiver, &msg.Content, &msg.Timestamp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
}
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing message ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	stmt, err := database.DB.Prepare("DELETE FROM messages WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}