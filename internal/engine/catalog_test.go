package engine

import (
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestLoadCatalogValidDirectory(t *testing.T) {
	catalog, err := LoadCatalog("testdata/validYAML")
	testutil.NoError(t, err, "LoadCatalog()")

	testutil.Equal(t, len(catalog.Runbooks), 1, "LoadCatalog() runbook count")

	want := testutil.CrashLoopBackOffRunbook()
	got := catalog.Runbooks[0]

	testutil.DeepEqual(t, got, want, "LoadCatalog()")
}

func TestLoadCatalogMissingDirectory(t *testing.T) {
	_, err := LoadCatalog("testdata/missing")
	testutil.ErrorContains(t, err, "LoadCatalog()", "read runbooks directory")
}

func TestLoadCatalogInvalidYAML(t *testing.T) {
	_, err := LoadCatalog("testdata/invalidYAML")
	testutil.ErrorContains(t, err, "LoadCatalog()", "parse runbook YAML")
}

func TestLoadCatalogInvalidRunbook(t *testing.T) {
	_, err := LoadCatalog("testdata/invalidRunbook")
	testutil.ErrorContains(t, err, "LoadCatalog()", "runbook id is required")
}

func TestLoadCatalogIgnoresNonRunbookFiles(t *testing.T) {
	catalog, err := LoadCatalog("testdata/mixed")
	testutil.NoError(t, err, "LoadCatalog()")

	testutil.Equal(t, len(catalog.Runbooks), 1, "LoadCatalog() runbook count")

	want := testutil.CrashLoopBackOffRunbook()
	got := catalog.Runbooks[0]

	testutil.DeepEqual(t, got, want, "LoadCatalog()")
}

func TestLoadCatalogLoadsMultipleRunbooks(t *testing.T) {
	catalog, err := LoadCatalog("testdata/multiple")
	testutil.NoError(t, err, "LoadCatalog()")

	testutil.Equal(t, len(catalog.Runbooks), 3, "LoadCatalog() runbook count")

	gotIDs := make(map[string]struct{})

	for _, runbook := range catalog.Runbooks {
		gotIDs[runbook.ID] = struct{}{}
	}

	wantIDs := []string{
		"kubernetes.pod.crashloopbackoff",
		"kubernetes.pod.pending",
		"kubernetes.pod.imagepullbackoff",
	}

	for _, id := range wantIDs {
		if _, ok := gotIDs[id]; !ok {
			t.Errorf("LoadCatalog() missing runbook %q", id)
		}
	}
}
