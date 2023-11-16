
components = components/hello.templ components/count.templ
go_files = cmd/server.go
css = input.css
node = node_modules/
out = out/cmd/

all: $(out) $(node) templates go tailwind


templates: $(components)
	templ generate $(components)

go: $(go_files)
	go build -o out/cmd/server -v $(go_files)

tailwind: $(css)
	npx tailwindcss -i input.css -o assets/css/index.css

$(out):
	mkdir -p out/cmd

$(node): 
	npm i

clean:
	rm -rf out/cmd
