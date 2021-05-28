## 01. Subproject modify : galera-init (src/github.com/cloudfoundry/galera-init)
``` 
$ vi db_helper/db_helper.go
```
func FormatDSN(config config.DBHelper) string {
                Passwd: config.Password,
                Net:    "unix",
<span style="color:green">+                Addr:   config.Socket,</span>
                AllowNativePasswords: true,
}

``` 
$ vi go.mod
```
<span style="color:red">-       github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c</span>
<span style="color:green">+        github.com/go-sql-driver/mysql v1.5.0</span>

``` 
$ vi go.sum
```
<span style="color:red">-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c h1:QD/OSWIQcR3PMs9GzsjN5QOVvxvDI+WrK0GbvNapPds=</span>
<span style="color:red">-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c/go.mod h1:zAC/RDZ24gD3HViQzih4MyKcchzm+sOG5ZlKdlhCg5w=</span>
<span style="color:green">+github.com/go-sql-driver/mysql v1.5.0 h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=</span>
<span style="color:green">+github.com/go-sql-driver/mysql v1.5.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=</span>

``` 
$ vi vendor/modules.txt
```
<span style="color:red">-# github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c</span>
<span style="color:green">+# github.com/go-sql-driver/mysql v1.5.0</span>

``` 
$ cd vendor/github.com/go-sql-driver
$ rm -rf mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.5.0
```

## 02. Subproject modify : cf-mysql-cluster-health-logger (src/github.com/cloudfoundry-incubator/cf-mysql-cluster-health-logger)
``` 
$ vi go.mod
```
<span style="color:red">-       github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c</span>
<span style="color:green">+        github.com/go-sql-driver/mysql v1.5.0</span>

``` 
$ vi go.sum
```
<span style="color:red">-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c h1:QD/OSWIQcR3PMs9GzsjN5QOVvxvDI+WrK0GbvNapPds=</span>
<span style="color:red">-github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c/go.mod h1:zAC/RDZ24gD3HViQzih4MyKcchzm+sOG5ZlKdlhCg5w=</span>
<span style="color:green">+github.com/go-sql-driver/mysql v1.5.0 h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=</span>
<span style="color:green">+github.com/go-sql-driver/mysql v1.5.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=</span>

``` 
$ vi vendor/modules.txt
```
<span style="color:red">-# github.com/go-sql-driver/mysql v1.2.1-0.20160802113842-0b58b37b664c</span>
<span style="color:green">+# github.com/go-sql-driver/mysql v1.5.0</span>

``` 
$ cd vendor/github.com/go-sql-driver
$ rm -rf mysql
$ git clone https://github.com/go-sql-driver/mysql.git -b v1.5.0
```
