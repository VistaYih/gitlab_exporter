package main

import (
	"choerodon.io/gitlab-exporter/pkg/utils"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"os"
)

func main() {
	client, err := utils.NewGitlabClient("./config/gitlab-config.yaml")
	if err != nil {
		fmt.Printf("gitlab get config failed: %s\n", err)
		os.Exit(1)
	}

	listOptions := gitlab.ListOptions{
		PerPage: 20,
		Page:    6,
	}
	projectOptions := gitlab.ListProjectsOptions{
		ListOptions: listOptions,
	}
	/*jobOptions := gitlab.ListJobsOptions{
		ListOptions: listOptions,
	}*/

	for {
		fmt.Printf("GET Projects page %d / pageSize %d\n", projectOptions.ListOptions.Page, projectOptions.ListOptions.PerPage)
		projects, resp, err := client.Projects.ListProjects(&projectOptions)
		fmt.Printf("Projects: %s\n", projects)
		if err != nil {
			fmt.Printf("\nGet project failed: %s\n", err)
			return
		}
		/*for _, project := range projects {
			fmt.Printf("\nNow fetch project: %s\n", project.ID)
			jobs, _, err := client.Jobs.ListProjectJobs(project.ID, &jobOptions)
			if err != nil {
				fmt.Printf("Get job failed: %s", err)
			}
			for _, job := range jobs {
				fmt.Printf("\tNow fetch job: %d--%s\n", job.ID, job.Name)
			}
		}*/
		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		projectOptions.Page = resp.NextPage
	}

}
