export const isPositiveNumber = (value: string) => {
  const errorMsg = 'Skal være et positivt tal.'
  const trimmed = value.trim()
  if (!trimmed) {
    return errorMsg
  }
  const normalized = trimmed.replace(',', '.')
  const validNumberPattern = /^[0-9]+(\.[0-9]+)?$/
  if (!validNumberPattern.test(normalized)) {
    return errorMsg
  }
  const number = parseFloat(normalized)
  if (isNaN(number) || number <= 0) {
    return errorMsg
  }
  return true
}
