package practice7_16

import (
	"html/template"
	"net/http"
	"regexp"
)

var html = `<DOC TYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>電卓</title>

	<style>
		.numbers {
			display: grid;
			grid-template-columns: 1fr 1fr 1fr 1fr;
			gap: 1rem;
		}
		.row5 {
			grid-row-start: 5;
		}
		.row4 {
			grid-row-start: 4;
		}
		.row3 {
			grid-row-start: 3;
		}
		.row2 {
			grid-row-start: 2;
		}
		.row1 {
			grid-row-start: 1;
		}
		.op {
			grid-column-start: 4;
		}
	</style>

</head>
<body>
<div>
	input: {{ .Input }}
</div>
<div>
	result: {{ .Result }}
</div>
<form action="">
	<div class="numbers">
		<input type="submit" name="in" value="0" class="row5">
		<input type="submit" name="in" value="1" class="row4">
		<input type="submit" name="in" value="2" class="row4">
		<input type="submit" name="in" value="3" class="row4">
		<input type="submit" name="in" value="4" class="row3">
		<input type="submit" name="in" value="5" class="row3">
		<input type="submit" name="in" value="6" class="row3">
		<input type="submit" name="in" value="7" class="row2">
		<input type="submit" name="in" value="8" class="row2">
		<input type="submit" name="in" value="9" class="row2">
		<input type="submit" name="in" value="." class="row5">
		<input type="submit" name="in" value="+" class="row5 op">
		<input type="submit" name="in" value="-" class="row4 op">
		<input type="submit" name="in" value="*" class="row3 op">
		<input type="submit" name="in" value="/" class="row2 op">
		<input type="submit" name="in" value="AC" class="row1 op">
	</div>
	<input type="hidden" name="input" value="{{ .Input }}">
</form>

</body>
</html>
`

type Server struct {
}

type viewParams struct {
	Input  string
	Result float64
}

func NewServer() *Server {
	return new(Server)
}

func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	var in string
	if validate(request) {
		in = request.Form.Get("in")
	}

	var input string
	var result float64
	if in == "AC" {
		input = ""
		result = 0
	} else {
		input = request.Form.Get("input")
		input += in

		calc := NewCalculator()
		result = calc.Calc(input)
	}

	params := viewParams{Input: input, Result: result}

	writer.Header().Set("Content-Type", "text/html")
	t := template.Must(template.New("html").Parse(html))
	t.Execute(writer, params)
}

func validate(request *http.Request) bool {
	in := request.Form.Get("in")
	if regexp.MustCompile("[0-9+*/.-]+").MatchString(in) {
		return true
	}
	if in == "AC" {
		return true
	}
	return false
}
