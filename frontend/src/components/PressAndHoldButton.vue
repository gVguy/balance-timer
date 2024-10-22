<script setup lang="ts">
import TooltipHint from './TooltipHint.vue'
import { usePressAndHold } from '@/composables/usePressAndHold'

defineProps<{
  text: string
  hint?: string
  secondary?: boolean
}>()

const emit = defineEmits(['confirm'])

const { isDown, onPointerDown } = usePressAndHold(() => {
  emit('confirm')
})

</script>

<template>
  <TooltipHint :text="hint">
    <button
      class="secondary press-and-hold-button"
      :class="{ down: isDown }"
      @pointerdown="onPointerDown"
    >
      {{ text }}
    </button>
  </TooltipHint>
</template>

<style scoped lang="scss">
.press-and-hold-button {
  position: relative;
  &.down {
    &::after {
      position: absolute;
      display: block;
      content: '';
      inset: 0;
      background: rgba(var(--fg-rgb), .4);
      transform-origin: left center;
      animation: hold-down 1.1s cubic-bezier(.3,.6,.7,.4);
      @keyframes hold-down {
        from {
          transform: scaleX(0);
        }
      }
    }
  }
}
</style>