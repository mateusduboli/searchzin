#!/usr/bin/env python3
import os
import re

include_pattern = r'^<!--\s+include\s+(?P<format>[\w]+)\s+(?P<include_file>[^\s]+)\s+-->$'
readme_location = os.path.dirname(__file__) + '/../README.md.tpl'

readme_replaced = ''
with open(readme_location, 'r') as readme:
    readme_contents = readme.read()
    readme_replacements = []
    for m in re.finditer(include_pattern, readme_contents, flags=re.MULTILINE):
        (content_start, content_end) = m.span()
        replacement = {
                'start': content_start,
                'end': content_end,
                'format': m.group('format'),
                'file': m.group('include_file')
        }
        readme_replacements.append(replacement)

    last_replacement_index = 0
    for replacement in readme_replacements:
        readme_replaced += readme_contents[last_replacement_index:replacement['start']]
        readme_replaced += '\n```' + replacement['format'] + '\n'
        with open(replacement['file']) as f:
            readme_replaced += f.read()
        readme_replaced += '```\n'
        last_replacement_index = replacement['end']
    readme_replaced += readme_contents[last_replacement_index:-1]

with open('README.md', 'r+') as f:
    old_readme = f.read()
    f.seek(0)
    f.write(readme_replaced)
    f.truncate()
    if old_readme != readme_replaced:
        print('README.md modified')
        exit(1)
    else:
        exit(0)
