## Go Simhash Experiment

### To Build

1. Use Go v1.11 or better. Make sure you have `GO111MODULE` set to `on` in your build environment, since this experiment uses Go modules.
2. Check out this repo into whatever directory, outside your `GOPATH`.
3. Inside the repo, run `go build ./...`

### To Run

First, be sure you have a tunnel open to the ContentID database. You will need the PGS AWS Salt Test SSH key.

```bash
ssh -i PATH_TO_PGS_SALT_TEST_KEY -L 54321:cluster01-content-id.cluster-c4uwuietvovl.us-east-1.rds.amazonaws.com:5432 ec2-user@cid-jump.pgs.io -N -f
```

Test your tunnel by connecting with `psql`:

```bash
psql -h localhost -p 54321 -U pgs -d content_id -W
```
That should ask you for the password. If you don't have it, ask EKINGERY or MSM.

Export the following variables into your environment:

```bash
export SHE_dbhost=localhost
export SHE_dbpass=PUT_THE_DB_PASSWORD_HERE
export SHE_dbname=content_id
export SHE_dbport=54321
export SHE_dbuser=pgs
```

