## 01. Check the commit contents
> [enable cce](https://github.com/PaaS-TA/pxc-release/commit/f5f04860195f400787e04a63f149a42fbe00cdaf)

## 02. Submodule modify : galera-init (src/github.com/cloudfoundry/galera-init)

> $ vi src/github.com/cloudfoundry/galera-init/db_helper/db_helper.go
```diff
 func FormatDSN(config config.DBHelper) string {
                Passwd: config.Password,
                Net:    "unix",
                Addr:   config.Socket,
+               AllowNativePasswords: true,
        }

```

> $ vi src/github.com/cloudfoundry/galera-init/go.mod
```diff
 require (
        github.com/Microsoft/go-winio v0.4.14 // indirect
        github.com/docker/docker v1.4.2-0.20181208172742-edf5134ba77d // indirect
        github.com/fsouza/go-dockerclient v1.3.6
-       github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c
+       github.com/go-sql-driver/mysql v1.5.0
        github.com/gogo/protobuf v1.2.1 // indirect
        github.com/google/uuid v1.1.0
        github.com/imdario/mergo v0.0.0-20160517064435-50d4dbd4eb0e // indirect

```

> $ vi src/github.com/cloudfoundry/galera-init/go.sum
```diff
 github.com/fsnotify/fsnotify v1.4.7 h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV
 github.com/fsnotify/fsnotify v1.4.7/go.mod h1:jwhsz4b93w/PPRr/qN1Yymfu8t87LnFCMoQvtojpjFo=
 github.com/fsouza/go-dockerclient v1.3.6 h1:oL0e3fpCjF+AHuUUBnwbkVcelFhxQifgTPQKipJPtnI=
 github.com/fsouza/go-dockerclient v1.3.6/go.mod h1:ptN6nXBwrXuiHAz2TYGOFCBB1aKGr371sGjMFdJEr1A=
-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c h1:QD/OSWIQcR3PMs9GzsjN5QOVvxvDI+WrK0GbvNapPds=
-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c/go.mod h1:zAC/RDZ24gD3HViQzih4MyKcchzm+sOG5ZlKdlhCg5w=
+github.com/go-sql-driver/mysql v1.5.0 h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=
+github.com/go-sql-driver/mysql v1.5.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=
 github.com/gogo/protobuf v1.2.0/go.mod h1:r8qH/GZQm5c6nD/R0oafs1akxWv10x8SbQlK7atdtwQ=
 github.com/gogo/protobuf v1.2.1 h1:/s5zKNz0uPFCZ5hddgPdo2TK2TVrUNMn0OOX8/aZMTE=
 github.com/gogo/protobuf v1.2.1/go.mod h1:hp+jE20tsWTFYpLwKvXlhS1hjn+gTNwPg2I6zVXpSg4=
```

> $ vi src/github.com/cloudfoundry/galera-init/vendor/modules.txt
```diff
 github.com/fsouza/go-dockerclient
 github.com/fsouza/go-dockerclient/internal/archive
 github.com/fsouza/go-dockerclient/internal/jsonmessage
 github.com/fsouza/go-dockerclient/internal/term
-# github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c
+# github.com/go-sql-driver/mysql v1.5.0
 ## explicit
 github.com/go-sql-driver/mysql
 # github.com/gogo/protobuf v1.2.1
```

> pull v1.5.0 go-sql-driver/mysql 
``` 
$ cd src/github.com/cloudfoundry/galera-init/vendor/github.com/go-sql-driver
$ rm -rf mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.5.0
```


## 02. Submodule modify : cf-mysql-cluster-health-logger (src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger)

> $ vi src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger/go.mod
```diff
 require (
        github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
-       github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c
+       github.com/go-sql-driver/mysql v1.5.0
        github.com/imdario/mergo v0.0.0-20160517064435-50d4dbd4eb0e // indirect
        github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
        github.com/onsi/ginkgo v1.12.0
```

> $ vi src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger/go.sum
```diff
 github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5/go.mod h1:a2zkGnVExMxdzMo3M0Hi/3sEU+cWnZpSni0O6/Yb/P0=
 github.com/fsnotify/fsnotify v1.4.7 h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=
 github.com/fsnotify/fsnotify v1.4.7/go.mod h1:jwhsz4b93w/PPRr/qN1Yymfu8t87LnFCMoQvtojpjFo=
-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c h1:QD/OSWIQcR3PMs9GzsjN5QOVvxvDI+WrK0GbvNapPds=
-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c/go.mod h1:zAC/RDZ24gD3HViQzih4MyKcchzm+sOG5ZlKdlhCg5w=
+github.com/go-sql-driver/mysql v1.5.0 h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=
+github.com/go-sql-driver/mysql v1.5.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=
 github.com/golang/protobuf v1.2.0 h1:P3YflyNX/ehuJFLhxviNdFxQPkGK5cDcApsge1SqnvM=
 github.com/golang/protobuf v1.2.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
 github.com/hpcloud/tail v1.0.0 h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=
```

> $ vi src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger/vendor/modules.txt
```diff
 # github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
 ## explicit
 github.com/erikstmartin/go-testdb
-# github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c
+# github.com/go-sql-driver/mysql v1.5.0
 ## explicit
 github.com/go-sql-driver/mysql
 # github.com/hpcloud/tail v1.0.0
```

> pull v1.5.0 go-sql-driver/mysql 
``` 
$ cd src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger/vendor/github.com/go-sql-driver
$ rm -rf mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.5.0
```