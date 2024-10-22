<script setup lang="ts">
import TooltipHint from './TooltipHint.vue'
import RoundPressAndHoldButton from './RoundPressAndHoldButton.vue'

defineProps<{
  title: string
  text: string
}>()

defineEmits(['confirm'])

</script>

<template>
  <TooltipHint v-bind="$props">
    <RoundPressAndHoldButton @confirm="$emit('confirm')">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
        <path d="M16,18H18V6H16M6,18L14.5,12L6,6V18Z" />
      </svg>
    </RoundPressAndHoldButton>
  </TooltipHint>
</template>
