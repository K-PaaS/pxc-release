## 01. Submodule modify : galera-init (src/github.com/cloudfoundry/galera-init)

> $ vi src/github.com/cloudfoundry/galera-init/db_helper/db_helper.go
```diff
 func FormatDSN(config config.DBHelper) string {
        skipBinLog := ""
+       allowNativePasswords := ""
+
        if config.SkipBinlog {
                skipBinLog = "?sql_log_bin=off"
+               allowNativePasswords = "&allowNativePasswords=true"
+       } else {
+               allowNativePasswords = "?allowNativePasswords=true"
        }
-       return fmt.Sprintf("%s:%s@unix(%s)/%s", config.User, config.Password, config.Socket, skipBinLog)
+       return fmt.Sprintf("%s:%s@unix(%s)/%s%s", config.User, config.Password, config.Socket, skipBinLog, allowNativePasswords)
 }
```

## 02. Submodule modify : galera-healthcheck - go-sql-driver (src/github.com/cloudfoundry-incubator/galera-healthcheck)

> $ vi src/github.com/cloudfoundry-incubator/galera-healthcheck/vendor/github.com/go-sql-driver/mysql/auth.go
```diff
func (mc *mysqlConn) auth(authData []byte, plugin string) ([]byte, error) {
        switch plugin {
        case "caching_sha2_password":
                authResp := scrambleSHA256Password(authData, mc.cfg.Passwd)
                return authResp, nil

        case "mysql_old_password":
                if !mc.cfg.AllowOldPasswords {
                        return nil, ErrOldPassword
                }
                if len(mc.cfg.Passwd) == 0 {
                        return nil, nil
                }
                // Note: there are edge cases where this should work but doesn't;
                // this is currently "wontfix":
                // https://github.com/go-sql-driver/mysql/issues/184
                authResp := append(scrambleOldPassword(authData[:8], mc.cfg.Passwd), 0)
                return authResp, nil

        case "mysql_clear_password":
                if !mc.cfg.AllowCleartextPasswords {
                        return nil, ErrCleartextPassword
                }
                // http://dev.mysql.com/doc/refman/5.7/en/cleartext-authentication-plugin.html
                // http://dev.mysql.com/doc/refman/5.7/en/pam-authentication-plugin.html
                return append([]byte(mc.cfg.Passwd), 0), nil

        case "mysql_native_password":
                if !mc.cfg.AllowNativePasswords {
                        return nil, ErrNativePassword
                }
                // https://dev.mysql.com/doc/internals/en/secure-password-authentication.html
                // Native password authentication only need and will need 20-byte challenge.
                authResp := scramblePassword(authData[:20], mc.cfg.Passwd)
                return authResp, nil


        case "sha256_password":
                if len(mc.cfg.Passwd) == 0 {
                        return []byte{0}, nil
                }
-               if mc.cfg.tls != nil || mc.cfg.Net == "unix" {
+               if mc.cfg.tls != nil {
                        // write cleartext auth packet
                        return append([]byte(mc.cfg.Passwd), 0), nil
                }
```
