package users

import (
	"encoding/json"
	"net/http"
	"users/internal/models"
	"users/internal/services/users"

	"github.com/sirupsen/logrus"
)

// ModifierUser
// @Tags         users
// @Summary      Modifier un utilisateur existant.
// @Description  Endpoint pour modifier un utilisateur existant.
// @Accept       json
// @Produce      json
// @Param        user body models.User true "Utilisateur à modifier"
// @Success      200            {string} string
// @Failure      400            "Requête invalide"
// @Failure      500            "Erreur interne du serveur"
// @Router       /users/{id} [put]

// EditUser modifie les détails d'un utilisateur existant.
func EditUser(w http.ResponseWriter, r *http.Request) {
	// Décodez le corps de la requête JSON dans une structure de données models.User

	var updatedUser models.User                         // Pas de pointeur ici, juste la structure User
	err := json.NewDecoder(r.Body).Decode(&updatedUser) // Utilisation de "&" pour obtenir l'adresse de la structure
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Requête invalide"))
		return
	}

	// Appeler la fonction de mise à jour de l'utilisateur dans le référentiel approprié
	_, err = users.UpdateUser(&updatedUser) // Passer la référence de la structure mise à jour
	if err != nil {
		logrus.Errorf("Erreur lors de la modification de l'utilisateur : %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Erreur interne du serveur"))
		return
	}

	// Répondre avec un statut 200 OK si tout s'est bien passé
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Utilisateur mis à jour avec succès"))
}
