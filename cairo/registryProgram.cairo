%builtins output pedersen

from starkware.cairo.common.cairo_builtins import HashBuiltin
from starkware.cairo.common.math import assert_not_equal
from starkware.cairo.common.hash import hash2
from starkware.cairo.common.alloc import alloc

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

func getOldRegistry() -> (list: felt*, n: felt):
    alloc_locals
    local n
    let (registry: felt*) = alloc()
    %{
        index = 0
        ids.n = len(program_input['oldRegistry'])
        for key in program_input['oldRegistry']:
            base_addr = ids.registry + index
            memory[base_addr] = int(key, 16)
            index += 1
    %}
    return (list=registry, n=n)
end

func getNewRegistry() -> (list: felt*, n: felt):
    alloc_locals
    local n
    let (registry: felt*) = alloc()
    %{
        index = 0
        ids.n = len(program_input['newRegistry'])
        for key in program_input['newRegistry']:
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

    let (oldRegistry, oldN_r) = getOldRegistry()
    local oldRegistry: felt* = oldRegistry
    local oldN_r: felt = oldN_r

    let (oldRegistryHash) = computeRegistryHash(oldRegistry, oldN_r, 0)
    local oldRegistryHash: felt = oldRegistryHash

    let (newRegistry, newN_r) = getNewRegistry()
    local newRegistry: felt* = newRegistry
    local newN_r: felt = newN_r

    let (newRegistryHash) = computeRegistryHash(newRegistry, newN_r, 0)

    assert_not_equal(oldRegistryHash, newRegistryHash)

    assert output.oldHash = oldRegistryHash
    assert output.newHash = newRegistryHash
    return ()
end

############################################################################### 
#  STRUCTS
############################################################################### 
struct Output:
    member oldHash: felt
    member newHash: felt
end