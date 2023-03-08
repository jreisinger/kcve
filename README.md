kcve is a small Go package and CLI tool to fetch, parse and enrich [official Kubernetes CVEs][1].

```
> kcve | head
CREATED     CLOSED      SUMMARY                            URL
2022-11-08  2022-11-10  Node address isn't always veri...  github.com/kubernetes/kubernetes/issues/113757
2022-11-08  2022-11-10  Unauthorized read of Custom Re...  github.com/kubernetes/kubernetes/issues/113756
2022-09-16  2022-09-16  Aggregated API server can caus...  github.com/kubernetes/kubernetes/issues/112513
2022-09-01  2022-09-15  `runAsNonRoot` logic bypass fo...  github.com/kubernetes/kubernetes/issues/112192
2021-09-13  2021-09-16  Symlink Exchange Can Allow Hos...  github.com/kubernetes/kubernetes/issues/104980
2021-05-18  2021-05-18  Holes in EndpointSlice Validat...  github.com/kubernetes/kubernetes/issues/102106
2021-04-23  2021-04-23  Processes may panic upon recei...  github.com/kubernetes/kubernetes/issues/101435
2021-03-10  2021-04-14  Validating Admission Webhook d...  github.com/kubernetes/kubernetes/issues/100096
2020-12-04  2020-12-07  Man in the middle using LoadBa...  github.com/kubernetes/kubernetes/issues/97076
```

[1]: https://kubernetes.io/docs/reference/issues-security/official-cve-feed/
