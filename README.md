# pgoproxy
for personal use only

Changes:

- Remove sumdb_client_ops.go:load() some extra logic, which use proxy as sumdb even if the GOSUMDB is not empty. 
- Modify cacher/disk.go, and remove the rest of cacher.