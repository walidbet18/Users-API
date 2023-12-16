package users

import (
	"net/http"
	"users/internal/services/users"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// SupprimerUser est un contrôleur pour supprimer un utilisateur.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'utilisateur à supprimer depuis les paramètres de la requête
	ctx := r.Context()
	userID, ok := ctx.Value("userId").(uuid.UUID)
	if !ok || userID == uuid.Nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("ID utilisateur manquant ou invalide"))
		return
	}

	// Appeler la fonction DeleteUser du repository pour supprimer l'utilisateur
	err := users.DeleteUser(userID)
	if err != nil {
		logrus.Errorf("Erreur lors de la suppression de l'utilisateur : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}

	// Répondre avec un statut 200 OK si tout s'est bien passé
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Utilisateur supprimé avec succès"))
}
