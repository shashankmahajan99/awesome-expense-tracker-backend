generate:   
	rm -rf generated
	mkdir generated
	cd api; buf generate; cd ..
	make convert
	cp -r generated/* api/

convert:
	for file in generated/*swagger.json; do \
		base=$$(basename $$file .swagger.json); \
		if [[ $$base != *_v3 ]]; then \
			echo "Converting $$file"; \
			swagger2openapi $$file -o "generated/$${base}_v3.swagger.json"; \
		fi \
	done