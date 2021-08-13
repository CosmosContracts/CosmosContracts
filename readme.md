# Junø

![JUNO TWITTTT](https://user-images.githubusercontent.com/79812965/128625844-41ad62ea-ef27-49ad-81cc-915a2d1a6fc7.png)

Open source platform for interoperable smart contracts which automatically executes, controls or documents a procedure of relevant events and actions 
according to the terms of such contract or agreement to be valid & usable across multiple sovereign networks.

Juno as a **sovereign public blockchain** in the Cosmos ecosystem, aims to provide a sandbox environment for the deployment 
of such interoperable smart contracts. The network serves as a **decentralized, permissionless & censorship resistant** avenue 
for developers to efficiently and securely launch application specific smart contracts using proven frameworks 
and compile them in various languages **Rust & Go** with the potential future addition of C and C++
Battle tested contract modules such as CosmWasm, will allow for decentralized applications (dapps) to be compiled on robust and secure multi-chain smart contracts.
EVM support and additional specialized modules to be introduced after genesis subject to onchain governance.

At the heart of Cosmos is the Inter Blockchain Communication Protocol (IBC), which sets the table for an interoperable base layer 0 
to now be used to transfer data packets across thousands of independent networks supporting IBC. 
Naturally, the next evolutionary milestone is to enable cross-network smart contracts.

The Juno blockchain is built using the **Cosmos SDK framework**. 
A generalized framework that simplifies the process of building secure blockchain applications on top of Tendermint BFT. 
It is based on two major principles: Modularity & capabilities-based security.

Agreement on the network is reached via **Tendermint BFT consensus**.

Tendermint BFT is a solution that packages the networking and consensus layers of a blockchain into a generic engine, 
allowing developers to focus on application development as opposed to the complex underlying protocol. 
As a result, Tendermint saves hundreds of hours of development time.

Juno originates from a **community driven initiative**, prompted by developers, validators & delegators in the Cosmos ecosystem.
The common vision is to preserve the neutrality, performance & reliability of the Cosmos Hub and offload smart contract deployment to a dedicated sister Hub. 
Juno plans to make an early connection to the Cosmos Hub enabling IBC transfers, cross-chain smart contracts and making use of shared security.

A decentralized launch of the Juno main-net is enabled by a large set of independent validators from across the blockchain space.
$Juno, the native asset has many functions like securing the Juno Hub & serving as a work token to give access to on-chain governance voting rights 
and to provide utility in the deployment and execution of smart contracts.


**What differentiates JUNO from other Smart Contract networks?**

⚪️ Interoperable smart contracts

⚪️ First mover advantage

⚪️ Open source

⚪️ Permissionless 

⚪️ Modular

⚪️ Wasm + (EVM)

⚪️ Compilation in multiple languages Rust & Go (C,C++)

⚪️ Highly scalable

⚪️ Ease of use

⚪️ Free & fair asset distribution (100% to staked atom only)

⚪️ Onchain governance

⚪️ Balanced governance (Zero top heavy control) 

⚪️ Grass roots community                                               
                                                     
⚪️ Decentralized
                                             




![GENESIS 3 (JUNO)](https://user-images.githubusercontent.com/79812965/128879584-a3cf4ac0-3ba8-4142-a4c4-7dc37880b6ad.png)





**Tokenomics & reward shedule** (updated on 24.07.2021)

⚪️ **Ticker**: JUNO

⚪️ **Decimals**: 6

⚪️ **Unit**: uJuno

⚪️ **Supply**: Snapshot of Cosmoshub-3 at 06:00 PM UTC on Feb 18th 2021

⚪️ **Rewards**: Fixed yearly reward schedule (Reward model below)

⚪️ **Community pool tax**: 5% of block rewards


✅ Circulating supply at genesis 64.903.243 $JUNO (64.9 Million)

✅ Max supply (After year 12): 185.562.268 JUNO (185.5 Million)


**Supply Breakdown**

⚪️ Stakedrop: 30.663.193 $JUNO

⚪️ Community Pool: 20.000.000 $JUNO

⚪️ Development Reserve (Multi-sig): 11.866.708 $JUNO

⚪️ Smart Contract Challenges: 2.373.341 $JUNO


**Genesis Distribution**

A 1:1 stakedrop is allocated to $ATOM stakers, giving the $JUNO genesis supply to staked $ATOM balances that had their assets bonded 
at the time of the Stargate snapshot on Feb. 18th 6:00 PM UTC. 
Addresses that qualify are included in the JUNO genesis block at launch. 
All exchange validators & their delegators are excluded from the genesis allocation. Additionally any unbonded ATOM at the time of the snapshot is excluded.
A whale cap was voted in by the community, effectively hard-capping $ATOM accounts at 50 thousand $ATOM in order to ensure a less top heavy allocation.
Approx 10% of the supply difference is allocated to the development reserve (multi-sig) address for the funding of core-development efforts. The remaining 90% of the excess supply to be allocated in the following ways (20 million $Juno community pool, Smart contract competition 2.373.341,66 million to be managed/distributed by the multi-sig committee. The remaining difference will not be included in the genesis file ie. burned)



**Reward Schedule**

⚪️ Phase 1: Fixed inflation 40% 

New Juno in year 1 = (+25.961.297)

Supply after year 1 = 90.864.540 $JUNO


⚪️ Phase 2: Fixed inflation 20% 

New Juno in year 2 = (+18.172.908)

Supply after year 2 = 109.037.449 JUNO


⚪️ Phase 3: Fixed inflation 10% 

New Juno in year 3= (+10.903.744)

Supply after year 3 = 119.941.194 JUNO


Once the inflation reaches 10% it gradually reduces on a fixed 1% basis each year.


⚪️ Phase 4 = Fixed 9% (+10.794.707) Supply = 130.735.901 JUNO

⚪️ Phase 5 = Fixed 8% (+10.458.872) Supply = 141.194.773 JUNO

⚪️ Phase 6 = Fixed 7% (+9.883.634) Supply = 151.078.407 JUNO

⚪️ Phase 7 = Fixed 6% (+9.064.704) Supply = 160.143.112  JUNO

⚪️ Phase 8 = Fixed 5% (+8.007.155) Supply = 168.150.267  JUNO

⚪️ Phase 9 = Fixed 4% (+6.726.010) Supply = 174.876.278 JUNO

⚪️ Phase 10 = Fixed 3% (+5.246.288) Supply = 180.122.566 JUNO

⚪️ Phase 11 = Fixed 2% (+3.602.451) Supply = 183.725.018 JUNO

⚪️ Phase 12 = Fixed 1% (+1.837.250) Supply = 185.562.268 JUNO 

JUNO MAX SUPPLY (185.5 Million)

After year 12 the inflation reward schedule ends. 
Network incentives would primarily come from smart contract usage & regular tx fees generated on the network.












**Juno** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Learn more

- [Juno](https://junochain.com)
- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Telegram](https://t.me/JunoNetwork)
- [Discord](https://discord.gg/QcWPfK4gJ2)
