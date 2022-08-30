package main

import "time"

type SnykResponse struct {
	Issues []struct {
		ID            string   `json:"id"`
		IssueType     string   `json:"issueType"`
		PkgName       string   `json:"pkgName"`
		PkgVersions   []string `json:"pkgVersions"`
		PriorityScore int      `json:"priorityScore"`
		Priority      struct {
			Score   int `json:"score"`
			Factors []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"factors"`
		} `json:"priority"`
		IssueData struct {
			ID          string `json:"id"`
			Title       string `json:"title"`
			Severity    string `json:"severity"`
			URL         string `json:"url"`
			Identifiers struct {
				Cve         []string      `json:"CVE"`
				Cwe         []string      `json:"CWE"`
				Alternative []interface{} `json:"ALTERNATIVE"`
			} `json:"identifiers"`
			Credit          []string `json:"credit"`
			ExploitMaturity string   `json:"exploitMaturity"`
			Semver          struct {
				Vulnerable []string `json:"vulnerable"`
			} `json:"semver"`
			PublicationTime       time.Time     `json:"publicationTime"`
			DisclosureTime        time.Time     `json:"disclosureTime"`
			CVSSv3                string        `json:"CVSSv3"`
			CvssScore             float64       `json:"cvssScore"`
			Language              string        `json:"language"`
			Patches               []interface{} `json:"patches"`
			NearestFixedInVersion string        `json:"nearestFixedInVersion"`
			IsMaliciousPackage    bool          `json:"isMaliciousPackage"`
		} `json:"issueData"`
		IsPatched bool `json:"isPatched"`
		IsIgnored bool `json:"isIgnored"`
		FixInfo   struct {
			IsUpgradable          bool   `json:"isUpgradable"`
			IsPinnable            bool   `json:"isPinnable"`
			IsPatchable           bool   `json:"isPatchable"`
			IsFixable             bool   `json:"isFixable"`
			IsPartiallyFixable    bool   `json:"isPartiallyFixable"`
			NearestFixedInVersion string `json:"nearestFixedInVersion"`
		} `json:"fixInfo"`
		Links struct {
			Paths string `json:"paths"`
		} `json:"links"`
	} `json:"issues"`
}
