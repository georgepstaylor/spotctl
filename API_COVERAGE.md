# Rackspace Spot CLI API Coverage

This document tracks the implementation status of the Rackspace Spot Public API endpoints in the CLI.

## API Coverage Table

| API Route                        | CLI Command                            | Status |
| -------------------------------- | -------------------------------------- | ------ |
| **Authentication**               |
| POST /oauth/token                | `spotctl` (automatic)                  | ✅     |
| **Regions**                      |
| GET /regions                     | `spotctl regions list`                 | ✅     |
| GET /regions/{name}              | `spotctl regions get <name>`           | ✅     |
| **Server Classes**               |
| GET /serverclasses               | `spotctl serverclasses list`           | ✅     |
| GET /serverclasses/{name}        | `spotctl serverclasses get <name>`     | ✅     |
| **Organizations**                |
| GET /organizations               | `spotctl organizations list`           | ✅     |
| **Cloudspaces**                  |
| GET /cloudspaces                 | `spotctl cloudspaces list`             | ✅     |
| POST /cloudspaces                | `spotctl cloudspaces create`           | ✅     |
| DELETE /cloudspaces/{name}       | `spotctl cloudspaces delete`           | ✅     |
| PATCH /cloudspaces/{name}        | `spotctl cloudspaces edit`             | ✅     |
| **Spot Node Pools**              |
| GET /spotnodepools               | `spotctl spotnodepool list`            | ✅     |
| POST /spotnodepools              | `spotctl spotnodepools create`         | ✅     |
| DELETE /spotnodepools            | `spotctl spotnodepools delete-all`     | ✅     |
| GET /spotnodepools/{name}        | `spotctl spotnodepool get <name>`      | ✅     |
| DELETE /spotnodepools/{name}     | `spotctl spotnodepools delete`         | ❌     |
| PATCH /spotnodepools/{name}      | `spotctl spotnodepools edit`           | ✅     |
| PUT /spotnodepools/{name}        | `spotctl spotnodepools edit`           | ❌     |
| **On-Demand Node Pools**         |
| GET /ondemandnodepools           | `spotctl ondemandnodepools list`       | ❌     |
| POST /ondemandnodepools          | `spotctl ondemandnodepools create`     | ❌     |
| DELETE /ondemandnodepools        | `spotctl ondemandnodepools delete-all` | ❌     |
| GET /ondemandnodepools/{name}    | `spotctl ondemandnodepools get <name>` | ❌     |
| DELETE /ondemandnodepools/{name} | `spotctl ondemandnodepools delete`     | ❌     |
| PATCH /ondemandnodepools/{name}  | `spotctl ondemandnodepools edit`       | ❌     |
| **Price Information**            |
| GET /price-history               | `spotctl price-history`                | ❌     |
| GET /percentile-info             | `spotctl percentile-info`              | ❌     |
| GET /market-price-capacity       | `spotctl market-price-capacity`        | ❌     |

Legend:

- ✅ Implemented
- ❌ Not Implemented

## Implementation Summary

**Implemented:** 14/24 endpoints (58.3%)
**Remaining:** 10/24 endpoints (41.7%)
