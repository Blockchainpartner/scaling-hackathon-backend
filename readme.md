# Turbo Proof

> We're building a dApp based on identity registries that use STARKs (with Cairo) to verify some aspects of your identity.
>
> Let's say you need to prove that you're 18+. KYC first to either an identity provider or a state. The identity provider then releases a registry of all people 18+. Private information are encrypted by the person involved. Then, when they need to prove to a company or anyone that they are 18+, they execute the Cairo program checking the registry and verifying that the person is actually well registered, proof is then sent to the prover while programs output will be sent to a smart contract. Based on the Cairo program hash and outputs, we can now compute a fact that can be verified by the prover.
>
> If the prover confirms that the fact is actually one, then the person would have confirmed that they are 18+.
> Such process can be used for many other use cases and we're excited to showcase some of them in the dApp.

## Project
A user can then select a trip and select a fare based on his age or disability (if applicable).  
Before completing the booking, we generate a proof using a Cairo program that the user belongs, for example, to the registry of people under 25.  
That way, the user proves to the train company that they qualify for the right fare without having to reveal anything else.  

## Cairo setup & flow

### 1 - Setup the cairo programs
> We first need to compile the registry program with `cairo-compile registryHash.cairo --output registryHash.json --simple`.
> We will then be able to get it's hash with `cairo-hash-program --program registryHash.json` : `0x505ae0c3821ab690edb0fe45a63615f3600e168f3e60da824af1f28df54ecb1`.
>
> Then, we will need to compite the identities program with `cairo-compile main.cairo --output cairo.json --simple`.
> We will then be able to get it's hash with `cairo-hash-program --program cairo.json` : `0x7f481cf54b923941934abbc67454374eaff004f6fef703783c80d9ddb3fc3b5`.

