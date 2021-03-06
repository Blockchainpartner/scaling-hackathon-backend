import sys
import argparse
import json
import math
import os
import subprocess
import tempfile
import time
import starkware.python.python_dependencies as python_dependencies

from typing import Dict
from http.server import BaseHTTPRequestHandler
from starkware.cairo.lang.compiler.program import Program, ProgramBase
from starkware.cairo.lang.vm.crypto import get_crypto_lib_context_manager
from starkware.cairo.lang.vm.cairo_run import load_program  # noqa
from starkware.cairo.lang.vm.cairo_runner import CairoRunner
from starkware.cairo.lang.vm.memory_dict import MemoryDict
from starkware.cairo.lang.vm.security import verify_secure_runner
from starkware.cairo.lang.vm.vm import VmException
from starkware.cairo.lang.compiler.expression_simplifier import to_field_element

def retrieveOutput(runner, exception, output_callback=to_field_element):
	if 'output_builtin' not in runner.builtin_runners:
		return

	output = []
	output_runner = runner.builtin_runners['output_builtin']
	_, size = output_runner.get_used_cells_and_allocated_size(runner)
	for i in range(size):
		val = runner.vm_memory.get(output_runner.base + i)
		if val is not None:
			if i == exception:
				toAppend = output_callback(val=val, prime=runner.program.prime)
				toAppend = hex(toAppend + runner.program.prime)
				output.append(toAppend)
			else:
				output.append(str(output_callback(val=val, prime=runner.program.prime)))
	return output

def cairo_run(args, program_input):
	program: ProgramBase = load_program(args.program)
	initial_memory = MemoryDict()
	steps_input = args.steps

	runner = CairoRunner(program=program, layout=args.layout, memory=initial_memory, proof_mode=args.proof_mode)

	runner.initialize_segments()
	end = runner.initialize_main_entrypoint()
	runner.initialize_vm(hint_locals={'program_input': program_input})

	try:
		additional_steps = 1 if args.proof_mode else 0
		max_steps = steps_input - additional_steps if steps_input is not None else None
		runner.run_until_pc(end, max_steps=max_steps)
		runner.original_steps = runner.vm.current_step
		runner.end_run()
	except (VmException, AssertionError) as exc:
		raise exc

	runner.read_return_values()
	runner.finalize_segments_by_effective_size()
	verify_secure_runner(runner)
	runner.relocate()
	output = retrieveOutput(runner, args.exception)

	return output

def prepare_cairo_run():
	sys.argv.extend(['--layout', 'small'])
	sys.argv.extend(['--print_output'])
	parser = argparse.ArgumentParser(description='A tool to run Cairo programs.')
	parser.add_argument('--program', type=argparse.FileType('r'))
	parser.add_argument('--program_input', type=str)
	parser.add_argument('--exception', type=int)
	parser.add_argument('--print_output', action='store_true')
	parser.add_argument('--layout', default='plain')
	parser.add_argument('--flavor', type=str, choices=['Debug', 'Release', 'RelWithDebInfo'])
	parser.add_argument('--steps', type=int)
	parser.add_argument('--proof_mode', action='store_true')
	parser.add_argument('--input', type=str)
	python_dependencies.add_argparse_argument(parser)
	args = parser.parse_args()


	with get_crypto_lib_context_manager(args.flavor):
		try:
			res = cairo_run(args, json.loads(args.program_input))
		except VmException as err:
			res = ["0", "0"]
		except AssertionError as err:
			print(f'Error: {err}', file=sys.stderr)
			res = 0
	return res

res = prepare_cairo_run()
print(res, end="")