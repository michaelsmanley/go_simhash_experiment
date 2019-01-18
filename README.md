## Go Simhash Experiment

### To Build

1. Use Go v1.11 or better. Make sure you have `GO111MODULE` set to `on` in your build environment, since this experiment uses Go modules.
2. Check out this repo into whatever directory, outside your `GOPATH`.
3. Inside the repo, run `go build ./...`

### To Prepare

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

### To Run

```bash
/simhash -n 18436187482098755291 -s 90 -d 7

Looking for hash 18436187482098755291 (ffda7ed5faffe6db), distance 7, since 90 days ago
533440 hashes returned
load took 5.86575223s
no matches found
hash search took 463.037µs
```

The experiment takes the following command line arguments:

`-n`: The hash, as a decimal integer, that you're looking for.

`-s`: The window backwards from now in days. Default is `10000` which should pick up all hashes for all time in the database.

`-d`: The Hamming distance to search for.

