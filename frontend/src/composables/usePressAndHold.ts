import { ref, watchEffect } from 'vue'

export function usePressAndHold(cb: () => void) {

  const TIMEOUT_DURATION = 1000

  const isDown = ref(false)

  const onPointerDown = () => {
    isDown.value = true
    addEventListener('pointerup', () => {
      isDown.value = false
    }, { once: true })
  }

  let timeout: any = 0

  watchEffect(() => {
    console.log('isDown', isDown.value)
    if (isDown.value) {
      timeout = setTimeout(() => {
        console.log('timeout fired')
        cb()
      }, TIMEOUT_DURATION)
    } else {
      clearTimeout(timeout)
    }
  })

  return {
    isDown,
    onPointerDown
  }

}
