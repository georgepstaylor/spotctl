# Rackspace Spot CLI API Coverage

This document tracks the implementation status of the Rackspace Spot Public API endpoints in the CLI.

## API Coverage Table

| API Route                        | CLI Command                           | Status |
| -------------------------------- | ------------------------------------- | ------ |
| **Authentication**               |
| POST /oauth/token                | `spotctl` (automatic)                 | ✅     |
| **Regions**                      |
| GET /regions                     | `spotctl regions list`                | ✅     |
| GET /regions/{name}              | `spotctl regions get <name>`          | ✅     |
| **Server Classes**               |
| GET /serverclasses               | `spotctl serverclasses list`          | ✅     |
| GET /serverclasses/{name}        | `spotctl serverclasses get <name>`    | ✅     |
| **Organizations**                |
| GET /organizations               | `spotctl organizations list`          | ✅     |
| **Cloudspaces**                  |
| GET /cloudspaces                 | `spotctl cloudspaces list`            | ✅     |
| POST /cloudspaces                | `spotctl cloudspaces create`          | ✅     |
| DELETE /cloudspaces/{name}       | `spotctl cloudspaces delete`          | ✅     |
| PATCH /cloudspaces/{name}        | `spotctl cloudspaces edit`            | ✅     |
| **Spot Node Pools**              |
| GET /spotnodepools               | `spotctl spotnodepool list`           | ✅     |
| POST /spotnodepools              | `spotctl spotnodepools create`        | ✅     |
| DELETE /spotnodepools            | `spotctl spotnodepools delete-all`    | ✅     |
| GET /spotnodepools/{name}        | `spotctl spotnodepool get <name>`     | ✅     |
| DELETE /spotnodepools/{name}     | `spotctl spotnodepool delete`         | ✅     |
| PATCH /spotnodepools/{name}      | `spotctl spotnodepools edit`          | ✅     |
| **On-Demand Node Pools**         |
| GET /ondemandnodepools           | `spotctl ondemandnodepool list`       | ✅     |
| POST /ondemandnodepools          | `spotctl ondemandnodepool create`     | ❌     |
| DELETE /ondemandnodepools        | `spotctl ondemandnodepool delete-all` | ❌     |
| GET /ondemandnodepools/{name}    | `spotctl ondemandnodepool get <name>` | ✅     |
| DELETE /ondemandnodepools/{name} | `spotctl ondemandnodepool delete`     | ❌     |
| PATCH /ondemandnodepools/{name}  | `spotctl ondemandnodepool edit`       | ❌     |
| **Price Information**            |
| GET /price-history               | `spotctl price-history`               | ❌     |
| GET /percentile-info             | `spotctl percentile-info`             | ❌     |
| GET /market-price-capacity       | `spotctl market-price-capacity`       | ❌     |

Legend:

- ✅ Implemented
- ❌ Not Implemented

## Implementation Summary

**Implemented:** 18/25 endpoints (72.0%)
**Remaining:** 7/25 endpoints (28.0%)
