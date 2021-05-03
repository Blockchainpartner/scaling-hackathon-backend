#!/usr/bin/env python3

import argparse
import json
import os
import subprocess
import sys
import tempfile
from typing import Optional

from starkware.cairo.bootloader.generate_fact import get_cairo_pie_fact_info
from starkware.cairo.bootloader.hash_program import compute_program_hash_chain
from starkware.cairo.lang.compiler.assembler import Program
from starkware.cairo.lang.vm.crypto import get_crypto_lib_context_manager
from starkware.cairo.sharp.client_lib import CairoPie, ClientLib
from starkware.cairo.sharp.fact_checker import FactChecker
from starkware.cairo.sharp.sharp_client import init_client  # noqa

def get_job_status(args, command_args):
    parser = argparse.ArgumentParser(
        description='Retreive the status of a SHARP Cairo job.')
    parser.add_argument('job_key', type=str, help='The key identifying the job.')

    parser.parse_args(command_args, namespace=args)

    client = init_client(bin_dir=args.bin_dir)
    return client.get_job_status(args.job_key)

def submit(args, command_args):
    parser = argparse.ArgumentParser()
    parser.add_argument('--source', type=str, required=False)
    parser.add_argument('--program', type=str, required=False)
    parser.add_argument('--program_input', type=str, required=False)
    parser.add_argument('--cairo_pie', type=str, required=False)

    parser.parse_args(command_args, namespace=args)

    is_not_none = lambda x: 1 if x is not None else 0
    assert (
        is_not_none(args.source) + is_not_none(args.program) + is_not_none(args.cairo_pie) == 1), \
        'Exactly one of --source, --program, --cairo_pie must be specified.'

    client = init_client(bin_dir=args.bin_dir)

    if args.cairo_pie is not None:
        assert args.program_input is None, \
            'Error: --program_input cannot be specified with --cairo_pie.'
        cairo_pie = CairoPie.from_file(args.cairo_pie)
    else:
        if args.program is not None:
            program = Program.Schema().load(json.load(open(args.program)))
        else:
            assert args.source is not None
            program = client.compile_cairo(source_code_path=args.source)
        cairo_pie = client.run_program(program=program, program_input_path=args.program_input)

    fact = client.get_fact(cairo_pie)
    job_key = client.submit_cairo_pie(cairo_pie=cairo_pie)
    return str(f'["{job_key}", "{fact}"]')

def prepare_cairo_run():
	subparsers = {
		'submit': submit,
		'status': get_job_status,
		# 'is_verified': is_verified,
	}

	parser = argparse.ArgumentParser(description='A tool to communicate with SHARP.')
	parser.add_argument('command', choices=subparsers.keys())
	parser.add_argument('--bin_dir', type=str, default='')
	parser.add_argument('--flavor', type=str, default='Release', choices=['Debug', 'Release', 'RelWithDebInfo'])
	args, unknown = parser.parse_known_args()

	with get_crypto_lib_context_manager(args.flavor):
		try:
			res = subparsers[args.command](args, unknown)
		except Exception as exc:
			res = ["", ""]
	return res

res = prepare_cairo_run()
print(res, end="")