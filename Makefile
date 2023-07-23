ROOT_CA_KEY=rootCA.key
ROOT_CA_CRT=rootCA.crt
ROOT_CA_SRL=rootCA.srl
SIGNER_CSR=signer.csr
SIGNER_CRT=signer.crt
SIGNER_KEY=signer.key
SCRIPT=scripts/gencert.sh

run_script:
	./$(SCRIPT) $(ROOT_CA_KEY) $(ROOT_CA_CRT) $(SIGNER_CSR) $(SIGNER_CRT)

clean:
	rm -f $(ROOT_CA_KEY) $(ROOT_CA_CRT) ${ROOT_CA_SRL} $(SIGNER_CSR) $(SIGNER_CRT) ${SIGNER_KEY}

.PHONY: run_script clean
