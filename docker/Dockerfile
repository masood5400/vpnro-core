FROM alpine:latest
ENV CONFIG='https://raw.githubusercontent.com/ircfspace/warpsub/main/export/warp#WARP%20(IRCF)'
ENV VERSION=v1.3.6
WORKDIR /hiddify
RUN apk add  curl tar gzip libc6-compat # iptables ip6tables

    RUN case "$(apk --print-arch)" in \
            x86_64) ARCH=amd64 ;; \
            i386) ARCH=386 ;; \
            armv7l) ARCH=arm ;; \
            aarch64) ARCH=arm64 ;; \
            *) echo "Unsupported architecture: $(apk --print-arch) $(uname -m)" && exit 1 ;; \
        esac && \
        curl -L -o vpnro-cli.tar.gz https://github.com/masood5400/vpnro-core/releases/download/${VERSION}/vpnro-cli-linux-$ARCH.tar.gz && \
        tar -xzf vpnro-cli.tar.gz && rm vpnro-cli.tar.gz
COPY hiddify.sh .
RUN chmod +x hiddify.sh
COPY hiddify.json ./data/

EXPOSE 2334
EXPOSE 2335
EXPOSE 6756
EXPOSE 6450


ENTRYPOINT [ "/hiddify/hiddify.sh" ]