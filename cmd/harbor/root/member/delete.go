package member

import (
	"context"
	"strconv"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/member"
	"github.com/goharbor/harbor-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Deletes the member of the given project and Member
func DeleteMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete project by name or id",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			projectName := args[0]
			var memberID []int64

			for i, mid := range args {
				if i != 0 {
					mID, _ := strconv.Atoi(mid)
					memberID = append(memberID, int64(mID))
				}
			}

			if len(args) > 1 {
				err = runDeleteMember(args[0], memberID)
			} else {
				projectName = utils.GetProjectNameFromUser()
				memberID = utils.GetMemberIDFromUser()
				err = runDeleteMember(projectName, memberID)
			}
			if err != nil {
				log.Errorf("failed to delete project: %v", err)
			}
		},
	}

	return cmd
}

func runDeleteMember(projectName, memberID string) error {
	credentialName := viper.GetString("current-credential-name")
	client := utils.GetClientByCredentialName(credentialName)
	ctx := context.Background()
	_, err := client.Member.DeleteProjectMember(
		ctx,
		&member.DeleteProjectMemberParams{ProjectNameOrID: projectName, Mid: memberID},
	)
	if err != nil {
		return err
	}

	log.Info("Member deleted successfully")
	return nil
}
