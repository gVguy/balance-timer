<script setup lang="ts">
import { Stop } from '@w/go/main/App'
import { ref, watchEffect } from 'vue'

const isDown = ref(false)


const onDown = () => {
  isDown.value = true
  addEventListener('pointerup', () => {
    isDown.value = false
  }, { once: true })
}

let timeout: any = 0

watchEffect(() => {
  if (isDown.value) {
    timeout = setTimeout(() => {
      Stop()
    }, 1500)
  } else {
    clearTimeout(timeout)
  }
})

</script>

<template>
  <div class="stop-button-container">
    <button
      class="secondary stop-button"
      :class="{ down: isDown }"
      @pointerdown="onDown"
    >
      End session
    </button>
    <span class="hint">
      Press and hold to quit
    </span>
  </div>
</template>

<style scoped lang="scss">
.stop-button-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 8px;
  .stop-button {
    position: relative;
    &.down {
      &::after {
        position: absolute;
        display: block;
        content: '';
        inset: 0;
        background: rgba(var(--fg-rgb), .4);
        transform-origin: left center;
        animation: hold-down 1.6s cubic-bezier(.3,.6,.7,.4);
        @keyframes hold-down {
          from {
            transform: scaleX(0);
          }
        }
      }
    }
    &:hover ~ .hint {
      opacity: 1;
    }
  }
  .hint {
    transition: opacity .4s;
    opacity: 0;
    color: rgba(var(--fg-rgb), .3);
  }
}
</style>