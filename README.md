# ansible-piMetrics
ansible for installing Grafana with InfluxDB on RaspberryPi

## All projects
### Secrets:
* Upload your ssh key to your pi else ansible will fail
* All roles have a copy of the same vault (temp until a golbal vault is created) - Update as necessary

  ```
  db_admin: 'xxx' #influx db admin username
  db_admin_pass: 'xxx' #influx db admin password
  ```

---
## Grafana with Influx Datastore on RaspberryPi3
### Notes:
* Home setup for Grafana and InfluxDB on a raspberryPi3
* Tested on Stretch - NOOBS

### Usage:
* Update hosts with IP or Hostname under the [piGrafana]

  ```
  --ask-sudo-pass may be required if running reboot role due to your local setup
  ansible-playbook piMetrics_setup.yml --ask-vault-pass -i hosts --ask-sudo-pass
  ```

---

```
Grafana with InfluxDB
https://github.com/fg2it/grafana-on-raspberry/wiki
https://www.circuits.dk/install-grafana-influxdb-raspberry/
```
