# ADRgen

> work in progress


[![<ORG_NAME>](https://circleci.com/gh/asiermarques/adrgen.svg?style=svg)](https://circleci.com/gh/asiermarques/adrgen)

Another little tool for generating Architecture Decision Records


## Getting started

### Build the binary
```
go build -o adrgen
```

### Initializing the project and configuration

We will use the **init** command specifying where the ADR files will be written. 

```
adrgen init "docs/adrs"
```

This command creates the following structure:

```
your_dir
├── adrgen.config.yml
└── docs
    └── adrs
        └── adr_template.md
```

As the result, we can see
* A config file is created
* A directory structure is created if it doesn't exist
* A markdown template is created in the desired directory


The adrgen.config.yml config file will be used by other commands in order to know how to operate with the ADR files.

It will include the following configuration keys:

| key                | type       | description                                                                   |
|--------------------|------------|-------------------------------------------------------------------------------|
| directory          | string     | the directory where the ADR files will be managed by adrgen                   |
| default_meta       | array      | the keys for meta that will include in all ADR files                          |
| supported_statuses | array      | the statuses that will be supported for the ADRs                              |
| default_status     | string     | the status that the ADR status will be set by default in the creation process |
| template_file      | string     | the template file that will be used to generate the ADR files                 |
| id_digit_number    | int        | the number of digits for the ADR identifier, for example: "0001-"             |

Example:

````
default_meta: []
default_status: proposed
directory: docs/adrs
supported_statuses:
- proposed
- accepted
- rejected
- superseeded
- amended
- deprecated
template_file: docs/adrs/adr_template.md
id_digit_number: 4
````

  

### Create a new ADR

**Simple Usage**

```
adrgen create "My new adr"
```

**Specify meta parameters**

```
adrgen create "My new adr" -m "components, technologies"
```

The meta parameters will add a meta section on the top of the generated file.  
```
---
components: ""
technologies: ""
---

# My new adr
```

They could be useful for an automated process that uses the ADR files to generate a living documentation website.

### Update the status of an ADR File

We can change the status for an ADR file specifying its ID and the new status with the **status** command

```
adrgen status 9 "accepted"
```


### To-Do

- [ ] Configure the ID format (for example, 0001-my-adr.md instead of 1-my-adr.md)  
- [ ] Supersede one ADR with another one  
- [ ] Amend one ADR with another one  
- [ ] Generating a trace with the ADR history  
- [ ] Change statuses in bulk
- [ ] Generating reports