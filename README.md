# dwitter
**dwitter** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com).

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

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
curl https://get.starport.com/xmonader/dwitter@latest! | sudo bash
```
`xmonader/dwitter` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)





```
  dwitterd git:(master) ✗ ./dwitterd tx dwitter create-tweet "aa" --from alice
{"body":{"messages":[{"@type":"/xmonader.dwitter.dwitter.MsgCreateTweet","creator":"cosmos19yaj85alahff2tt0a6lgpkcvdak29g9qpngqh4","content":"aa"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: 0A2A0A282F786D6F6E616465722E647769747465722E647769747465722E4D73674372656174655477656574
events:
- attributes:
  - index: true
    key: ZmVl
    value: ""
  type: tx
- attributes:
  - index: true
    key: YWNjX3NlcQ==
    value: Y29zbW9zMTl5YWo4NWFsYWhmZjJ0dDBhNmxncGtjdmRhazI5ZzlxcG5ncWg0LzI=
  type: tx
- attributes:
  - index: true
    key: c2lnbmF0dXJl
    value: dHRrSUpNeHhqZTM0NzIzaDZhdW9Ka1Fsc2t6TmM3dXRiQzVjUDhsUC9lc3JCYlpyaE83b1ZtTHhhWUpvWjU2emdvalVuVUtvNzVxbCtmUjhlaGtFN3c9PQ==
  type: tx
- attributes:
  - index: true
    key: YWN0aW9u
    value: Y3JlYXRlX3R3ZWV0
  type: message
gas_used: "44661"
gas_wanted: "200000"
height: "657"
info: ""
logs:
- events:
  - attributes:
    - key: action
      value: create_tweet
    type: message
  log: ""
  msg_index: 0
raw_log: '[{"events":[{"type":"message","attributes":[{"key":"action","value":"create_tweet"}]}]}]'
timestamp: ""
tx: null
txhash: E508A45B246777058D202E4E6E47A64C5E74F274F3D1FF70299317BE13C8A56B
➜  dwitterd git:(master) ✗ ./dwitterd tx dwitter create-tweet "aab" --from alice
{"body":{"messages":[{"@type":"/xmonader.dwitter.dwitter.MsgCreateTweet","creator":"cosmos19yaj85alahff2tt0a6lgpkcvdak29g9qpngqh4","content":"aab"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: 0A2E0A282F786D6F6E616465722E647769747465722E647769747465722E4D7367437265617465547765657412020801
events:
- attributes:
  - index: true
    key: ZmVl
    value: ""
  type: tx
- attributes:
  - index: true
    key: YWNjX3NlcQ==
    value: Y29zbW9zMTl5YWo4NWFsYWhmZjJ0dDBhNmxncGtjdmRhazI5ZzlxcG5ncWg0LzM=
  type: tx
- attributes:
  - index: true
    key: c2lnbmF0dXJl
    value: TlI0WmJmcTZqbXhYenJGT1d5SXhwYWhKOGtpRnZjaTRzOXk4RkJjQWFCNGNabkJpUS96SEVLQ1hNaTlFbmFCMXpqaDl5R2I5c2htQi9NN2FJbkg2SkE9PQ==
  type: tx
- attributes:
  - index: true
    key: YWN0aW9u
    value: Y3JlYXRlX3R3ZWV0
  type: message
gas_used: "44785"
gas_wanted: "200000"
height: "662"
info: ""
logs:
- events:
  - attributes:
    - key: action
      value: create_tweet
    type: message
  log: ""
  msg_index: 0
raw_log: '[{"events":[{"type":"message","attributes":[{"key":"action","value":"create_tweet"}]}]}]'
timestamp: ""
tx: null
txhash: 60A9C0AF5054F2D67197C178E7A52F9A56333D871C2A5938E809F598BC776A5E
➜  dwitterd git:(master) ✗ ./dwitterd tx dwitter create-tweet "aabc" --from alice
{"body":{"messages":[{"@type":"/xmonader.dwitter.dwitter.MsgCreateTweet","creator":"cosmos19yaj85alahff2tt0a6lgpkcvdak29g9qpngqh4","content":"aabc"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
code: 0
codespace: ""
data: 0A2E0A282F786D6F6E616465722E647769747465722E647769747465722E4D7367437265617465547765657412020802
events:
- attributes:
  - index: true
    key: ZmVl
    value: ""
  type: tx
- attributes:
  - index: true
    key: YWNjX3NlcQ==
    value: Y29zbW9zMTl5YWo4NWFsYWhmZjJ0dDBhNmxncGtjdmRhazI5ZzlxcG5ncWg0LzQ=
  type: tx
- attributes:
  - index: true
    key: c2lnbmF0dXJl
    value: azRDWmdId1U2WUlzNTR0TjVYOXdJejJlT1UwYjNURTFvRkU5Y29ocGNZZEVqNmRsQ3FwZTkxSGJWRDVOdjdLQ1dOclYya21kbXBHdUJrMTdrUDVCVXc9PQ==
  type: tx
- attributes:
  - index: true
    key: YWN0aW9u
    value: Y3JlYXRlX3R3ZWV0
  type: message
gas_used: "44825"
gas_wanted: "200000"
height: "666"
info: ""
logs:
- events:
  - attributes:
    - key: action
      value: create_tweet
    type: message
  log: ""
  msg_index: 0
raw_log: '[{"events":[{"type":"message","attributes":[{"key":"action","value":"create_tweet"}]}]}]'
timestamp: ""
tx: null
txhash: 98BBF4BF1838479B41F83B5DF7CA010C7B7D0255D09767A1B437EE9641DCD7C5
➜  dwitterd git:(master) ✗ ./dwitterd q dwitter tweets                         
Tweet:
- content: aa
  creator: cosmos19yaj85alahff2tt0a6lgpkcvdak29g9qpngqh4
  id: "0"
- content: aab
  creator: cosmos19yaj85alahff2tt0a6lgpkcvdak29g9qpngqh4
  id: "1"
- content: aabc
  creator: cosmos19yaj85alahff2tt0a6lgpkcvdak29g9qpngqh4
  id: "2"
pagination:
  next_key: null
  total: "3"


 ```