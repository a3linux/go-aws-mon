## Setup systemd sceduled task (tested on CoreOS)
* Put go-aws-mon.service and go-aws-mon.timer to /etc/systemd/system/
* If you don't used CoreOS, please change the ExecStart line in go-aws-mon.service to where you put go-aws-mon.sh
* Run the following commands,
    ```
    sudo systemctl enable go-aws-mon.service
    sudo systemctl enable go-aws-mon.timer
    sudo systemctl start go-aws-mon.service
    sudo systemctl start go-aws-mon.timer
    ```

* Enjoy your metric on AWS CloudWatch
