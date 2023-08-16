FROM golang:1.21-alpine3.18
# install packages
RUN apk add --update --no-cache bash make git zsh curl tmux

# Make zsh your default shell for tmux
RUN echo "set-option -g default-shell /bin/zsh" >> /root/.tmux.conf

# install oh-my-zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"

VOLUME /go_gitmoji_cli
WORKDIR /go_gitmoji_cli

COPY go-gitmoji-cli /usr/bin/

ENTRYPOINT ["/bin/zsh"]
