package main

type Status int

const (
	Unknow Status = iota
	Success
	Fail
	Denied
)

var statusValues = map[Status]string{
	Unknow:  "unknow",
	Success: "success",
	Fail:    "fail",
	Denied:  "denied",
}

// ALERTA: Esta abordagem deve ser usada com cautela,
// pois corre o risco de "out of bound" caso o número de
// declarações no bloco `const` seja diferente da quantidade
// de entradas na variável de valores (statusValues).
// O compilador não irá alertá-lo sobre o possível erro "out of bound"
// Uma possível alternativa é utilizar um bloco switch para retornar as strings:
//
// switch s {
// case Success:
//
//	return "success"
//
// case Fail:
//
//	return "fail"
//
// case Denied:
//
//	return "denied"
//
// default:
//
//	  return "unknow"
//	}
func (s Status) String() string {
	return statusValues[s]
}

func (s Status) EnumIndex() int {
	return int(s)
}

func StatusFrom(status string) Status {
	for k, v := range statusValues {
		if v == status {
			return k
		}
	}
	return Unknow
}

func main() {

}
