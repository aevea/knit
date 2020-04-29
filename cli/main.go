package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"golang.org/x/oauth2"

	"github.com/aevea/merge-master/internal/github"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "merge-master",
		Short: "TODO",
		Long:  "TODO",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("There is no root command. Please check merge-master --help.")
			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	oldestPRCmd := &cobra.Command{
		Use:   "pr-info",
		Short: "Gets the name and date of the longest open PR",
		RunE: func(cmd *cobra.Command, args []string) error {
			repository := cmd.Flag("repository").Value.String()
			token := cmd.Flag("token").Value.String()

			if repository == "" {
				return errors.New("missing repository")
			}

			if token == "" {
				return errors.New("missing token")
			}

			src := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			httpClient := oauth2.NewClient(context.Background(), src)

			repoName := strings.Split(repository, "/")

			githubClient := github.NewGithubClient(httpClient, repoName[0], repoName[1])

			oldestPR, err := githubClient.OldestPR()

			if err != nil {
				return err
			}

			t := table.NewWriter()

			t.AppendRow(
				table.Row{
					"Longest open PR",
					fmt.Sprintf("%.0f days", oldestPR.OpenFor.Hours()/24),
					oldestPR.URL,
				},
			)

			fmt.Println(t.Render())

			return nil
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	oldestPRCmd.PersistentFlags().String("repository", "", "repository in the format of owner/repository")
	oldestPRCmd.PersistentFlags().String("token", "", "token for github API")

	rootCmd.AddCommand(oldestPRCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
