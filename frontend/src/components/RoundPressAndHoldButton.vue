<script setup lang="ts">
import { usePressAndHold } from '@/composables/usePressAndHold'
import { computed } from 'vue'

const emit = defineEmits(['confirm'])

const { onPointerDown, isDown } = usePressAndHold(() => {
  emit('confirm')
})

const SIZE = 48
const STROKE_WIDTH = 5

const radius = computed(() => (SIZE / 2) - STROKE_WIDTH)
const circumference = computed(() => 2 * Math.PI * radius.value)

</script>

<template>
  <button
    class="secondary round-press-and-hold-button"
    @pointerdown="onPointerDown"
    :style="{
      width: SIZE + 'px',
      height: SIZE + 'px'
    }"
  >
    <svg :width="SIZE" :height="SIZE" >
      <circle
        class="circle"
        :class="{ 'down': isDown }"
        :style="{
          strokeDasharray: circumference,
          strokeDashoffset: isDown ? 0 : circumference
        }"
        cx="50%"
        cy="50%"
        :r="radius"
        fill="none"
        :stroke-width="STROKE_WIDTH"
      />
    </svg>
    <div class="icon-container">
      <slot />
    </div>
  </button>
</template>

<style scoped lang="scss">
.round-press-and-hold-button {
  border-radius: 50%;
  position: relative;
  overflow: clip;
  &:deep(*) {
    cursor: pointer;
  }
  .icon-container {
    width: 24px;
    height: 24px;
    &:deep(svg) {
      width: 100%;
      height: 100%;
      fill: rgba(var(--fg-rgb), .7);
      transition: fill .4s;
    }
  }
  &:hover {
    .icon-container:deep(svg) {
      fill: var(--fg-color);
    }
  }
  svg {
    position: absolute;
    rotate: -90deg;
    scale: 1.1;
    circle {
      stroke: rgba(var(--fg-rgb), .7);
      &.down {
        transition: stroke-dashoffset 1.1s cubic-bezier(.3,.6,.7,.4);
      }
    }
  }
}
</style>
