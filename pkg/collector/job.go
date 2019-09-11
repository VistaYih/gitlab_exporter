package collector // choerodon.io/gitlab_exporter/pkg/collector
import (
	"choerodon.io/gitlab-exporter/pkg/utils"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/xanzy/go-gitlab"
	"os"
)

type gitlabJobCollector struct {
	Runner          string
	Status          string
	Project         string
	GroupName       string
	JobDurationDesc *prometheus.Desc
}

func init() {
	registerCollector("gitlab_job", defaultEnabled, NewGitlabJobCollector)
}

func NewGitlabJobCollector() (Collector, error) {
	return &gitlabJobCollector{
		GroupName: "",
		Project:   "",
		Runner:    "",
		Status:    "",
		JobDurationDesc: prometheus.NewDesc(
			"gitlab_job_duration_time",
			"Duration time of utils job",
			[]string{"runner", "pipeline", "job", "status"},
			nil,
		),
	}, nil
}

func (g *gitlabJobCollector) Update(ch chan<- prometheus.Metric) error {
	if err := g.updateJobDurationTime(ch); err != nil {
		return err
	}
	return nil
}

func (g *gitlabJobCollector) ReallyGitlabJobDurationTime() (jobDurationTime []gitlab.Job) {
	client, err := utils.NewGitlabClient("./config/gitlab-config.yaml")
	if err != nil {
		fmt.Printf("gitlab get config failed: %s", err)
		os.Exit(1)
	}
	projects, _, err := client.Projects.ListProjects(&gitlab.ListProjectsOptions{})
	for _, project := range projects {
		log.Infof("New fetch project: %s", project.ID)
		jobs, _, err := client.Jobs.ListProjectJobs(project.ID, &gitlab.ListJobsOptions{})
		if err != nil {
			log.Warnf("Get job failed: %s", err)
		}
		for _, job := range jobs {
			log.Infof("New fetch job: %s", job.ID)
			jobDurationTime = append(jobDurationTime, job)
		}
	}

	return
}

func (g *gitlabJobCollector) updateJobDurationTime(ch chan<- prometheus.Metric) error {
	jobList := g.ReallyGitlabJobDurationTime()

	for _, job := range jobList {
		ch <- prometheus.MustNewConstMetric(
			g.JobDurationDesc,
			prometheus.GaugeValue,
			job.Duration,
			job.Runner.Name,
			string(job.Pipeline.ID),
			job.Name,
			job.Status,
		)
	}

	return nil
}
