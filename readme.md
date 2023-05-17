# exabits

**exabits** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Run a single node testnet

You can also run a local testnet using a single node. On a local testnet, you will be the sole validator signing blocks.

**Step 1. Create network and account**

First, initialize your genesis file to bootstrap your network. Create a name for your local testnet and provide a moniker to refer to your node:

```bash
exabitd init <node_moniker> --chain-id=<testnet_name>
```

Next, create a Exabits account by running the following command:

```bash
exabitd keys add <account_name>
```

**Step 2. Add account to genesis**

Next, add your account to genesis and set an initial balance to start. Run the following commands to add your account and set the initial balance:

```bash
exabitd add-genesis-account $(exabitd keys show <account_name> -a) 100000000ebc
exabitd gentx <account_name> 100000000ebc --chain-id=<testnet_name>
exabitd collect-gentxs
```

**Step 3. Run Mises daemon**

Now you can start your private Exabits network:

```bash
exabitsd start
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run dev
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/exabits@latest! | sudo bash
```

`username/exabits` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

-   [Ignite CLI](https://ignite.com/cli)
-   [Tutorials](https://docs.ignite.com/guide)
-   [Ignite CLI docs](https://docs.ignite.com)
-   [Cosmos SDK docs](https://docs.cosmos.network)
-   [Developer Chat](https://discord.gg/ignite)
