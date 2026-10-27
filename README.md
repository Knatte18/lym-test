# lym-test

Fixture host repo for the Loomyard **weft** sandbox. Paired with `lym-test-weft`.

Carries a `services/api/` subpath used to exercise **weft path mirroring** — how the
weft worktree mirrors host subpaths. This is *not* an invitation to run `lym` from that
subdirectory: `lym` resolves against the current directory's own `_lym/` (it does not
walk up to a parent), so it is expected to work only from the directory where it was
initialized — this repo's root. Running `lym` from an uninitialized subdirectory
correctly reports `not initialized here; run "lym init"`; that is expected behaviour,
not a defect.
