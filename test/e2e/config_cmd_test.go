// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package e2e

import (
	"testing"

	"github.com/goharbor/harbor-cli/cmd/harbor/root"
	"github.com/goharbor/harbor-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func Test_ConfigCmd(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config"})
	err := rootCmd.Execute()
	assert.Nil(t, err)
}

func Test_ConfigListCmd(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "list"})
	err := rootCmd.Execute()
	assert.Nil(t, err)
}

func Test_ConfigGetCmd_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "get", "credentials.serveraddress"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
}

func Test_ConfigGetCmd_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "get", "serveraddress"})
	err = rootCmd.Execute()
	assert.Error(t, err, "Expected an error when getting a non-existent config item")
}

func Test_ConfigGetCmd_CredentialName_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "get", "credentials.serveraddress", "--name", "harbor-cli@http://demo.goharbor.io"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
}

func Test_ConfigGetCmd_CredentialName_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "get", "credentials.serveraddress", "--name", "harbor-cli@http://goharbor.io"})
	err = rootCmd.Execute()
	assert.Error(t, err, "Expected an error when getting a non-existent credential name")
}

func Test_ConfigUpdateCmd_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "update", "credentials.serveraddress", "http://demo.goharbor.io"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
}

func Test_ConfigUpdateCmd_CredentialName_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "update", "credentials.serveraddress", "http://demo.goharbor.io", "--name", "harbor-cli@http://demo.goharbor.io"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
}

func Test_ConfigUpdateCmd_CredentialName_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "update", "credentials.serveraddress", "http://demo.goharbor.io", "--name", "harbor-cli@http://goharbor.io"})
	err = rootCmd.Execute()
	assert.Error(t, err, "Expected an error when setting a non-existent credential name")
}

func Test_ConfigUpdateCmd_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "update", "serveraddress", "http://demo.goharbor.io"})
	err = rootCmd.Execute()
	assert.Error(t, err, "Expected an error when setting a non-existent config item")
}

func Test_ConfigDeleteCmd_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "delete", "credentials.serveraddress"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
	config, err := utils.GetCurrentHarborConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert.Empty(t, config.Credentials[0].ServerAddress)
}

func Test_ConfigDeleteCmd_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "delete", "serveraddress"})
	err = rootCmd.Execute()
	assert.Error(t, err, "Expected an error when deleting a non-existent config item")
}

func Test_ConfigDeleteCmd_CredentialName_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "delete", "credentials.serveraddress", "--name", "harbor-cli@http://demo.goharbor.io"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
	config, err := utils.GetCurrentHarborConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert.Empty(t, config.Credentials[0].ServerAddress)
}

func Test_ConfigDeleteCmd_CredentialName_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "delete", "credentials.serveraddress", "--name", "harbor-cli@http://goharbor.io"})
	err = rootCmd.Execute()
	assert.Error(t, err, "Expected an error when deleting a non-existent credential name")
}

func Test_ConfigDeleteCmd_Current_Flag_Success(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	testConfig := &utils.HarborConfig{
		CurrentCredentialName: "harbor-cli@http://demo.goharbor.io",
		Credentials: []utils.Credential{
			{
				Name:          "harbor-cli@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "harbor-cli",
				Password:      "Harbor12345",
			},
			{
				Name:          "admin@http://demo.goharbor.io",
				ServerAddress: "http://demo.goharbor.io",
				Username:      "admin",
				Password:      "Admin12345",
			},
		},
	}
	err := utils.UpdateConfigFile(testConfig)
	if err != nil {
		t.Fatal(err)
	}
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "delete", "--current"})
	err = rootCmd.Execute()
	assert.NoError(t, err)
	config, err := utils.GetCurrentHarborConfig()
	if err != nil {
		t.Fatal(err)
	}
	assert.Empty(t, config.CurrentCredentialName)
	assert.NotEmpty(t, config.Credentials)
	assert.NoError(t, err)
}

func Test_ConfigDeleteCmd_Current_Flag_With_Item_Failure(t *testing.T) {
	tempDir := t.TempDir()
	data := Initialize(t, tempDir)
	defer ConfigCleanup(t, data)
	SetMockKeyring(t)
	rootCmd := root.RootCmd()
	rootCmd.SetArgs([]string{"config", "delete", "credentials.serveraddress", "--current"})
	err := rootCmd.Execute()
	assert.Error(t, err, "Expected an error when specifying both --current and an item")
}
