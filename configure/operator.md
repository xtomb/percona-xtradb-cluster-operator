Custom Resource options
==============================================================

The operator is configured via the spec section of the [deploy/cr.yaml](https://github.com/Percona-Lab/percona-xtradb-cluster-operator/blob/master/deploy/cr.yaml) file. This file contains the following spec sections to configure three main subsystems of the cluster: 

| Key      | Value Type | Description                               |
|----------|------------|-------------------------------------------|
| pxc      | subdoc     | Percona XtraDB Cluster general section    |
| proxysql | subdoc     | ProxySQL section                          |
| pmm      | subdoc     | Percona Monitoring and Management section |
| backup   | subdoc     | Percona XtraDB Cluster backups section    |

### PXC Section

The ``pxc`` section in the deploy/cr.yaml file contains general configuration options for the Percona XtraDB Cluster.

| Key                            | Value Type | Example   | Description |
|--------------------------------|------------|-----------|-------------|
|size                            | int        | `3`       |  The size of the Percona XtraDB Cluster, must be >= 3 for [High-Availability](hhttps://www.percona.com/doc/percona-xtradb-cluster/5.7/intro.html) |
|image                           | string     |`perconalab/pxc-openshift:0.1.0` | Percona XtraDB Cluster docker image to use                                                                     |
|resources.requests.memory       | string     | `1G`      | [Kubernetes Memory requests](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a PXC container                                                               |
|resources.requests.cpu          | string     | `600m`    | [Kubernetes CPU requests](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a PXC container |
|resources.limits.memory         | string     | `1G`      | [Kubernetes Memory limit](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a PXC container |
|resources.limits.cpu            | string     | `1`       | [Kubernetes CPU limit](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a PXC container |
|volumeSpec.storageClass         | string     | `standard`| Set the [Kubernetes Storage Class](https://kubernetes.io/docs/concepts/storage/storage-classes/) to use with the PXC [Persistent Volume Claim](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims)                     |
|volumeSpec.accessModes          | array      | `[ "ReadWriteOnce" ]` | [Kubernetes Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) access modes for the PerconaXtraDB Cluster  |
|volumeSpec.size                 | string     | `6Gi`     | The [Kubernetes Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) size for the Percona XtraDB Cluster                            |
|affinity.topologyKey            | string     |`kubernetes.io/hostname`| The [Kubernetes topologyKey](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) node affinity constraint|
|affinity.advanced               | subdoc     |           | If available, it makes [topologyKey](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) node affinity constraint to be ignored |
|nodeSelector                    | label      | `disktype: ssd`        | The [Kubernetes nodeSelector](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector) constraint|
|tolerations                     | subdoc     | `node.alpha.kubernetes.io/unreachable` | The [Kubernetes Pod tolerations] (https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/#concepts)            |
|priorityClassName               | string     | `high-priority`  | The [Kuberentes Pod priority class](https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/#priorityclass) |
|annotations.iam.amazonaws.com/role | string |`role-arn`| The [AWS IAM role](https://kubernetes-on-aws.readthedocs.io/en/latest/user-guide/iam-roles.html)                              |
|imagePullSecrets.name           | string     | `private-registry-credentials` | [Kubernetes imagePullSecret](https://kubernetes.io/docs/concepts/configuration/secret/#using-imagepullsecrets) for the Percona XtraDB Cluster docker image |
|labels                          | label      | `rack: rack-22` | The [Kubernetes affinity labels](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/)                       |

### ProxySQL Section

The ``proxysql`` section in the deploy/cr.yaml file contains configuration options for the ProxySQL daemon.

| Key                            | Value Type | Example   | Description |
|--------------------------------|------------|-----------|-------------|
|enabled                         | boolean    | `true`    | Enables or disables [load balancing with ProxySQL](https://www.percona.com/doc/percona-xtradb-cluster/5.7/howtos/proxysql.html) [Service](https://kubernetes.io/docs/concepts/services-networking/service/) |
|size                            | int        | `1`       | The number of the ProxySQL daemons [to provide load balancing](https://www.percona.com/doc/percona-xtradb-cluster/5.7/howtos/proxysql.html), must be = 1 in current release|
|image                           | string     |`perconalab/proxysql-openshift:0.1.0` | ProxySQL docker image to use |
|resources.requests.memory       | string     | `1G`      | [Kubernetes Memory requests](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a ProxySQL container                                                      |
|resources.requests.cpu          | string     | `600m`    | [Kubernetes CPU requests](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a ProxySQL container                                                               |
|resources.limits.memory| string     | `1G`      | [Kubernetes Memory limit](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a ProxySQL container                                                               |
|resources.limits.cpu   | string     | `700m`    | [Kubernetes CPU limit](https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#resource-requests-and-limits-of-pod-and-container) for a ProxySQL container                                                               |
|volumeSpec.storageClass         | string     | `standard`| The [Kubernetes Storage Class](https://kubernetes.io/docs/concepts/storage/storage-classes/) to use with the ProxySQL [Persistent Volume Claim](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims)            |
|volumeSpec.accessModes          | array      | `[ "ReadWriteOnce" ]` | [Kubernetes Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) access modes for ProxySQL  |
|volumeSpec.size                 | string     | `2Gi`     | The [Kubernetes Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) size for ProxySQL                             |
|affinity.topologyKey            | string     |`failure-domain.beta.kubernetes.io/zone`| The [Kubernetes topologyKey](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) node affinity constraint|
|affinity.advanced               | subdoc     |           | If available, it makes [topologyKey](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#inter-pod-affinity-and-anti-affinity-beta-feature) node affinity constraint to be ignored |
|nodeSelector           | label      | `disktype: ssd`        | The [Kubernetes nodeSelector](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector) affinity constraint|
|tolerations            | subdoc     | `node.alpha.kubernetes.io/unreachable` | The [Kubernetes Pod tolerations] (https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/#concepts)            |

### PMM Section

The ``pmm`` section in the deploy/cr.yaml file contains configuration options for Percona Monitoring and Management.

| Key       | Value Type | Example               | Description                    |
|-----------|------------|-----------------------|--------------------------------|
|enabled    | boolean    | `false`               | Enables or disables [monitoring Percona XtraDB Cluster with PMM](https://www.percona.com/doc/percona-xtradb-cluster/LATEST/manual/monitoring.html#using-pmm) |
|image      | string     |`perconalab/pmm-client`| PMM Client docker image to use |
|serverHost | string     | `monitoring-service`  | Address of the PMM Server to collect data from the Cluster |
|serverUser | string     | `pmm`                 | The [PMM Server user](https://www.percona.com/doc/percona-monitoring-and-management/glossary.option.html#term-server-user). The PMM Server Password should be configured via secrets. |

## backup section

The ``backup`` section in the [deploy/cr.yaml](https://github.com/Percona-Lab/percona-xtradb-cluster-operator/blob/master/deploy/cr.yaml) file contains the following configuration options for the regular Percona XtraDB Cluster backups.

| Key                            | Value Type | Default   | Description |
|--------------------------------|------------|-----------|-------------|
|image                           | string     | `perconalab/backupjob-openshift:0.2.0` | Percona XtraDB Cluster docker image to use for the backup functionality                                                                       |
|imagePullSecrets.name           | string     | `private-registry-credentials`  | [Kubernetes imagePullSecret](https://kubernetes.io/docs/concepts/configuration/secret/#using-imagepullsecrets) for the specified docker image |
|schedule.name                   | string     | `sat-night-backup` | Name of the backup             |
|schedule.schedule               | string     | `0 0 * * 6`        | Scheduled time to make a backup, specified in the [crontab format](https://en.wikipedia.org/wiki/Cron)                                                        |
|schedule.keep                   | int        | `3`       | Number of backups to store             |
|volume.storageClass             | string     | `standard`| Set the [Kubernetes Storage Class](https://kubernetes.io/docs/concepts/storage/storage-classes/) to use with the PXC backups [Persistent Volume Claim](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims)                     |
|volume.size                     | string     | `6Gi`     | The [Kubernetes Persistent Volume](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) size for the Percona XtraDB Cluster backups                           |