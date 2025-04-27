module main

go 1.24.2

replace bloomfilter/bloom => ./pkg

require bloomfilter/bloom v0.0.0-00010101000000-000000000000

require github.com/spaolacci/murmur3 v1.1.0 // indirect
