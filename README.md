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

## Running the application

1. Create a `.env` file and add the following:

```
RABBIT_MQ_CONNECTION=amqp://guest:guest@localhost:5672/
```

2. Set up and run the infrastructure using docker compose

```sh
task up
```

3. Start the Peril server

```sh
task start:server
```

