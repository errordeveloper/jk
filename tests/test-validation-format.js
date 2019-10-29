import { formatError } from '@jkcfg/std/schema';
import { print } from '@jkcfg/std';

const justMsg = { msg: 'Just an error message' };
print(formatError(justMsg));

const msgAndPath = { msg: 'This field was wrong', path: '<root>.foo.bar' };
print(formatError(msgAndPath));

const msgPathStart = {
  msg: 'The field here was wrong',
  path: '<root>.foo.bar',
  start: { line: 3, column: 20 },
};
print(formatError(msgPathStart));

const all = {
  msg: 'The field here was wrong',
  path: '<root>.foo.bar',
  start: { line: 3, column: 20 },
  end: { line: 5, column: 0 },
};
print(formatError(all));
