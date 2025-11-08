NAME=vvedenie_v_go

$(NAME).pdf: $(NAME).tex text/*.tex
	xelatex -shell-escape -interaction=nonstopmode $(NAME).tex|tail -n 20; echo -e \\a
	#xelatex -interaction=nonstopmode $(NAME).tex|tail -n 20; echo -e \\a

all: $(NAME).pdf

run: all

.phony: clean cleanall run

clean:
	@rm -rf *.aux text/*.aux *.toc *.ent *.log text/*.log *.idx *.minted _minted *.gz >/dev/null 2>/dev/null

cleanall: clean
	@rm -f $(NAME).pdf >/dev/null 2>/dev/null

