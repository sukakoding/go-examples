## Issues
- Dilayer manakah custom error (NewProblem) harus dibuat?
- Bagaimana cara menghandle error yang dihasilkan dari package lain?
	* package eksternal dibungkus (abstraksi)
		- concern jika termasuk builtin bisa jadi banyak banget
- Perlukah dibuatkan function respond berdasarkan status?
	* 
- Perlukah menambahkan parameter alternatifStatus di function Error, untuk handling eksternal error?
	*
- Perlukah menggunakan merinci semua error eksternal yg belum ditangkap atau cukup dg error default?

- cara nge-produce error
	* default by function
- errors per app atau per domain/package?

## TODO
- Custom error, extended dari Problem
- Load list error dari yaml

## Reference
- https://golang.org/pkg/errors/
- https://www.digitalocean.com/community/tutorials/creating-custom-errors-in-go