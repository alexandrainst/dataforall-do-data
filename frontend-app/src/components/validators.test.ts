import { isPositiveNumber } from './validators'

test('isPositiveNumber', () => {
  const errorMsg = 'Skal være et positivt tal.'
  expect(isPositiveNumber('')).toBe(errorMsg)
  expect(isPositiveNumber('abc')).toBe(errorMsg)
  expect(isPositiveNumber('-1')).toBe(errorMsg)
  expect(isPositiveNumber('0')).toBe(errorMsg)
  expect(isPositiveNumber('4,b')).toBe(errorMsg)
  expect(isPositiveNumber('4,b,c')).toBe(errorMsg)
  expect(isPositiveNumber('4,b.c')).toBe(errorMsg)
  expect(isPositiveNumber(',,')).toBe(errorMsg)
  expect(isPositiveNumber('..')).toBe(errorMsg)
  expect(isPositiveNumber('4..5')).toBe(errorMsg)
  expect(isPositiveNumber('4,,5')).toBe(errorMsg)
  expect(isPositiveNumber('1')).toBe(true)
  expect(isPositiveNumber('1.5')).toBe(true)
  expect(isPositiveNumber('1,5')).toBe(true)
  expect(isPositiveNumber('123')).toBe(true)
  expect(isPositiveNumber('  4,5  ')).toBe(true)
  expect(isPositiveNumber('4,533')).toBe(true)
})
