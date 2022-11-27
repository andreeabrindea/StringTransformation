# GO String Transformation Service 

### Description: 
##### A program in Golang which receives a CSV formatted string input(ex: input_example), and returns a new string(ex: output_example) with the following requirements:
```
input_example:

----------------------------------------------------

full_name, email, location

Anita, anita@email.com, California

Aron, aron.bla@email.com, California

Aron, aron.bla@email.com, California

Cosmin, kox@bla.com, Giurgiu

Crina, ggl@test.com, Letcani

Bogdan, vox@example.com, Resita

----------------------------------------------------

output_example:

----------------------------------------------------

A:

Anita, anita@email.com, California

Aron, aron.bla@email.com, California


C:

Cosmin, kox@bla.com, Giurgiu

Crina, ggl@test.com, Letcani

B:

Bogdan, vox@example.com, Resita
----------------------------------------------------
```
### User Guide:
##### clone this repository and to run the program you need to run the following commands:
```
 go run String_transformation.go

```
##### and to run the tests:
```
 go test .

```
