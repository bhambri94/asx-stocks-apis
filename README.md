### asx-stocks-apis

Setup

```

git clone https://github.com/bhambri94/asx-stocks-apis.git

cd asx-stocks-apis/

docker build -t asx-stocks-apis:v1.0 .

docker images ls

docker run -it --name asx-stocks-apis -v $PWD/src:/go/src/asx-stocks-apis asx-stocks-apis:v1.0

```

#### Cron job

To setup a Daily Cron job, please follow following steps:
 
```
cd asx-stocks-apis/

Vi bash.sh

```
```
#!/bin/bash
sudo /usr/bin/docker restart asx-stocks-apis
```

Save the sheet script and run command 

```
chmod 777 bash.sh

Crontab -e

* 9 * * * /path_to_asx-stocks-apis_repo/bash.sh > /path_to_asx-stocks-apis_repo/asx-stocks-apis.logs

```
This above command written in crontab will run the script daily at 9AM UTC time.
