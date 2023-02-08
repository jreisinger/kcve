package gh

import (
	"fmt"
	"testing"
)

var headers = `HTTP/2.0 200 OK
Accept-Ranges: bytes
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Used, X-RateLimit-Resource, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type, X-GitHub-SSO, X-GitHub-Request-Id, Deprecation, Sunset
Cache-Control: public, max-age=60, s-maxage=60
Content-Security-Policy: default-src 'none'
Content-Type: application/json; charset=utf-8
Date: Wed, 08 Feb 2023 13:32:07 GMT
Etag: W/"ff0c446de0ab420a0b6a3a0f31ce7505749cd96b51475bb07dfe82a64d4026e4"
Last-Modified: Sat, 12 Nov 2022 03:44:24 GMT
Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
Server: GitHub.com
Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
Vary: Accept, Accept-Encoding, Accept, X-Requested-With
X-Content-Type-Options: nosniff
X-Frame-Options: deny
X-Github-Api-Version-Selected: 2022-11-28
X-Github-Media-Type: github.v3; format=json
X-Github-Request-Id: E038:E2C7:1520B47:156F070:63E3A457
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 59
X-Ratelimit-Reset: 1675866727
X-Ratelimit-Resource: core
X-Ratelimit-Used: 1
X-Xss-Protection: 0`

