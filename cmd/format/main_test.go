package main_test

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"regexp"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	"github.com/pivotal/monitoring-indicator-protocol/pkg/go_test"
)

func TestFormatBinary(t *testing.T) {
	g := NewGomegaWithT(t)

	binPath, err := go_test.Build("./", "-race")
	g.Expect(err).ToNot(HaveOccurred())

	t.Run("shows CLI version", func(t *testing.T) {
		g := NewGomegaWithT(t)

		binPath, err := go_test.Build("./", "-race", "-ldflags", "-X main.Version=1.1 -X main.OS=darwin")
		g.Expect(err).ToNot(HaveOccurred())
		cmd := exec.Command(binPath, "-version")

		buffer := bytes.NewBuffer(nil)
		session, _ := gexec.Start(cmd, buffer, nil)
		g.Eventually(session, 5).Should(gexec.Exit(0))

		output := buffer.String()
		g.Expect(output).To(ContainSubstring("cli version 1.1 darwin"))
	})

	t.Run("complains if no indicators file path specified", func(t *testing.T) {
		g := NewGomegaWithT(t)

		cmd := exec.Command(binPath)

		session, _ := gexec.Start(cmd, nil, nil)
		g.Eventually(session, 5).Should(gexec.Exit(1))
	})

	t.Run("complains when there is erb that did not get filled in", func(t *testing.T) {
		g := NewGomegaWithT(t)

		cmd := exec.Command(binPath,
			"-format", "html",
			"-indicators", "test_fixtures/erb-bad-doc.yml")

		session, err := gexec.Start(cmd, nil, nil)
		g.Expect(err).ToNot(HaveOccurred())

		g.Eventually(session, 5).Should(gexec.Exit(1))
		g.Expect(session.Err).To(gbytes.Say(regexp.QuoteMeta("found raw un-interpolated ERB, please check your input for ERB")))
	})

	t.Run("outputs formatted HTML", func(t *testing.T) {
		g := NewGomegaWithT(t)

		cmd := exec.Command(binPath,
			"-format", "html",
			"-metadata", "deployment=[deployment],source_id=[source_id]",
			"-indicators", "../../example_indicators.yml")

		buffer := bytes.NewBuffer(nil)

		session, err := gexec.Start(cmd, buffer, nil)
		g.Expect(err).ToNot(HaveOccurred())

		g.Eventually(session, 5).Should(gexec.Exit(0))

		html := buffer.String()

		t.Run("It displays document title and description", func(t *testing.T) {
			g := NewGomegaWithT(t)
			g.Expect(html).To(ContainSubstring(`<title>Monitoring Document Product</title>`))
			g.Expect(html).To(ContainSubstring(`<h1>Monitoring Document Product</h1>`))
			g.Expect(html).To(ContainSubstring(`Document description`))
		})

		t.Run("It displays indicator sections", func(t *testing.T) {
			g := NewGomegaWithT(t)
			g.Expect(html).To(ContainSubstring(`<h2><a id="indicators"></a>Indicators</h2>`))
			g.Expect(html).To(ContainSubstring(`This section includes indicators`))

			g.Expect(html).To(ContainSubstring(`<h3><a id="doc_performance_indicator"></a>Doc Performance Indicator</h3>`))

			g.Expect(html).To(ContainSubstring(`avg_over_time(demo_latency{source_id="[source_id]",deployment="[deployment]"}[5m])`))
		})

		t.Run("It does not have multiple % signs", func(t *testing.T) {
			g := NewGomegaWithT(t)

			g.Expect(html).ToNot(ContainSubstring("%%"))
		})
	})

	t.Run("outputs bookbinder formatted HTML", func(t *testing.T) {
		g := NewGomegaWithT(t)

		cmd := exec.Command(binPath,
			"-format", "bookbinder",
			"-metadata", "deployment=[deployment],source_id=[source_id]",
			"-indicators", "../../example_indicators.yml")

		buffer := bytes.NewBuffer(nil)

		session, err := gexec.Start(cmd, buffer, nil)
		g.Expect(err).ToNot(HaveOccurred())

		g.Eventually(session, 5).Should(gexec.Exit(0))

		html := buffer.String()

		t.Run("It displays document title and description", func(t *testing.T) {
			g := NewGomegaWithT(t)
			g.Expect(html).To(ContainSubstring(`title: Monitoring Document Product`))
			g.Expect(html).To(ContainSubstring(`Document description`))
		})

		t.Run("It displays indicator sections", func(t *testing.T) {
			g := NewGomegaWithT(t)
			g.Expect(html).To(ContainSubstring(`## <a id="indicators"></a>Indicators`))
			g.Expect(html).To(ContainSubstring(`This section includes indicators`))

			g.Expect(html).To(ContainSubstring(`### <a id="doc_performance_indicator"></a>Doc Performance Indicator`))

			g.Expect(html).To(ContainSubstring(`avg_over_time(demo_latency{source_id="[source_id]",deployment="[deployment]"}[5m])`))
		})

		t.Run("It does not have multiple % signs", func(t *testing.T) {
			g := NewGomegaWithT(t)

			g.Expect(html).ToNot(ContainSubstring("%%"))
		})
	})

	t.Run("outputs prometheus alert configuration", func(t *testing.T) {
		t.Run("with no metadata flag", func(t *testing.T) {
			g := NewGomegaWithT(t)

			cmd := exec.Command(binPath,
				"-format", "prometheus-alerts",
				"-indicators", "../../example_indicators.yml")

			buffer := bytes.NewBuffer(nil)

			session, err := gexec.Start(cmd, buffer, nil)
			g.Expect(err).ToNot(HaveOccurred())

			g.Eventually(session, 5).Should(gexec.Exit(0))

			prometheusAlertConfigurationYML := buffer.String()

			fileBytes, err := ioutil.ReadFile("test_fixtures/prometheus_alert.yml")
			g.Expect(err).ToNot(HaveOccurred())

			g.Expect(prometheusAlertConfigurationYML).To(MatchYAML(fileBytes))
		})

		t.Run("with metadata flag", func(t *testing.T) {
			g := NewGomegaWithT(t)

			cmd := exec.Command(binPath,
				"-format", "prometheus-alerts",
				"-metadata", "deployment=my-other-service-deployment",
				"-indicators", "../../example_indicators.yml")

			buffer := bytes.NewBuffer(nil)

			session, err := gexec.Start(cmd, buffer, nil)
			g.Expect(err).ToNot(HaveOccurred())

			g.Eventually(session, 5).Should(gexec.Exit(0))

			prometheusAlertConfigurationYML := buffer.String()

			g.Expect(prometheusAlertConfigurationYML).To(ContainSubstring(`deployment: my-other-service-deployment`))
		})
	})

	t.Run("outputs grafana dashboards", func(t *testing.T) {
		t.Run("with no metadata flag", func(t *testing.T) {
			g := NewGomegaWithT(t)

			cmd := exec.Command(binPath,
				"-format", "grafana",
				"-indicators", "../../example_indicators.yml")

			buffer := bytes.NewBuffer(nil)

			session, err := gexec.Start(cmd, buffer, nil)
			g.Expect(err).ToNot(HaveOccurred())

			g.Eventually(session, 5).Should(gexec.Exit(0))

			text := buffer.String()

			t.Run("it outputs indicators titles", func(t *testing.T) {
				g := NewGomegaWithT(t)
				g.Expect(text).To(ContainSubstring(`## title\nDoc Performance Indicator`))
				g.Expect(text).To(ContainSubstring(`"expr":"avg_over_time(demo_latency{source_id=\"my-metric-source\",deployment=\"my-service-deployment\"}[5m])"`))
			})
		})

		t.Run("with metadata flag", func(t *testing.T) {
			g := NewGomegaWithT(t)

			cmd := exec.Command(binPath,
				"-format", "grafana",
				"-metadata", "deployment=my-other-service-deployment",
				"-indicators", "../../example_indicators.yml")

			buffer := bytes.NewBuffer(nil)

			session, err := gexec.Start(cmd, buffer, nil)
			g.Expect(err).ToNot(HaveOccurred())

			g.Eventually(session, 5).Should(gexec.Exit(0))

			text := buffer.String()

			t.Run("it outputs indicators titles", func(t *testing.T) {
				g := NewGomegaWithT(t)
				g.Expect(text).To(ContainSubstring(`## title\nDoc Performance Indicator`))
				g.Expect(text).To(ContainSubstring(`"expr":"avg_over_time(demo_latency{source_id=\"my-metric-source\",deployment=\"my-other-service-deployment\"}[5m])"`))
			})
		})
	})
}
