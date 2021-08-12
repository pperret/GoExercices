TOPTARGETS :=build test clean

SUBDIRS := $(wildcard Chapter-*/.)

$(TOPTARGETS): go.mod $(SUBDIRS)
$(SUBDIRS):
	cd $@ ; $(MAKE) $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)

go.mod:
	go mod init GoExercices
	