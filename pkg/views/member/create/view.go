package create

import (
	"errors"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

type MemberUser struct {
	UserID   int
	Username string
}

type MemberGroup struct {
	ID          int
	GroupName   string
	GroupType   int
	LdapGroupDN string
}

type CreateView struct {
	ProjectNameOrID string
	RoleID          int
	MemberUser      *models.UserEntity
	MemberGroup     *models.UserGroup
}

func CreateMemberView(createView *CreateView) {
	theme := huh.ThemeCharm()
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Username").
				Value(&createView.MemberUser.Username).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("email cannot be empty")
					}
					return nil
				}),
		),
	).WithTheme(theme).Run()
	if err != nil {
		log.Fatal(err)
	}
}
