# lst

Displays a tree view of directories.

Built in go for learning purposes.

Based on `tree` by Steve Barker 
([source](http://mama.indstate.edu/users/ice/tree/)).

### Installing

[Install go](http://golang.org/) and then run:

```
go get github.com/mcls/lst
```

### Example

```
$ lst --level 2 activerecord/
activerecord/
|-- CHANGELOG.md
|-- MIT-LICENSE
|-- README.rdoc
|-- RUNNING_UNIT_TESTS.rdoc
|-- Rakefile
|-- activerecord.gemspec
|-- examples/
|   |-- performance.rb
|   `-- simple.rb
|-- lib/
|   |-- active_record/
|   |-- active_record.rb
|   `-- rails/
`-- test/
    |-- active_record/
    |-- assets/
    |-- cases/
    |-- config.example.yml
    |-- config.rb
    |-- fixtures/
    |-- migrations/
    |-- models/
    |-- schema/
    `-- support/
```
