package engine

import (
	"reflect"
	"testing"

	"github.com/sgrinan/runbeacon/internal/testutil"
)

func TestLoadCatalogValidDirectory(t *testing.T) {
	catalog, err := LoadCatalog("testdata/validYAML")
	testutil.NoError(t, err, "LoadCatalog()")

	if len(catalog.Runbooks) != 1 {
		t.Fatalf("LoadCatalog() runbook count = %d, want 1", len(catalog.Runbooks))
	}

	want := testutil.ValidRunbook()
	got := catalog.Runbooks[0]
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadCatalog() = %#v, want %#v", got, want)
	}
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

	if len(catalog.Runbooks) != 1 {
		t.Fatalf("LoadCatalog() runbook count = %d, want 1", len(catalog.Runbooks))
	}

	want := testutil.ValidRunbook()
	got := catalog.Runbooks[0]
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("LoadCatalog() = %#v, want %#v", got, want)
	}
}

func TestLoadCatalogLoadsMultipleRunbooks(t *testing.T) {
	catalog, err := LoadCatalog("testdata/multiple")
	testutil.NoError(t, err, "LoadCatalog()")

	if len(catalog.Runbooks) != 3 {
		t.Fatalf("LoadCatalog() runbook count = %d, want 3", len(catalog.Runbooks))
	}

	gotIDs := make(map[string]bool)

	for _, runbook := range catalog.Runbooks {
		gotIDs[runbook.ID] = true
	}

	wantIDs := []string{
		"kubernetes.pod.crashloopbackoff",
		"kubernetes.pod.pending",
		"kubernetes.pod.imagepullbackoff",
	}

	for _, id := range wantIDs {
		if !gotIDs[id] {
			t.Errorf("LoadCatalog() missing runbook %q", id)
		}
	}
}