var body = `{"url":"https://api.github.com/repos/kubernetes/kubernetes/issues/113757","repository_url":"https://api.github.com/repos/kubernetes/kubernetes","labels_url":"https://api.github.com/repos/kubernetes/kubernetes/issues/113757/labels{/name}","comments_url":"https://api.github.com/repos/kubernetes/kubernetes/issues/113757/comments","events_url":"https://api.github.com/repos/kubernetes/kubernetes/issues/113757/events","html_url":"https://github.com/kubernetes/kubernetes/issues/113757","id":1440935103,"node_id":"I_kwDOAToIks5V4uy_","number":113757,"title":"CVE-2022-3294: Node address isn't always verified when proxying","user":{"login":"tallclair","id":29742491,"node_id":"MDQ6VXNlcjI5NzQyNDkx","avatar_url":"https://avatars.githubusercontent.com/u/29742491?v=4","gravatar_id":"","url":"https://api.github.com/users/tallclair","html_url":"https://github.com/tallclair","followers_url":"https://api.github.com/users/tallclair/followers","following_url":"https://api.github.com/users/tallclair/following{/other_user}","gists_url":"https://api.github.com/users/tallclair/gists{/gist_id}","starred_url":"https://api.github.com/users/tallclair/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/tallclair/subscriptions","organizations_url":"https://api.github.com/users/tallclair/orgs","repos_url":"https://api.github.com/users/tallclair/repos","events_url":"https://api.github.com/users/tallclair/events{/privacy}","received_events_url":"https://api.github.com/users/tallclair/received_events","type":"User","site_admin":false},"labels":[{"id":105146071,"node_id":"MDU6TGFiZWwxMDUxNDYwNzE=","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/kind/bug","name":"kind/bug","color":"e11d21","default":false,"description":"Categorizes issue or PR as related to a bug."},{"id":116712923,"node_id":"MDU6TGFiZWwxMTY3MTI5MjM=","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/area/security","name":"area/security","color":"d93f0b","default":false,"description":null},{"id":136601536,"node_id":"MDU6TGFiZWwxMzY2MDE1MzY=","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/area/apiserver","name":"area/apiserver","color":"0052cc","default":false,"description":null},{"id":173493835,"node_id":"MDU6TGFiZWwxNzM0OTM4MzU=","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/sig/api-machinery","name":"sig/api-machinery","color":"d2b48c","default":false,"description":"Categorizes an issue or PR as relevant to SIG API Machinery."},{"id":1199275492,"node_id":"MDU6TGFiZWwxMTk5Mjc1NDky","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/committee/security-response","name":"committee/security-response","color":"c0ff4a","default":false,"description":"Denotes an issue or PR intended to be handled by the product security committee."},{"id":2389856656,"node_id":"MDU6TGFiZWwyMzg5ODU2NjU2","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/triage/accepted","name":"triage/accepted","color":"8fc951","default":false,"description":"Indicates an issue or PR is ready to be actively worked on."},{"id":3603068678,"node_id":"LA_kwDOAToIks7WwncG","url":"https://api.github.com/repos/kubernetes/kubernetes/labels/official-cve-feed","name":"official-cve-feed","color":"0052cc","default":false,"description":"Issues or PRs related to CVEs officially announced by Security Response Committee (SRC)"}],"state":"closed","locked":false,"assignee":null,"assignees":[],"milestone":null,"comments":2,"created_at":"2022-11-08T21:33:26Z","updated_at":"2022-11-12T03:44:24Z","closed_at":"2022-11-10T17:39:50Z","author_association":"MEMBER","active_lock_reason":null,"body":"CVSS Rating: [CVSS:3.1/AV:N/AC:H/PR:H/UI:N/S:U/C:H/I:H/A:H](https://www.first.org/cvss/calculator/3.1#CVSS:3.1/AV:N/AC:H/PR:H/UI:N/S:U/C:H/I:H/A:H)\r\n\r\nA security issue was discovered in Kubernetes where users may have access to secure endpoints in the control plane network. Kubernetes clusters are only affected if an untrusted user can modify Node objects and send proxy requests to them.\r\n\r\nKubernetes supports node proxying, which allows clients of kube-apiserver to access endpoints of a Kubelet to establish connections to Pods, retrieve container logs, and more. While Kubernetes already validates the proxying address for Nodes, a bug in kube-apiserver made it possible to bypass this validation. Bypassing this validation could allow authenticated requests destined for Nodes to to the API server's private network.\r\n\r\n### Am I vulnerable?\r\n\r\nClusters are affected by this vulnerability if there are endpoints that the kube-apiserver has connectivity to that users should not be able to access. This includes:\r\n\r\n- kube-apiserver is in a separate network from worker nodes\r\n- localhost services\r\n\r\nmTLS services that accept the same client certificate as nodes may be affected. The severity of this issue depends on the privileges & sensitivity of the exploitable endpoints.\r\n\r\nClusters that configure the egress selector to use a proxy for cluster traffic may not be affected.\r\n\r\n\r\n#### Affected Versions\r\n\r\n- Kubernetes kube-apiserver <= v1.25.3\r\n- Kubernetes kube-apiserver <= v1.24.7\r\n- Kubernetes kube-apiserver <= v1.23.13\r\n- Kubernetes kube-apiserver <= v1.22.15\r\n\r\n### How do I mitigate this vulnerability?\r\n\r\nUpgrading the **kube-apiserver** to a fixed version mitigates this vulnerability.\r\n\r\nAside from upgrading, configuring an [egress proxy for egress to the cluster network](https://kubernetes.io/docs/tasks/extend-kubernetes/setup-konnectivity/) can mitigate this vulnerability.\r\n\r\n#### Fixed Versions\r\n\r\n- Kubernetes kube-apiserver v1.25.4\r\n- Kubernetes kube-apiserver v1.24.8\r\n- Kubernetes kube-apiserver v1.23.14\r\n- Kubernetes kube-apiserver v1.22.16\r\n\r\n**Fix impact:** In some cases, the fix can break clients that depend on the nodes/proxy subresource, specifically if a kubelet advertises a localhost or link-local address to the Kubernetes control plane.\r\n\r\n### Detection\r\n\r\nNode create & update requests may be included in the Kubernetes audit log, and can be used to identify requests for IP addresses that should not be permitted. Node proxy requests may also be included in audit logs.\r\n\r\nIf you find evidence that this vulnerability has been exploited, please contact security@kubernetes.io\r\n\r\n#### Acknowledgements\r\n\r\nThis vulnerability was reported by Yuval Avrahami of Palo Alto Networks.\r\n\r\n<!-- labels -->\r\n/area security\r\n/kind bug\r\n/committee security-response\r\n/label official-cve-feed\r\n/sig api-machinery\r\n/area apiserver","closed_by":{"login":"tallclair","id":29742491,"node_id":"MDQ6VXNlcjI5NzQyNDkx","avatar_url":"https://avatars.githubusercontent.com/u/29742491?v=4","gravatar_id":"","url":"https://api.github.com/users/tallclair","html_url":"https://github.com/tallclair","followers_url":"https://api.github.com/users/tallclair/followers","following_url":"https://api.github.com/users/tallclair/following{/other_user}","gists_url":"https://api.github.com/users/tallclair/gists{/gist_id}","starred_url":"https://api.github.com/users/tallclair/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/tallclair/subscriptions","organizations_url":"https://api.github.com/users/tallclair/orgs","repos_url":"https://api.github.com/users/tallclair/repos","events_url":"https://api.github.com/users/tallclair/events{/privacy}","received_events_url":"https://api.github.com/users/tallclair/received_events","type":"User","site_admin":false},"reactions":{"url":"https://api.github.com/repos/kubernetes/kubernetes/issues/113757/reactions","total_count":0,"+1":0,"-1":0,"laugh":0,"hooray":0,"confused":0,"heart":0,"rocket":0,"eyes":0},"timeline_url":"https://api.github.com/repos/kubernetes/kubernetes/issues/113757/timeline","performed_via_github_app":null,"state_reason":"completed"}`

func TestRemoveHeaders(t *testing.T) {
	tests := []struct {
		in   []byte
		want []byte
	}{
		{
			in:   []byte(fmt.Sprintf("%s\r\n\r\n%s", headers, body)),
			want: []byte(body),
		},
		{
			in:   []byte{},
			want: []byte{},
		},
		{
			in:   []byte(headers),
			want: []byte{},
		},
	}
	for _, test := range tests {
		if got := removeResponseHeaders(test.in); string(got) != string(test.want) {
			t.Errorf("want %s, got %s", test.want, got)
		}
	}
}
