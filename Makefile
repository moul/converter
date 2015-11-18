.PHONY: convey
convey:
	goconvey -cover -port=10042 -workDir="$(realpath .)" -depth=1
