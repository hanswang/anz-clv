# anz-clv #

A tool to validate credit limit hierarchies.

### Demo ###
![demo](https://github.com/hanswang/anz-clv/blob/master/clvdemo.gif)

### Makefile Command ###
- `make build` build the binary output as `clv`
- `make lint` format the source code
- `make test` run unit test
- `make cover` running test and generate coverage report in `cover.html`
- `make all` run `generate`, `lint` and `test`

### Code Structure ###
```
                                  ┌───┬──┐
     ┌──────┐                     │   └──┤
     │      │   Read file (-f)    │ CSV  │
     │ Main │◄────────────────────┤ File │
     │      │   Verbose Log (-v)  │      │
     └──┬───┘                     └──────┘
        │
        │Buffer
        │bytes
        │
   ┌────▼───────┐ Rows of Str ┌────────────┐
   │            ├────────────►│            │
   │            │             │ Row Parser │
   │            │◄────────────┤            │
   │            │   Entities  └────────────┘
   │ Processor  │
   │            │  Entities   ┌────────────┐
   │            ├────────────►│            │
   │            │             │ Entity     │
   │            │◄────────────┤ Aggregator │
   │            │   Reports   │            │
   └─────┬──────┘             └────────────┘
         │
         │
         │
         │
    ┌────▼──────┐
    │           │
    │ View      │ Display details
    │ Generator │ inside each report
    │           │
    └───────────┘
```

