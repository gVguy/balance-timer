<template>
  <div class="circle-progress">
    <svg :width="size" :height="size" :viewBox="`0 0 ${size} ${size}`" fill="none" xmlns="http://www.w3.org/2000/svg">
      <circle
        :cx="size / 2"
        :cy="size / 2"
        :r="radius"
        stroke="rgba(255, 255, 255, 0.25)"
        stroke-width="2"
      />
      <circle
        class="progress"
        :cx="size / 2"
        :cy="size / 2"
        :r="radius"
        stroke="var(--accent-color)"
        stroke-width="2"
        :stroke-dasharray="strokeDasharrayComplete + ' ' + strokeDasharrayMax"
      />
    </svg>
    <span class="value-display">
      {{ formattedValue }}
    </span>
  </div>
</template>

<script setup lang="ts">
import { useSession } from '@/composables/useSession'
import { computed } from 'vue'

const { session } = useSession()

const size = 222

const limit = computed(() => session.value.timeInitial)
const current = computed(() => session.value.timeRemaining)

const formattedValue = computed(() => session.value.timeRemainingString)

const radius = computed(() => {
  return Number(size) / 2 - 1
})

const strokeDasharrayMax = computed(() => {
  // r * п * 2
  return radius.value * Math.PI * 2
})

const strokeDasharrayComplete = computed(() => {
  return (current.value * strokeDasharrayMax.value) / limit.value
})

</script>

<style scoped lang="scss">
.circle-progress {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  .progress {
    transform-origin: center;
    transform: rotate(-90deg);
    transition: stroke-dasharray 1s linear;
  }
  .value-display {
    position: absolute;
    font-size: 28px;
  }
}
</style>
