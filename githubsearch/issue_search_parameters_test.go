package githubsearch

import (
	"testing"
	"time"
)

func TestIssueSearchParametersStringEmpty(t *testing.T) {
	empty := IssueSearchParameters{}
	actual := empty.String()
	if actual != "" {
		t.Fatalf("Expected no query string, but got: %q", actual)
	}
}

func TestIssueSearchParametersStringSomeFields(t *testing.T) {
	some := IssueSearchParameters{
		Repository: &RepositoryName{Owner: "defunkt", Name: "pjax"},
		State:      Unmerged,
		Type:       PullRequest,
	}
	actual := some.String()
	expected := "is:pr is:unmerged repo:defunkt/pjax"
	if actual != expected {
		t.Fatalf("Expected %q, but got: %q", expected, actual)
	}
}

func TestIssueSearchParametersStringEveryField(t *testing.T) {
	all := IssueSearchParameters{
		Type:           Issue,
		Scope:          BodyScope,
		Visibility:     Private,
		Author:         "defunkt1",
		Assignees:      []string{"defunkt2", "defunkt3"},
		Mentions:       "defunkt4",
		Commenter:      "defunkt5",
		Involves:       "defunkt6",
		Team:           "github/ceo",
		State:          Open,
		Labels:         []string{"good-first-issue", "bug fix"},
		Milestone:      "GitHub 1.0",
		ProjectBoard:   "github/dotcom/2",
		MissingField:   Milestone,
		Language:       "ruby",
		Status:         Pending,
		HeadBranchName: "search-api",
		BaseBranchName: "master",
		CreatedAt:      &TimeParameters{Time: time.Date(2007, 01, 02, 19, 52, 41, 0, time.UTC)},
		UpdatedAt:      &TimeParameters{Time: time.Date(2007, 02, 02, 19, 52, 41, 0, time.UTC), Modifier: GreaterThanOrEqualTo},
		MergedAt:       &TimeParameters{Time: time.Date(2007, 03, 02, 19, 52, 41, 0, time.UTC), Modifier: LessThan},
		ClosedAt: &TimeParameters{
			Time:     time.Date(2007, 04, 02, 19, 52, 41, 0, time.UTC),
			Modifier: Range,
			EndTime:  time.Date(2007, 05, 02, 19, 52, 41, 0, time.UTC),
		},
		NumComments:         &NumericalParameters{Count: 100, Modifier: GreaterThan},
		User:                "defunkt7",
		Organization:        "github",
		Repository:          &RepositoryName{Owner: "github", Name: "dotcom"},
		Reviewed:            ChangesRequestedReview,
		ReviewedBy:          "defunkt8",
		ReviewRequested:     "defunkt9",
		TeamReviewRequested: "github/ceo",
		Query:               "Nice Work!",
	}
	actual := all.String()
	expected := `is:issue in:body is:private author:defunkt1 assignee:defunkt2 assignee:defunkt3 mentions:defunkt4 commenter:defunkt5 involves:defunkt6 team:github/ceo is:open label:good-first-issue label:"bug fix" milestone:"GitHub 1.0" project:github/dotcom/2 no:milestone language:ruby status:pending head:search-api base:master created:2007-01-02T19:52:41+0000 updated:>=2007-02-02T19:52:41+0000 merged:<2007-03-02T19:52:41+0000 closed:2007-04-02T19:52:41+0000..2007-05-02T19:52:41+0000 comments:>100 user:defunkt7 org:github repo:github/dotcom review:changes_requested reviewed-by:defunkt8 review-requested:defunkt9 team-review-requested:github/ceo Nice Work!`
	if actual != expected {
		t.Fatalf("Expected %q, but got: %q", expected, actual)
	}
}
