package config

discovery = {
 "bundles": {
    "authz": {
      "service": "storage-service",
      "resource": "policy.tar.gz",
      "polling": {
        "min_delay_seconds": 1,
        "max_delay_seconds": 20
      }
    }
  },
    "decision_logs": {
        "reporting": {
            "upload_size_limit_bytes": 1048576,
            "min_delay_seconds": 30,
            "max_delay_seconds": 30,
        },
        "service": "decision-log-service"
    }
}