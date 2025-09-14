# learn-pub-sub-starter (Peril)

## Setup

```sh
# Install `mise`
$ brew install mise

# Activate `mise` in your .profile
# Activation options: https://mise.jdx.dev/installing-mise.html#shells
#
# Run ONE of these:
$ echo 'eval "$(mise activate bash)"' >> ~/.bashrc
$ echo 'eval "$(mise activate zsh)"' >> "${ZDOTDIR-$HOME}/.zshrc"
$ echo 'mise activate fish | source' >> ~/.config/fish/config.fish

# Open and validate `mise.toml` before trusting (inside repository root)
$ mise trust

# Install the toolings
$ mise install

# Install projects' dependencies
$ task setup
```

## RabbitMQ

```sh
# Start RabbitMQ server
$ task rabbit:start

# Stop RabbitMQ server
$ task rabbit:stop

# View RabbitMQ logs
$ task rabbit:logs
```

## Running the application

```sh
# Start the RabbitMQ server
$ task rabbit:start

# Start our Peril server after the RabbitMQ server has booted
$ task start:server
```

