# Decision Log Payload Size Debugging

Created for [question here](https://github.com/orgs/open-policy-agent/discussions/332) in order to replicate issue.

Steps to use this repo:

1. Clone this repo
2. Install OPA, make sure is in path
2. Run `opa build -o storage/policy.tar.gz -b storage/policy/`
3. Run `opa build -o storage/discovery.tar.gz -b storage/disc`
4. In terminal 1, `cd storage && go run main.go`
5. In terminal 2, `cd declog && go run main.go`
6. In terminal 3, `opa run -s --config-file=config.yaml`
7. Observe the logs in terminal 3, has OPA loaded the bundles ok?
8. In terminal 4, `cd bench && go run main.go`
9. Observe the logs in terminal 2, after 30s OPA sends decision logs. They should all be less than 1048576 bytes in size.

The input used in the benchmark is around 450kB. I see log sizes of 594kB.