--------------
### 2 - Deploying the contract
> We need to deploy the contract on ropsten with the following arguments :
> - _registriesProgramHash = `0x505ae0c3821ab690edb0fe45a63615f3600e168f3e60da824af1f28df54ecb1`
> - _identitiesProgramHash = `0x7f481cf54b923941934abbc67454374eaff004f6fef703783c80d9ddb3fc3b5`
> - cairoVerifier = `0x4Cd99A1FC780d874a34008fdC6da2961d540fE64`
>  
> The SmartContract is deployed here : [`0x274889F6864Bc0493BfEe3CF292A2A0ba1A76951`](https://ropsten.etherscan.io/address/0x274889F6864Bc0493BfEe3CF292A2A0ba1A76951).
>
>There is 2 notable functions: 
>```go
>function updateRegistry(uint256 registryKey, uint256 oldRegistryHash, uint256 newRegistryHash) onlyOwner public returns (bytes32)
>```
>This function will be used to update a specific registry, aka the list of addresses in it. Only the owner can update this list. We have one list per key.
>
>```go
>function proveIdentity(uint256 registryKey, uint256 hash, uint256 registryHash) public returns (bytes32)
>```
>This function will be used verify, directly on the smartcontract, that the output provided are valid, aka that the proof is valid.

-----------

### 3 - Creating some identities
>In order to work with this cairo program, we will first need some identities with the following informations :
>| ADDRESS (secret) | ENCRYPTED_ADDRESS | ENCRYPTION_KEY |
>| --- | --- | --- |
>| `0x02cEDd50ef234A2ee3CD3B87120e4367B37a3E61` | `0x6a3044f3172368c38bf9e73393275fd7e50a5e951ff8c09e4dd7b22d95a34f5` | `209076780032371316007324417983357412752269248774484056340191384479060305304` |
>| `0x9BfaDaa08AAf30EA4D6D00e86ed24EfB48554BD8` | `0x6ed120ae443c419aee7e3ee2c98575edef4612ad381fe88d817f04e9e5911ae` | `421816447642649854423897232527100597415899437911145406545557814348817499871` |
>| `0x09f3e857E4b14B7EF4D8F1a023e23F18f5e7ACF8` | `0x138e341ac1472ad700053c6a0e0fcdf41e7e4851ba2917fa9075f5e31e5ee2e` | `2671956092323553410592816235593792137884111017312175836053044049667031985252` |
>| `0xE5fb37aAe298C435Cb30B50DB950522856b8603B` | `0x755bf706ee73a22d3720569d0d3e4a7f9d28de4993a3b9c98014831c520a210` | `1759491603449138067611099643394087734194409180470721848330741395166328985900` |
>| `0x4600dFcd1e585BC39432CD1dceF973acA6225562` | `0x53465c95bcfb00da727a9ffd24c5bf324329de9fc27c57e53a86a23cfda1044` | `2023474883526219552093585673411987674931316574701963193192088079542665588336` |
>| `0x0568E0779FD984bb3E38921c466c0a5F89752e00` | `0x84afc9b9cdd81df7574ff788449dc3b6c0d6b5cf8a8f8f7536332d28b0add9` | `3124938967293004376366233214170809276646139449441931607366050588873625446273` |
> 
>  
>The `ADDRESS` is a secret. **This is what we will be try to prove** (that the address is in the registry). We are displaying it only for testing purpose.  
>The `ENCRYPTED_ADDRESS` matches `hash(ADDRESS, ENCRYPTION_KEY)`. This is the data that will be displayed in the registry.  
>The `ENCRYPTION_KEY` is a hash of a randomly generated password assciated to the user's account (`Web3.utils.randomHex(31)`) and the registryKey (here: `123456`) -> `hash(randomHex(31), 132456)`.
>
>With theses informations, we will produce the following JSON file :
>```json
>{
>	"oldRegistry": [],
>	"newRegistry": [
>		"0x6a3044f3172368c38bf9e73393275fd7e50a5e951ff8c09e4dd7b22d95a34f5",
>		"0x6ed120ae443c419aee7e3ee2c98575edef4612ad381fe88d817f04e9e5911ae",
>		"0x138e341ac1472ad700053c6a0e0fcdf41e7e4851ba2917fa9075f5e31e5ee2e",
>		"0x755bf706ee73a22d3720569d0d3e4a7f9d28de4993a3b9c98014831c520a210",
>		"0x53465c95bcfb00da727a9ffd24c5bf324329de9fc27c57e53a86a23cfda1044",
>		"0x84afc9b9cdd81df7574ff788449dc3b6c0d6b5cf8a8f8f7536332d28b0add9"
>	]
>}
>```
>This file will be used as our initial `registryHashInput` for the registry `123456`.  

--------------

>### 4 - Proving the identities registry
>By running `cairo-run --program=registryHash.json --print_output --layout=small --program_input=registryHashInput.json`, we will have the following output :
>```
>Program output:
>  0
>  2672368424450419303116754025943367656383112088074817038016074650154564643
>```
>Where `0` is our old hash, and `2672368424450419303116754025943367656383112088074817038016074650154564643` the new one.
>
>Now, let's prove that this hash is legit for our Cairo program, using the Cairo Sharp prover : `cairo-sharp submit --program registryHash.json --program_input registryHashInput.json`. 
>```
>Running...
>Submitting to SHARP...
>Job sent.
>Job key: e79ccff2-4376-4e9e-8b81-cc7e1fc86507
>Fact: 0xc076e47e63f5f79c38cadfee3ef7fbe2eb70d45c22e0e7ae5d10344b5b32d392
>```
>
>Here, we are telling to the Sharp Prover to compute our inputs with our program and to produce a `Fact` matching `hash(programHash, programOutputs)`. Our smartcontract will be able to check if this `fact` is valid with the prover's own smartContract, by calling the `isValid(fact)` function.  
>Our smartcontract will have to recompute the fact from the output, just to be sure.  
>We can use `cairo-sharp status e79ccff2-4376-4e9e-8b81-cc7e1fc86507` to check if the job has been processed (is live onChain).
>```
>PROCESSED
>```
>
>Now we can submit our transaction to the smartContract in order to confirm the proof with the following tx:
>```
>updateRegistry(
>	123456,
>	0,
>	2672368424450419303116754025943367656383112088074817038016074650154564643
>)
>```
> ⚠️ If the number `2672368424450419303116754025943367656383112088074817038016074650154564643` is negative, we would have to convert it to it's positive countervalue by doing the following math : `hex(MY_NEGATIVE_NUMBER + 2**251 + 17*2**192 + 1)` in order to work with Cairo's Prime & pedersen hashes.
>
>The transaction can [be found here](https://ropsten.etherscan.io/tx/0x5d4ffc61ae08c86e72b37d96a61e7fdfef05836d02dc2119abf59317b85fa28c). We now have some users in out registry.
>
>------------
>
>### 5 - Add more identities
>Yeah new users.
>| ADDRESS (secret) | ENCRYPTED_ADDRESS | ENCRYPTION_KEY |
>| --- | --- | --- |
>| `0x9E63B020ae098E73cF201EE1357EDc72DFEaA518` | `0x2c3dfa2e5c789f23006973dabc65c817704156344d2fb49aac7f237fb96c6ab` | `3262016890316122496475965907754361478299744245975029426120053541882877319917` |
>| `0xf51f06C26cA59954Ddce132E2a8EB026B0116658` | `0x3da30f7925ce33a7e93b186b8bd52ff7f64c52a42d79bebe891e132cb3d2b77` | `1953370059234117649198160759513056361430396932740856017941367296144751022849` |
>
>
>So updated JSON :
>```json
>{
>	"oldRegistry": [
>		"0x6a3044f3172368c38bf9e73393275fd7e50a5e951ff8c09e4dd7b22d95a34f5",
>		"0x6ed120ae443c419aee7e3ee2c98575edef4612ad381fe88d817f04e9e5911ae",
>		"0x138e341ac1472ad700053c6a0e0fcdf41e7e4851ba2917fa9075f5e31e5ee2e",
>		"0x755bf706ee73a22d3720569d0d3e4a7f9d28de4993a3b9c98014831c520a210",
>		"0x53465c95bcfb00da727a9ffd24c5bf324329de9fc27c57e53a86a23cfda1044",
>		"0x84afc9b9cdd81df7574ff788449dc3b6c0d6b5cf8a8f8f7536332d28b0add9"
>	],
>	"newRegistry": [
>		"0x6a3044f3172368c38bf9e73393275fd7e50a5e951ff8c09e4dd7b22d95a34f5",
>		"0x6ed120ae443c419aee7e3ee2c98575edef4612ad381fe88d817f04e9e5911ae",
>		"0x138e341ac1472ad700053c6a0e0fcdf41e7e4851ba2917fa9075f5e31e5ee2e",
>		"0x755bf706ee73a22d3720569d0d3e4a7f9d28de4993a3b9c98014831c520a210",
>		"0x53465c95bcfb00da727a9ffd24c5bf324329de9fc27c57e53a86a23cfda1044",
>		"0x84afc9b9cdd81df7574ff788449dc3b6c0d6b5cf8a8f8f7536332d28b0add9",
>		"0x2c3dfa2e5c789f23006973dabc65c817704156344d2fb49aac7f237fb96c6ab",
>		"0x3da30f7925ce33a7e93b186b8bd52ff7f64c52a42d79bebe891e132cb3d2b77"
>	]
>}
>```
>
>So new output
>```
>Program output:
>  2672368424450419303116754025943367656383112088074817038016074650154564643
>  -727187437430369499786061208688114383190983125936168187665846498250270412883
>```
>
>So new sharp 
>```
>Running...
>Submitting to SHARP...
>Job sent.
>Job key: 3967b1f1-e994-426f-9872-485bb1eb612e
>Fact: 0xba3365415dd3f5d4e92c4055cfb48eb561d773fe3e58a603db07dc692aef6cf5
>```
>
>So new TX (with the prime dark magic)
>```
>updateRegistry(
>	123456,
>	2672368424450419303116754025943367656383112088074817038016074650154564643,
>	2891315351235761713911261574406955722432124089395428512307245557885601607598
>)
>```
>
>And the [TX is here](https://ropsten.etherscan.io/tx/0xd3263c34b3509ad131e86c04a34ebf66162a3197b9a8f9bebaa9650e527019d2), with a new hash in the smartContract for the registryKey `123456` : `2891315351235761713911261574406955722432124089395428512307245557885601607598`
>
>------------------
>
>### 5 - Proving i'm IN !
>In order to prove that I am in this registry, we will use the second Cairo Program (the one with the following hash: `0x505ae0c3821ab690edb0fe45a63615f3600e168f3e60da824af1f28df54ecb1`). The idea is to prove that from my address and my secret (I am the only one to know theses informations), I can retrieve one of the entry on the registry.  
>In order to do that we will need a new input file with my `secret`, my `addresse`, and the public registry of all the addresses registered.  
>Let's say my address is `0x9E63B020ae098E73cF201EE1357EDc72DFEaA518`. With the table above, we can retrieve my secret, aka `3262016890316122496475965907754361478299744245975029426120053541882877319917`. Obviously, theses informations are here only for the purpose of this example.  
>```json
>{
>	"secret": "3262016890316122496475965907754361478299744245975029426120053541882877319917",
>	"address": "0x9E63B020ae098E73cF201EE1357EDc72DFEaA518",
>	"registry": [
>		"0x6a3044f3172368c38bf9e73393275fd7e50a5e951ff8c09e4dd7b22d95a34f5",
>		"0x6ed120ae443c419aee7e3ee2c98575edef4612ad381fe88d817f04e9e5911ae",
>		"0x138e341ac1472ad700053c6a0e0fcdf41e7e4851ba2917fa9075f5e31e5ee2e",
>		"0x755bf706ee73a22d3720569d0d3e4a7f9d28de4993a3b9c98014831c520a210",
>		"0x53465c95bcfb00da727a9ffd24c5bf324329de9fc27c57e53a86a23cfda1044",
>		"0x84afc9b9cdd81df7574ff788449dc3b6c0d6b5cf8a8f8f7536332d28b0add9",
>		"0x2c3dfa2e5c789f23006973dabc65c817704156344d2fb49aac7f237fb96c6ab",
>		"0x3da30f7925ce33a7e93b186b8bd52ff7f64c52a42d79bebe891e132cb3d2b77"
>	]
>}
>```
>With this `input.json` file, I can run the following command to get the output bellow : `cairo-run --program=cairo.json --print_output --layout=small --program_input=input.json`.
>```
>Program output:
>  829738850691260145934808358025321514871032962339753968951012900191018601261
>  -727187437430369499786061208688114383190983125936168187665846498250270412883
>```
>This output contains 2 informations needed to prove the information :
>- `829738850691260145934808358025321514871032962339753968951012900191018601261` is a specific hash corresponding to the element in the registry I match + my secret. In our case it's `hash(0x2c3dfa2e5c789f23006973dabc65c817704156344d2fb49aac7f237fb96c6ab, 3262016890316122496475965907754361478299744245975029426120053541882877319917)`.
>- `-727187437430369499786061208688114383190983125936168187665846498250270412883` is the registry hash.
>
>Just like we did for the registry update, we will prove this with the Sharp Prover with `cairo-sharp submit --program cairo.json --program_input input.json`
>```
>Running...
>Submitting to SHARP...
>Job sent.
>Job key: 866a23ce-2320-4233-99cc-9dd06ad7c889
>Fact: 0xabab4d2b54dd76e3b6095d6b9a1351aace18e8e9f07696e5f9d00c2a8d689bca
>```
>
>And then, we will [submit a new transaction](https://ropsten.etherscan.io/tx/0x17ba3a1630131b1e5459c8d2897b24e88d13a44dce98f862f4d636994900d419) to prove our fact :
>```
>proveIdentity(
>	123456,
>	829738850691260145934808358025321514871032962339753968951012900191018601261,
>	2891315351235761713911261574406955722432124089395428512307245557885601607598
>)
>```
>
>And the transaction is proved. That's it.