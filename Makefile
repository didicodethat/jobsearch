TEMPL = ${HOME}/go/bin/templ

templates: */*.templ
	${TEMPL} generate

development: templates
	go run main.go d
