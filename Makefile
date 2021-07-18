TOPTARGETS :=build test clean

SUBDIRS := $(wildcard Chapter-*/.)

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	cd $@ ; $(MAKE) $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)
