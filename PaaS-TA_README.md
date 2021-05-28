## 01. Subproject modify : galera-init (src/github.com/cloudfoundry/galera-init)

> $ vi db_helper/db_helper.go
```diff
		func FormatDSN(config config.DBHelper) string {
                Passwd: config.Password,
                Net:    "unix",
                Addr:   config.Socket,
+               AllowNativePasswords: true,
        }
```

> $ vi go.mod
```diff
 require (
        github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
-       github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c
+       github.com/go-sql-driver/mysql v1.5.0
        github.com/imdario/mergo v0.0.0-20160517064435-50d4dbd4eb0e // indirect
        github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
        github.com/onsi/ginkgo v1.12.0
```

> $ vi go.sum
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

> $ vi vendor/modules.txt
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

> download go-sql-driver v1.5.0
``` 
$ cd vendor/github.com/go-sql-driver
$ rm -rf mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.5.0
```


## 02. Subproject modify : cf-mysql-cluster-health-logger (src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger)

> $ vi go.mod
```diff
 require (
        github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
-       github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c
+       github.com/go-sql-driver/mysql v1.5.0
        github.com/imdario/mergo v0.0.0-20160517064435-50d4dbd4eb0e // indirect
        github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
        github.com/onsi/ginkgo v1.12.0
```

> $ vi go.sum
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

> $ vi vendor/modules.txt
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

> download go-sql-driver v1.5.0
``` 
$ cd vendor/github.com/go-sql-driver
$ rm -rf mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.5.0
```