%builtins output pedersen

from starkware.cairo.common.cairo_builtins import HashBuiltin
from starkware.cairo.common.math import assert_not_zero
from starkware.cairo.common.hash import hash2
from starkware.cairo.common.alloc import alloc

func checkIdentity{pedersen_ptr: HashBuiltin*}(address: felt, identity: felt*) -> (success: felt, outputHash: felt):
    alloc_locals
    local secret
    assert_not_zero(address)

    %{ ids.secret = int(program_input['secret'], 10) %}

    let output = 0
    let (res) = hash2{hash_ptr=pedersen_ptr}(address, secret)
    
    if res == [identity]:
        let (outputHash) = hash2{hash_ptr=pedersen_ptr}(res, secret)
        return (1, outputHash)
    end
    return (0, 0)
end

############################################################################### 
#  processRegistry: for each adresse in the registry, we will check if
#  the pedersen hash of (address, secret) matches one element of the identities
#  This function is a loop.
############################################################################### 
func processRegistry{pedersen_ptr: HashBuiltin*}(registry: felt*, n_r: felt, address: felt) -> (outputHash: felt):
    if n_r == 0:
        return (0)
    end

    let (result, outputHash) = checkIdentity(address, registry)
    if result == 1:
        return (outputHash)
    end

    let (outputHash) = processRegistry(registry + 1, n_r - 1, address)
    return (outputHash)
end

############################################################################### 
#  getRegistry: Retrieve the list of all the encrypted address in the registry
############################################################################### 
func computeRegistryHash{pedersen_ptr: HashBuiltin*}(registry: felt*, size: felt, tempHash: felt) -> (registryHash: felt):
    if size == 0:
        return (tempHash)
    end

    let (tempHash) = hash2{hash_ptr=pedersen_ptr}(tempHash, [registry])
    let (tempHash) = computeRegistryHash(registry + 1, size - 1, tempHash)
    return (tempHash)
end

func getRegistry() -> (list: felt*, n: felt):
    alloc_locals
    local n
    let (registry: felt*) = alloc()
    %{
        index = 0
        ids.n = len(program_input['registry'])
        for key in program_input['registry']:
            base_addr = ids.registry + index
            memory[base_addr] = int(key, 16)
            index += 1
    %}
    return (list=registry, n=n)
end

############################################################################### 
#  Entry point, main function, wich will returns the Output struct
############################################################################### 
func main{output_ptr: felt*, pedersen_ptr: HashBuiltin*} ():
    alloc_locals
    local address
    let output = cast(output_ptr, Output*)
    let output_ptr = output_ptr + Output.SIZE

    %{ids.address = int(program_input['address'], 16)%}
    assert_not_zero(address)

    let (registry, n_r) = getRegistry()
    local registry: felt* = registry
    local n_r: felt = n_r

    let (outputHash) = processRegistry(registry, n_r, address)
    local outputHash: felt = outputHash
    assert_not_zero(outputHash)

    let (registryHash) = computeRegistryHash(registry, n_r, 0)

    assert output.hash = outputHash
    assert output.registryHash = registryHash
    return ()
end

############################################################################### 
#  STRUCTS
############################################################################### 
struct Output:
    member hash: felt
    member registryHash: felt
end