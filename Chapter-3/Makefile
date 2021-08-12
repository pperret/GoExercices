TOPTARGETS :=build test clean

SUBDIRS := $(wildcard Exercice-*/.)

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	cd $@ ; $(MAKE) $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)

	