package users

import (
	"encoding/json"
	"net/http"
	"users/internal/models"
	"users/internal/repositories/users"
	_ "users/internal/repositories/users"

	"github.com/sirupsen/logrus"
)

// Ajouteruser
// @Tags         users
// @Summary      Ajouter une nouvelle user.
// @Description  Endpoint pour ajouter une nouvelle user.
// @Accept       json
// @Produce      json
// @Param        user body models.user true "Nouvelle user à ajouter"
// @Success      201            {string} string
// @Failure      400            "Requête invalide"
// @Failure      500            "Erreur interne du serveur"
// @Router       /users [post]
func AddUser(w http.ResponseWriter, r *http.Request) {
	// Décodez le corps de la requête JSON dans une structure de données models.User
	var nouvelleUser models.User
	err := json.NewDecoder(r.Body).Decode(&nouvelleUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Requête invalide"))
		return
	}

	// Appeler le service pour ajouter la nouvelle user
	err = users.AddUser(&nouvelleUser)
	if err != nil {
		logrus.Errorf("Erreur lors de l'ajout de la user : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}

	// Répondre avec un statut 201 Created si tout s'est bien passé
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("user ajoutée avec succès"))
}
