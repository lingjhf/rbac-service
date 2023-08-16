export function mediaQuery(
  query: string,
  callback: (mq: MediaQueryList | MediaQueryListEvent) => void
) {
  const mediaQuery = window.matchMedia(query)
  callback(mediaQuery)
  mediaQuery.addEventListener('change', callback)
}

export function onMove(options: {
  update?: (e: MouseEvent) => void
  end?: (e: MouseEvent) => void
}) {
  const move = (e: MouseEvent) => {
    options.update?.(e)
  }
  const endMove = (e: MouseEvent) => {
    options.end?.(e)
    window.removeEventListener('mousemove', move)
    window.removeEventListener('mouseup', endMove)
  }
  window.addEventListener('mousemove', move)
  window.addEventListener('mouseup', endMove)
}
