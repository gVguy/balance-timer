<script setup lang="ts">
import { Stop } from '@w/go/main/App'
import TooltipHint from './TooltipHint.vue'
import RoundPressAndHoldButton from './RoundPressAndHoldButton.vue'
</script>

<template>
  <TooltipHint
    title="End session"
    text="Press and hold to return to the New Session screen"
  >
    <RoundPressAndHoldButton @confirm="Stop">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <path d="M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z" />
      </svg>
    </RoundPressAndHoldButton>
  </TooltipHint>
</template>
