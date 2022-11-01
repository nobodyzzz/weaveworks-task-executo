Examples of usage:
```
~/dev/weaveworks-task-executor$ go build main.go
~/dev/weaveworks-task-executor$ ./main example.yml
[example.yml] 2022/11/01 11:24:59 START task 'Create root directory'
[example.yml] 2022/11/01 11:24:59 END task 'Create root directory'
[example.yml] 2022/11/01 11:24:59 START task 'Create VERSION file'
[example.yml] 2022/11/01 11:24:59 END task 'Create VERSION file'
[example.yml] 2022/11/01 11:24:59 START task 'Set VERSION'
[example.yml] 2022/11/01 11:24:59 END task 'Set VERSION'
[example.yml] 2022/11/01 11:24:59 START task 'Clean up'
[example.yml] 2022/11/01 11:24:59 END task 'Clean up'
~/dev/weaveworks-task-executor$ ./main example.yml bad-example.yml
[example.yml] 2022/11/01 11:25:03 START task 'Create root directory'
[example.yml] 2022/11/01 11:25:03 END task 'Create root directory'
[example.yml] 2022/11/01 11:25:03 START task 'Create VERSION file'
[example.yml] 2022/11/01 11:25:03 END task 'Create VERSION file'
[example.yml] 2022/11/01 11:25:03 START task 'Set VERSION'
[example.yml] 2022/11/01 11:25:03 END task 'Set VERSION'
[example.yml] 2022/11/01 11:25:03 START task 'Clean up'
[example.yml] 2022/11/01 11:25:03 END task 'Clean up'
[bad-example.yml] 2022/11/01 11:25:03 START task 'Create root directory'
[bad-example.yml] 2022/11/01 11:25:03 Task 'Create root directory' failed (mkdir /root/tmp/project: no such file or directory)
[bad-example.yml] 2022/11/01 11:25:03 SKIP task 'Create VERSION file'
[bad-example.yml] 2022/11/01 11:25:03 START task 'Set VERSION'
[bad-example.yml] 2022/11/01 11:25:03 Task 'Set VERSION' failed (open /root/tmp/project/VERSION: no such file or directory)
[bad-example.yml] 2022/11/01 11:25:03 START task 'Clean up'
[bad-example.yml] 2022/11/01 11:25:03 END task 'Clean up'
[bad-example.yml] 2022/11/01 11:25:03 Unknown task type 'fix_the_world'. Supported tasks: [create_dir create_file put_content rm_dir rm_file]
~/dev/weaveworks-task-executor$ ./main -silent example.yml bad-example.yml
~/dev/weaveworks-task-executor$
```

Possible imporvments:
  - ability to set permissions for `create_dir` and `create_file` tasks
  - more robust testing
  - probably many more =)
